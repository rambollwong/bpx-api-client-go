package bpx

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	WsStatusDisconnected = iota
	WsStatusConnecting
	WsStatusConnected
	WsStatusDisconnecting
)

type WsSubscribeMessage struct {
	Method    string   `json:"method,omitempty"`
	Params    []string `json:"params,omitempty"`
	Signature []string `json:"signature,omitempty"`
}

// WsMessageHandler is a function type for handling WebSocket messages
type WsMessageHandler func(msg []byte)

// WsMessage represents a WebSocket message
type WsMessage[D any] struct {
	Stream string `json:"stream"`
	Data   D      `json:"data,omitempty"`
}

type WsOrderData struct {
	EventType              string `json:"e,omitempty"`
	EventTimeMicroSeconds  int64  `json:"E,omitempty"`
	Symbol                 string `json:"s,omitempty"`
	ClientOrderId          uint64 `json:"c,omitempty"`
	Side                   string `json:"S,omitempty"`
	OrderType              string `json:"o,omitempty"`
	TimeInForce            string `json:"f,omitempty"`
	Quantity               string `json:"q,omitempty"`
	QuoteQuantity          string `json:"Q,omitempty"`
	Price                  string `json:"p,omitempty"`
	TriggerPrice           string `json:"P,omitempty"`
	TriggerBy              string `json:"B,omitempty"`
	TakeProfitTriggerPrice string `json:"a,omitempty"`
	StopLossTriggerPrice   string `json:"b,omitempty"`
	TakeProfitTriggerBy    string `json:"d,omitempty"`
	StopLossTriggerBy      string `json:"g,omitempty"`
	TriggerQuantity        string `json:"Y,omitempty"`
	OrderState             string `json:"X,omitempty"`
	OrderExpiryReason      string `json:"R,omitempty"`
	OrderId                string `json:"i,omitempty"`
	TradeId                int64  `json:"t,omitempty"`
	FileQuantity           string `json:"l,omitempty"`
	ExecutedQuantity       string `json:"z,omitempty"`
	ExecutedQuoteQuantity  string `json:"Z,omitempty"`
	FilePrice              string `json:"L,omitempty"`
	IsMaker                bool   `json:"m,omitempty"`
	Fee                    string `json:"n,omitempty"`
	FeeSymbol              string `json:"N,omitempty"`
	SelfTradePrevention    string `json:"V,omitempty"`
	EngineTimeMicroSeconds int64  `json:"T,omitempty"`
	Origin                 string `json:"O,omitempty"`
	RelatedOrderId         string `json:"I,omitempty"`
	StrategyId             int64  `json:"H,omitempty"`
	PostOnly               bool   `json:"y,omitempty"`
	ReduceOnly             bool   `json:"r,omitempty"`
}

type WsOrderDataHandler func(data WsOrderData)

type WsPositionData struct {
	EventType                 string `json:"e,omitempty"`
	EventTimeMicroSeconds     int64  `json:"E,omitempty"`
	Symbol                    string `json:"s,omitempty"`
	BreakEventPrice           string `json:"b,omitempty"`
	EntryPrice                string `json:"B,omitempty"`
	InitialMarginFraction     string `json:"f,omitempty"`
	MarkPrice                 string `json:"M,omitempty"`
	MaintenanceMarginFraction string `json:"m,omitempty"`
	NetQuantity               string `json:"q,omitempty"`
	NetExposureQuantity       string `json:"Q,omitempty"`
	NetExposureNotional       string `json:"n,omitempty"`
	PositionId                int64  `json:"i,omitempty"`
	PnlRealized               string `json:"p,omitempty"`
	PnlUnrealized             string `json:"P,omitempty"`
	EngineTimeMicroSeconds    int64  `json:"T,omitempty"`
}

type WsPositionDataHandler func(data WsPositionData)

type WsBookTickerData struct {
	EventType              string `json:"e,omitempty"`
	EventTimeMicroSeconds  int64  `json:"E,omitempty"`
	Symbol                 string `json:"s,omitempty"`
	InsideAskPrice         string `json:"a,omitempty"`
	InsideAskQuantity      string `json:"A,omitempty"`
	InsideBidPrice         string `json:"b,omitempty"`
	InsideBidQuantity      string `json:"B,omitempty"`
	EventUpdateId          int64  `json:"u,omitempty"`
	EngineTimeMicroSeconds int64  `json:"T,omitempty"`
}

type WsBookTickerDataHandler func(data WsBookTickerData)

type WsDepthData struct {
	EventType              string     `json:"e,omitempty"`
	EventTimeMicroSeconds  int64      `json:"E,omitempty"`
	Symbol                 string     `json:"s,omitempty"`
	Asks                   [][]string `json:"a,omitempty"`
	Bids                   [][]string `json:"b,omitempty"`
	FirstUpdateId          int64      `json:"U,omitempty"`
	LastUpdateId           int64      `json:"u,omitempty"`
	EngineTimeMicroSeconds int64      `json:"T,omitempty"`
}

type WsDepthDataHandler func(data WsDepthData)

type WsKlineData struct {
	EventType             string `json:"e,omitempty"`
	EventTimeMicroSeconds int64  `json:"E,omitempty"`
	Symbol                string `json:"s,omitempty"`
	StartTime             string `json:"t,omitempty"`
	CloseTime             string `json:"T,omitempty"`
	OpenPrice             string `json:"o,omitempty"`
	ClosePrice            string `json:"c,omitempty"`
	HighPrice             string `json:"h,omitempty"`
	LowPrice              string `json:"l,omitempty"`
	BaseAssetVolume       string `json:"v,omitempty"`
	NumberOfTrades        int64  `json:"n,omitempty"`
	IsClosed              bool   `json:"X,omitempty"`
}

type WsKlineDataHandler func(data WsKlineData)

type WsLiquidationData struct {
	EventType              string `json:"e,omitempty"`
	EventTimeMicroSeconds  int64  `json:"E,omitempty"`
	Quantity               string `json:"q,omitempty"`
	Price                  string `json:"p,omitempty"`
	Side                   string `json:"S,omitempty"`
	Symbol                 string `json:"s,omitempty"`
	EngineTimeMicroSeconds int64  `json:"T,omitempty"`
}

type WsLiquidationDataHandler func(data WsLiquidationData)

type WsMarkPriceData struct {
	EventType                   string `json:"e,omitempty"`
	EventTimeMicroSeconds       int64  `json:"E,omitempty"`
	Symbol                      string `json:"s,omitempty"`
	MarkPrice                   string `json:"p,omitempty"`
	EstimatedFundingRate        string `json:"f,omitempty"`
	IndexPrice                  string `json:"i,omitempty"`
	NextFundingTimeMicroSeconds int64  `json:"n,omitempty"`
}

type WsMarkPriceDataHandler func(data WsMarkPriceData)

type WsTickerData struct {
	EventType             string `json:"e,omitempty"`
	EventTimeMicroSeconds int64  `json:"E,omitempty"`
	Symbol                string `json:"s,omitempty"`
	FirstPrice            string `json:"o,omitempty"`
	LastPrice             string `json:"c,omitempty"`
	HighPrice             string `json:"h,omitempty"`
	LowPrice              string `json:"l,omitempty"`
	BaseAssetVolume       string `json:"v,omitempty"`
	QuoteAssetVolume      string `json:"V,omitempty"`
	NumberOfTrades        int64  `json:"n,omitempty"`
}

type WsTickerDataHandler func(data WsTickerData)

type WsOpenInterestData struct {
	EventType             string `json:"e,omitempty"`
	EventTimeMicroSeconds int64  `json:"E,omitempty"`
	Symbol                string `json:"s,omitempty"`
	OpenInterest          string `json:"o,omitempty"`
}

type WsOpenInterestDataHandler func(data WsOpenInterestData)

type WsTradeData struct {
	EventType              string `json:"e,omitempty"`
	EventTimeMicroSeconds  int64  `json:"E,omitempty"`
	Symbol                 string `json:"s,omitempty"`
	Price                  string `json:"p,omitempty"`
	Quantity               string `json:"q,omitempty"`
	BuyerOrderId           string `json:"b,omitempty"`
	SellerOrderId          string `json:"a,omitempty"`
	TradeId                int64  `json:"t,omitempty"`
	EngineTimeMicroSeconds int64  `json:"T,omitempty"`
	IsBuyerMaker           bool   `json:"m,omitempty"`
}

type WsTradeDataHandler func(data WsTradeData)

func (ws *WsClient) setStatus(status int) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.status = status
}

func (ws *WsClient) getStatus() int {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.status
}

func (ws *WsClient) setConn(conn *websocket.Conn) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.done = make(chan struct{})
	ws.ws = conn
	ws.status = WsStatusConnected
}

func (ws *WsClient) getConn() *websocket.Conn {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.ws
}

func (ws *WsClient) sign() (sig []string) {
	ms := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	sigStr := fmt.Sprintf("instruction=subscribe&timestamp=%s&window=%s", ms, ws.Window)
	apiSecret, _ := base64.StdEncoding.DecodeString(strings.TrimSpace(ws.Secret))
	pki := ed25519.NewKeyFromSeed(apiSecret)
	signature := base64.StdEncoding.EncodeToString(ed25519.Sign(pki, []byte(sigStr)))
	verifyingKey := ws.Key
	sig = []string{
		verifyingKey,
		signature,
		ms,
		ws.Window,
	}
	return
}

func (ws *WsClient) pushErr(err error) {
	select {
	case ws.errChan <- err:
	default:
	}
}

func (ws *WsClient) writeSubscribeMsg(stream string, isPrivate bool) error {
	subMsg := WsSubscribeMessage{
		Method:    "SUBSCRIBE",
		Params:    []string{stream},
		Signature: nil,
	}
	if isPrivate {
		subMsg.Signature = ws.sign()
	}

	return ws.ws.WriteJSON(subMsg)
}

func (ws *WsClient) resubscribe() error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	f := func(stream string) error {
		var isPrivate bool
		if strings.HasPrefix(stream, "account.") {
			isPrivate = true
		}
		if err := ws.writeSubscribeMsg(stream, isPrivate); err != nil {
			return err
		}
		return nil
	}

	for stream := range ws.handlers {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.orderSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.positionSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.bookTickerSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.depthSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.klineSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.liquidationSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.markPriceSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.tickerSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.tradeSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	for stream := range ws.openInterestSub {
		if err := f(stream); err != nil {
			return err
		}
	}

	return nil
}

func (ws *WsClient) tryReconnect(reason error) {
	ws.Disconnect(reason)
	if !ws.autoReconnect {
		return
	}
	if err := ws.Connect(); err != nil {
		ws.pushErr(fmt.Errorf("reconnect error: %w", err))
	} else {
		return
	}
	for {
		select {
		case <-ws.ctx.Done():
			return
		case <-time.After(time.Second):
			if err := ws.Connect(); err != nil {
				ws.pushErr(fmt.Errorf("reconnect error: %w", err))
				continue
			}
			return
		}
	}
}

func (ws *WsClient) keepAlive() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ws.ctx.Done():
			ws.Disconnect(context.Canceled)
			return
		case <-ws.done:
			return
		case <-ticker.C:
			if ws.getStatus() == WsStatusConnected {
				if err := ws.getConn().WriteMessage(websocket.PingMessage, []byte("Ping")); err != nil {
					return
				}
			}
		}
	}
}

func (ws *WsClient) handleMessages() {
	for {
		select {
		case <-ws.ctx.Done():
			return
		case <-ws.done:
			return
		default:
			// go on
		}

		conn := ws.getConn()
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			ws.tryReconnect(err)
			return
		}

		switch msgType {
		case websocket.PingMessage:
			if err := conn.WriteMessage(websocket.PongMessage, []byte("Pong")); err != nil {
				ws.pushErr(fmt.Errorf("write pong error: %w", err))
			}
			continue
		case websocket.PongMessage:
			continue
		case websocket.CloseMessage:
			ws.tryReconnect(errors.New("closed by server"))
		default:
			for _, hook := range ws.onMessageHooks {
				hook(msg)
			}
		}

		var wsMsg WsMessage[json.RawMessage]
		if err := json.Unmarshal(msg, &wsMsg); err != nil {
			ws.pushErr(fmt.Errorf("unmarshal ws message error: %w", err))
			continue
		}

		if msgHandlers, ok := ws.handlers[wsMsg.Stream]; ok {
			for _, msgHandler := range msgHandlers {
				msgHandler(wsMsg.Data)
			}
		}

		streamSplit := strings.Split(wsMsg.Stream, ".")
		streamType := streamSplit[0]
		switch streamType {
		case "account":
			switch streamSplit[1] {
			case "orderUpdate":
				ws.mu.RLock()
				handlers, ok := ws.orderSub[wsMsg.Stream]
				ws.mu.RUnlock()
				if !ok {
					continue
				}
				var data WsOrderData
				if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
					ws.pushErr(fmt.Errorf("unmarshal ws order data message error: %w", err))
					continue
				}
				for _, handler := range handlers {
					handler(data)
				}
			case "positionUpdate":
				ws.mu.RLock()
				handlers, ok := ws.positionSub[wsMsg.Stream]
				ws.mu.RUnlock()
				if !ok {
					continue
				}
				var data WsPositionData
				if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
					ws.pushErr(fmt.Errorf("unmarshal ws position data message error: %w", err))
					continue
				}
				for _, handler := range handlers {
					handler(data)
				}
			case "rfqUpdate":
				// todo
			default:
				// ignore
			}
			continue
		case "bookTicker":
			ws.mu.RLock()
			handlers, ok := ws.bookTickerSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsBookTickerData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws book ticker data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		case "depth":
			ws.mu.RLock()
			handlers, ok := ws.depthSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsDepthData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws depth data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		case "kline":
			ws.mu.RLock()
			handlers, ok := ws.klineSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsKlineData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws kline data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		case "liquidation":
			ws.mu.RLock()
			handlers, ok := ws.liquidationSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsLiquidationData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws liquidation data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		case "markPrice":
			ws.mu.RLock()
			handlers, ok := ws.markPriceSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsMarkPriceData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws mark price data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		case "ticker":
			ws.mu.RLock()
			handlers, ok := ws.tickerSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsTickerData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws ticker data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		case "openInterest":
			ws.mu.RLock()
			handlers, ok := ws.openInterestSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsOpenInterestData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws open interest data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		case "trade":
			ws.mu.RLock()
			handlers, ok := ws.tradeSub[wsMsg.Stream]
			ws.mu.RUnlock()
			if !ok {
				continue
			}
			var data WsTradeData
			if err := json.Unmarshal(wsMsg.Data, &data); err != nil {
				ws.pushErr(fmt.Errorf("unmarshal ws trade data message error: %w", err))
				continue
			}
			for _, handler := range handlers {
				handler(data)
			}
		default:
			ws.pushErr(fmt.Errorf("unknown stream: %s", wsMsg.Stream))
		}
		continue
	}
}
