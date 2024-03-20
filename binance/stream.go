package binance

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
)

// WsUserDataStream serve user data handler with listen key
func WsUserDataStream(listenKey string) {
	wsUserDataHandler := func(event *binance.WsUserDataEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsUserDataServe(listenKey, wsUserDataHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func StartUserStream(binanceClient *binance.Client) string {
	fmt.Println("Start User stream")
	key, err := binanceClient.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(key)
	return key
}

func KeepAliveUserStream(binanceClient *binance.Client, listenKey string) {
	fmt.Println("Keep Alive User stream")
	err := binanceClient.NewKeepaliveUserStreamService().ListenKey(listenKey).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CloseUserStream(binanceClient *binance.Client, listenKey string) {
	fmt.Println("Close User stream")
	err := binanceClient.NewCloseUserStreamService().ListenKey(listenKey).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
}
