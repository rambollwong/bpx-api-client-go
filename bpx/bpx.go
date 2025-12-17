package bpx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	EndpointApi = "https://api.backpack.exchange/"
	EndpointWs  = "wss://ws.backpack.exchange/"
)

type Client struct {
	Key    string
	Secret string
	Proxy  string
	Window string

	httpCli *http.Client
}

func NewClient(key, secret string) *Client {
	cli := &Client{
		Key:    key,
		Secret: secret,
		Window: "60000",
	}
	cli.httpCli = &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			DisableCompression: false,
			Proxy:              http.ProxyFromEnvironment,
		},
	}

	return cli
}

func (c *Client) Asserts() Assets {
	return Assets{c: c}
}

func (c *Client) BorrowLendMarkets() BorrowLendMarkets {
	return BorrowLendMarkets{c: c}
}

func (c *Client) Markets() Markets {
	return Markets{c: c}
}

func (c *Client) System() System {
	return System{c: c}
}

func (c *Client) Trades() Trades {
	return Trades{c: c}
}

func (c *Client) Account() Account {
	return Account{c: c}
}

func (c *Client) BorrowLend() BorrowLend {
	return BorrowLend{c: c}
}

func (c *Client) Capital() Capital {
	return Capital{c: c}
}

func (c *Client) Futures() Futures {
	return Futures{c: c}
}

func (c *Client) History() History {
	return History{c: c}
}

func (c *Client) Order() Order {
	return Order{c: c}
}

type WsClient struct {
	done chan struct{}

	Key           string
	Secret        string
	Proxy         string
	Window        string
	autoReconnect bool

	mu     sync.RWMutex
	ws     *websocket.Conn
	status int // 0: disconnected, 1: connecting, 2: connected 3:disconnecting

	handlers        map[string][]WsMessageHandler
	orderSub        map[string][]WsOrderDataHandler
	positionSub     map[string][]WsPositionDataHandler
	bookTickerSub   map[string][]WsBookTickerDataHandler
	depthSub        map[string][]WsDepthDataHandler
	klineSub        map[string][]WsKlineDataHandler
	liquidationSub  map[string][]WsLiquidationDataHandler
	markPriceSub    map[string][]WsMarkPriceDataHandler
	tickerSub       map[string][]WsTickerDataHandler
	tradeSub        map[string][]WsTradeDataHandler
	openInterestSub map[string][]WsOpenInterestDataHandler

	onConnectHooks    []func()
	onMessageHooks    []func([]byte)
	onDisconnectHooks []func(error)

	errChan chan error
}

func NewWsClient(key, secret string) *WsClient {
	ws := &WsClient{
		done:    make(chan struct{}),
		Key:     key,
		Secret:  secret,
		Proxy:   "",
		Window:  "60000",
		status:  WsStatusDisconnected,
		errChan: make(chan error, 1),
	}
	return ws
}

func (ws *WsClient) WithAutoReconnect(auto bool) *WsClient {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.autoReconnect = auto
	return ws
}

func (ws *WsClient) WithProxy(proxy string) *WsClient {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.Proxy = proxy
	return ws
}

func (ws *WsClient) WithWindow(window string) *WsClient {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.Window = window
	return ws
}

func (ws *WsClient) OnConnect(hook func()) *WsClient {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onConnectHooks = append(ws.onConnectHooks, hook)
	return ws
}

func (ws *WsClient) OnMessage(hook func([]byte)) *WsClient {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onMessageHooks = append(ws.onMessageHooks, hook)
	return ws
}

func (ws *WsClient) OnDisconnect(hook func(error)) *WsClient {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onDisconnectHooks = append(ws.onDisconnectHooks, hook)
	return ws
}

func (ws *WsClient) ErrChan() <-chan error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.errChan
}

func (ws *WsClient) Connect(ctx context.Context) error {
	if ws.getStatus() != WsStatusDisconnected {
		return fmt.Errorf("websocket connection is not disconnected")
	}
	ws.setStatus(WsStatusConnecting)
	var dialer websocket.Dialer
	if ws.Proxy != "" {
		proxyUrl, err := url.Parse(ws.Proxy)
		if err != nil {
			return fmt.Errorf("parse proxy url error: %w", err)
		}
		dialer = websocket.Dialer{Proxy: http.ProxyURL(proxyUrl)}
	} else {
		dialer = websocket.Dialer{Proxy: http.ProxyFromEnvironment}
	}
	conn, _, err := dialer.DialContext(ctx, EndpointWs, nil)
	if err != nil {
		ws.setStatus(WsStatusDisconnected)
		return fmt.Errorf("dial error: %w", err)
	}
	ws.setConn(conn)

	// Start message handling
	go ws.handleMessages(ctx)

	// Start keep alive
	go ws.keepAlive(ctx)

	// Call onConnectHooks
	for _, hook := range ws.onConnectHooks {
		hook()
	}

	ws.setStatus(WsStatusConnected)

	// auto resubscribe
	if err := ws.resubscribe(); err != nil {
		return err
	}

	return nil
}

func (ws *WsClient) Disconnect(err error) {
	ws.setStatus(WsStatusDisconnecting)

	select {
	case <-ws.done:
	default:
		close(ws.done)
	}

	conn := ws.getConn()
	_ = conn.Close()

	ws.setStatus(WsStatusDisconnected)
	for _, hook := range ws.onDisconnectHooks {
		hook(err)
	}
}

func (ws *WsClient) Subscribe(stream string, isPrivate bool, handler WsMessageHandler) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.handlers == nil {
		ws.handlers = make(map[string][]WsMessageHandler)
	}
	handlers, ok := ws.handlers[stream]
	if !ok {
		handlers = make([]WsMessageHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.handlers[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, isPrivate); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeOrderUpdate(symbol string, handler WsOrderDataHandler) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "account.orderUpdate"
	if symbol != "" {
		stream += "." + symbol
	}

	if ws.orderSub == nil {
		ws.orderSub = make(map[string][]WsOrderDataHandler)
	}
	handlers, ok := ws.orderSub[stream]
	if !ok {
		handlers = make([]WsOrderDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.orderSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, true); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribePositionUpdate(symbol string, handler WsPositionDataHandler) error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "account.positionUpdate"
	if symbol != "" {
		stream += "." + symbol
	}

	if ws.positionSub == nil {
		ws.positionSub = make(map[string][]WsPositionDataHandler)
	}
	handlers, ok := ws.positionSub[stream]
	if !ok {
		handlers = make([]WsPositionDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.positionSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, true); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeBookTicker(symbol string, handler WsBookTickerDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "bookTicker." + symbol

	if ws.bookTickerSub == nil {
		ws.bookTickerSub = make(map[string][]WsBookTickerDataHandler)
	}
	handlers, ok := ws.bookTickerSub[stream]
	if !ok {
		handlers = make([]WsBookTickerDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.bookTickerSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeDepth(symbol, aggregatedInterval string, handler WsDepthDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "depth"
	if aggregatedInterval != "" {
		stream += "." + aggregatedInterval
	}
	stream += "." + symbol

	if ws.depthSub == nil {
		ws.depthSub = make(map[string][]WsDepthDataHandler)
	}
	handlers, ok := ws.depthSub[stream]
	if !ok {
		handlers = make([]WsDepthDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.depthSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeKline(symbol, interval string, handler WsKlineDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}
	if interval == "" {
		return fmt.Errorf("interval cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "kline." + interval + "." + symbol
	if ws.klineSub == nil {
		ws.klineSub = make(map[string][]WsKlineDataHandler)
	}
	handlers, ok := ws.klineSub[stream]
	if !ok {
		handlers = make([]WsKlineDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.klineSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeLiquidation(symbol string, handler WsLiquidationDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "liquidation." + symbol

	if ws.liquidationSub == nil {
		ws.liquidationSub = make(map[string][]WsLiquidationDataHandler)
	}
	handlers, ok := ws.liquidationSub[stream]
	if !ok {
		handlers = make([]WsLiquidationDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.liquidationSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeMarkPrice(symbol string, handler WsMarkPriceDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "markPrice." + symbol
	if ws.markPriceSub == nil {
		ws.markPriceSub = make(map[string][]WsMarkPriceDataHandler)
	}
	handlers, ok := ws.markPriceSub[stream]
	if !ok {
		handlers = make([]WsMarkPriceDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.markPriceSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeTicker(symbol string, handler WsTickerDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "ticker." + symbol
	if ws.tickerSub == nil {
		ws.tickerSub = make(map[string][]WsTickerDataHandler)
	}
	handlers, ok := ws.tickerSub[stream]
	if !ok {
		handlers = make([]WsTickerDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.tickerSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeOpenInterest(symbol string, handler WsOpenInterestDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "openInterest." + symbol
	if ws.openInterestSub == nil {
		ws.openInterestSub = make(map[string][]WsOpenInterestDataHandler)
	}
	handlers, ok := ws.openInterestSub[stream]
	if !ok {
		handlers = make([]WsOpenInterestDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.openInterestSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) SubscribeTrade(symbol string, handler WsTradeDataHandler) error {
	if symbol == "" {
		return fmt.Errorf("symbol cannot be empty")
	}

	ws.mu.Lock()
	defer ws.mu.Unlock()

	stream := "trade." + symbol
	if ws.tradeSub == nil {
		ws.tradeSub = make(map[string][]WsTradeDataHandler)
	}
	handlers, ok := ws.tradeSub[stream]
	if !ok {
		handlers = make([]WsTradeDataHandler, 0)
	}

	handlers = append(handlers, handler)
	ws.tradeSub[stream] = handlers

	if !ok && ws.status == WsStatusConnected {
		// write subscribe message
		if err := ws.writeSubscribeMsg(stream, false); err != nil {
			return err
		}
	}

	return nil
}
