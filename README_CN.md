# Backpack Exchange API 客户端 (Go语言)

[[English](./README.md)]

Backpack Exchange 的 Go 语言 SDK，用于与 Backpack Exchange 交易所进行交互。

## 目录

- [安装](#安装)
- [快速开始](#快速开始)
- [API 文档](#api文档)
- [REST API](#rest-api)
- [WebSocket API](#websocket-api)
- [功能列表](#功能列表)
- [许可证](#许可证)

## 安装

```bash
go get github.com/rambollwong/bpx-api-client-go
```

## 快速开始

```go
package main

import (
    "context"
    "fmt"
    "log"

	"bpx-api-client-go/bpx"
	"bpx-api-client-go/types"
)

func main() {
    // 创建 REST API 客户端
    client := bpx.NewClient("your_api_key", "your_api_secret")

    // 获取账户信息
    accountInfo, err := client.Account().GetAccount(types.GetAccountReq{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("账户信息: %+v\n", accountInfo)
}
```

## API 文档

官方 API 文档: https://docs.backpack.exchange/

## REST API

### 初始化客户端

```go
client := bpx.NewClient("your_api_key", "your_api_secret")
```

### 支持的模块

- [x] 账户 (Account)
- [x] 资产 (Assets)
- [x] 借贷市场 (Borrow/Lend Markets)
- [x] 借贷 (Borrow/Lend)
- [x] 资金 (Capital)
- [x] 合约 (Futures)
- [x] 历史记录 (History)
- [x] 市场数据 (Markets)
- [x] 订单 (Order)
- [x] 系统 (System)
- [x] 交易 (Trades)

### 使用示例

获取市场数据:

```go
markets, err := client.Markets().GetMarkets()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("市场数据: %+v\n", markets)
```

下单:
```go
quantity := "5"
postOnly := true
price := "0.25070"
resp, err := c.ExecuteOrder(types.ExecuteOrderReq{
    OrderType: "Limit",
    Symbol:    "DOGE_USDC_PERP",
    Side:      types.SideBid,
    Quantity:  &quantity,
    PostOnly:  &postOnly,
    Price:     &price,
}
if err != nil {
    log.Fatal(err)
}
fmt.Printf("订单响应: %+v\n", orderResp)
```

## WebSocket API

### 初始化 WebSocket 客户端

```go
wsClient := bpx.NewWsClient("your_api_key", "your_api_secret")
```

### 连接

```go
err := wsClient.Connect()
if err != nil {
    log.Fatal(err)
}
defer wsClient.Disconnect(nil)
```

### 订阅数据流

订阅订单更新:
```go
err = wsClient.SubscribeOrderUpdate("", func(data bpx.WsOrderData) {
    fmt.Printf("订单更新: %+v\n", data)
})
if err != nil {
    log.Fatal(err)
}
```

订阅深度数据:
```go
err = wsClient.SubscribeDepth("SOL_USDC", "", func(data bpx.WsDepthData) {
    fmt.Printf("深度数据: %+v\n", data)
})
if err != nil {
    log.Fatal(err)
}
```

支持的数据流类型:
- 订单更新 (Order Update)
- 持仓更新 (Position Update)
- 最优挂单 (Book Ticker)
- 深度数据 (Depth)
- K线数据 (Kline)
- 强制平仓 (Liquidation)
- 标记价格 (Mark Price)
- 行情数据 (Ticker)
- 持仓量 (Open Interest)
- 交易数据 (Trade)

## 功能列表

| 功能 | 状态 |
|------|------|
| REST API 调用 | ✅ 已完成 |
| WebSocket 连接 | ✅ 已完成 |
| 自动重连 | ✅ 已完成 |
| 签名验证 | ✅ 已完成 |
| 错误处理 | ✅ 已完成 |

## 许可证

本项目采用 MIT 许可证。详情请见 [LICENSE](LICENSE) 文件。