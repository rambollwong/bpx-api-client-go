package bpx

import (
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"bpx-api-client-go/types"
)

func (c *Client) SetProxy(proxy string) error {
	c.Proxy = proxy

	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return fmt.Errorf("parse proxy url error: %w", err)
	}
	c.httpCli.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyUrl)

	return nil
}
func (c *Client) SetHttpTimeout(timeout time.Duration) {
	c.httpCli.Timeout = timeout
}

func (c *Client) SetHttpDisableCompression(disable bool) {
	c.httpCli.Transport.(*http.Transport).DisableCompression = disable
}

func (c *Client) SetHttpDisableKeepAlives(disable bool) {
	c.httpCli.Transport.(*http.Transport).DisableKeepAlives = disable
}

func (c *Client) SetHttpIdleConnTimeout(timeout time.Duration) {
	c.httpCli.Transport.(*http.Transport).IdleConnTimeout = timeout
}

func (c *Client) SetHttpResponseHeaderTimeout(timeout time.Duration) {
	c.httpCli.Transport.(*http.Transport).ResponseHeaderTimeout = timeout
}

func (c *Client) SetHttpMaxIdleConns(maxIdleConns int) {
	c.httpCli.Transport.(*http.Transport).MaxIdleConns = maxIdleConns
}

func DoRequest[B, R any](
	ctx context.Context,
	method, url string,
	body B,
	c *Client) (resp R, rh *types.ResponseHeaders, err error) {
	ms := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)

	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return resp, nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	switch ins := any(body).(type) {
	case types.AuthenticatedRequest:
		if err := ins.Validate(); err != nil {
			return resp, nil, err
		}
		req.Header.Set("X-Window", c.Window)
		req.Header.Set("X-API-Key", c.Key)
		req.Header.Set("X-Timestamp", ms)
		// 计算签名
		queryParams := ins.BuildQueryParams()
		sig := Sign(c, ins.Instruction(), ms, queryParams)
		req.Header.Set("X-Signature", sig)
		if method == http.MethodGet {
			req.URL.RawQuery = queryParams
		} else {
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return resp, nil, err
			}
			req.Body = io.NopCloser(strings.NewReader(string(jsonBody)))
		}
	case types.QueryParamsBuilder:
		if err := ins.Validate(); err != nil {
			return resp, nil, err
		}
		req.URL.RawQuery = ins.BuildQueryParams()
	default:
		// ignore
	}

	// 发送请求并处理响应
	res, err := c.httpCli.Do(req)
	if err != nil {
		return resp, nil, err
	}
	defer res.Body.Close()

	bz, err := io.ReadAll(res.Body)
	if err != nil {
		return resp, nil, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		// TODO：这里使用了类型反射，如果对性能有要求，需要改为使用类型确定的handler定义方式来处理
		switch r := any(resp).(type) {
		case types.ResponseHeadersReader:
			err = json.Unmarshal(bz, &resp)
			if err != nil {
				return resp, nil, err
			}
			rh, err = r.ReadResponseHeaders(res.Header)
			if err != nil {
				return resp, nil, err
			}
		case string:
			resp = any(string(bz)).(R)
		case *string:
			s := string(bz)
			resp = any(&s).(R)
		case struct{}:
			// ignore
		default:
			err = json.Unmarshal(bz, &resp)
			if err != nil {
				return resp, nil, err
			}
		}
	case http.StatusNoContent:
		return resp, nil, types.WrapCodeMessageError(types.CodeMessage{Code: "204", Message: "Not found"})
	//case http.StatusBadRequest:
	//	return resp, types.WrapCodeMessageError(types.CodeMessage{Code: "400", Message: fmt.Sprintf("Bad request: %s", string(bz))})
	default:
		cm := &types.CodeMessage{}
		err := json.Unmarshal(bz, cm)
		if err != nil {
			return resp, nil, err
		}
		return resp, nil, types.WrapCodeMessageError(*cm)
	}

	return resp, rh, err
}

func Sign(c *Client, instruction, ms, queryString string) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("instruction=")
	strBuilder.WriteString(instruction)
	if len(queryString) > 0 {
		strBuilder.WriteString("&")
		strBuilder.WriteString(queryString)
	}
	strBuilder.WriteString("&timestamp=")
	strBuilder.WriteString(ms)
	strBuilder.WriteString("&window=")
	strBuilder.WriteString(c.Window)

	apiSecret, _ := base64.StdEncoding.DecodeString(strings.TrimSpace(c.Secret))

	pki := ed25519.NewKeyFromSeed(apiSecret)
	return base64.StdEncoding.EncodeToString(ed25519.Sign(pki, []byte(strBuilder.String())))
}
