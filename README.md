# Backpack Exchange API Client (Go)

[[中文](./README_CN.md)]

Go SDK for the Backpack Exchange, a cryptocurrency exchange platform.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [API Documentation](#api-documentation)
- [REST API](#rest-api)
- [WebSocket API](#websocket-api)
- [Features](#features)
- [License](#license)

## Installation

```bash
go get github.com/rambollwong/bpx-api-client-go
```

## Quick Start

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
    // Create REST API client
    client := bpx.NewClient("your_api_key", "your_api_secret")

    // Get account information
    accountInfo, err := client.Account().GetAccount(types.GetAccountReq{})
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Account Info: %+v\n", accountInfo)
}
```

## API Documentation

Official API Documentation: https://docs.backpack.exchange/

## REST API

### Initialize Client

```go
client := bpx.NewClient("your_api_key", "your_api_secret")
```

### Supported Modules

- [x] Account
- [x] Assets
- [x] Borrow/Lend Markets
- [x] Borrow/Lend
- [x] Capital
- [x] Futures
- [x] History
- [x] Markets
- [x] Order
- [x] System
- [x] Trades

### Usage Examples

Get Markets:
```go
markets, err := client.Markets().GetMarkets()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Markets: %+v\n", markets)
```

Place Order:
```go
quantity := "5"
postOnly := true
price := "0.25070"
resp, err := client.Order().ExecuteOrder(types.ExecuteOrderReq{
    OrderType: "Limit",
    Symbol:    "DOGE_USDC_PERP",
    Side:      "Bid",
    Quantity:  &quantity,
    PostOnly:  &postOnly,
    Price:     &price,
})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Order Response: %+v\n", resp)
```

Get Account Balances:
```go
balances, err := client.Account().GetAccount(types.GetAccountReq{})
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Account Balances: %+v\n", balances)
```

## WebSocket API

### Initialize WebSocket Client

```go
wsClient := bpx.NewWsClient("your_api_key", "your_api_secret")
```

### Connect

```go
err := wsClient.Connect()
if err != nil {
    log.Fatal(err)
}
defer wsClient.Disconnect(nil)
```

### Subscribe to Streams

Subscribe to Order Updates:
```go
err = wsClient.SubscribeOrderUpdate("", func(data bpx.WsOrderData) {
    fmt.Printf("Order Update: %+v\n", data)
})
if err != nil {
    log.Fatal(err)
}
```

Subscribe to Depth Data:
```go
err = wsClient.SubscribeDepth("SOL_USDC", "", func(data bpx.WsDepthData) {
    fmt.Printf("Depth Data: %+v\n", data)
})
if err != nil {
    log.Fatal(err)
}
```

Supported Stream Types:
- Order Update
- Position Update
- Book Ticker
- Depth
- Kline
- Liquidation
- Mark Price
- Ticker
- Open Interest
- Trade

## Features

| Feature | Status |
|---------|--------|
| REST API Calls | ✅ Completed |
| WebSocket Connection | ✅ Completed |
| Auto Reconnection | ✅ Completed |
| Signature Verification | ✅ Completed |
| Error Handling | ✅ Completed |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.