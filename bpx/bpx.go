package bpx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"
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

	ctx context.Context
}

func NewClientWithContext(ctx context.Context, key, secret string) *Client {
	cli := &Client{
		Key:    key,
		Secret: secret,
		Window: "60000",
		ctx:    ctx,
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

func NewClient(key, secret string) *Client {
	return NewClientWithContext(context.Background(), key, secret)
}

func (c *Client) SetProxy(proxy string) error {
	c.Proxy = proxy

	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return fmt.Errorf("parse proxy url error: %w", err)
	}
	c.httpCli.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyUrl)

	return nil
}

func (c *Client) Asserts() Assets {
	return Assets{c}
}

func (c *Client) BorrowLendMarkets() BorrowLendMarkets {
	return BorrowLendMarkets{c}
}

func (c *Client) Markets() Markets {
	return Markets{c}
}

func (c *Client) System() System {
	return System{c}
}

func (c *Client) Trades() Trades {
	return Trades{c}
}

func (c *Client) Account() Account {
	return Account{c}
}

func (c *Client) BorrowLend() BorrowLend {
	return BorrowLend{c}
}

func (c *Client) Capital() Capital {
	return Capital{c}
}

func (c *Client) Futures() Futures {
	return Futures{c}
}

func (c *Client) History() History {
	return History{c}
}

func (c *Client) Order() Order {
	return Order{c}
}
