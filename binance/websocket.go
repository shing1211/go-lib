package binance

import (
	"fmt"

	"github.com/adshao/go-binance/v2"
)

// WsCombinedAggTradeStreams is similar to WsAggTradeServe, but it handles multiple symbolx
func WsCombinedAggTradeStreams(symbols []string) {
	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsCombinedAggTradeServe(symbols, wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsAggTradeStreams serve websocket aggregate handler with a symbol
func WsAggTradeStreams(symbol string) {
	wsAggTradeHandler := func(event *binance.WsAggTradeEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAggTradeServe(symbol, wsAggTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsTradeStreams serve websocket handler with a symbol
func WsTradeStreams(symbol string) {
	wsTradeHandler := func(event *binance.WsTradeEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsTradeServe(symbol, wsTradeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsCombinedPriceDepthStreams is similar to WsDepthServe, but it for multiple symbols
func WsCombinedPriceDepthStreams(symbols []string) {
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsCombinedDepthServe(symbols, wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsPriceDepthStreams serve websocket depth handler with an arbitrary endpoint address
func WsPriceDepthStreams(symbol string) {
	wsDepthHandler := func(event *binance.WsDepthEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsDepthServe(symbol, wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsKlineStream serve websocket kline handler with a symbol and interval like 15m, 30s
func WsKlineStream(symbol string, interval string) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsKlineServe(symbol, interval, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsCombinedKlineStream is similar to WsKlineServe, but it handles multiple symbols with it interval
func WsCombinedKlineStream(symbolIntervalPair map[string]string) {
	wsKlineHandler := func(event *binance.WsKlineEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsCombinedKlineServe(symbolIntervalPair, wsKlineHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsBookTickerStreams serve websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
func WsBookTickerStreams(symbol string) {
	wsBookTickerHandler := func(event *binance.WsBookTickerEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsBookTickerServe(symbol, wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsAllBookTickerStreams serve websocket that pushes updates to the best bid or ask price or quantity in real-time for all symbols.
func WsAllBookTickerStreams() {
	wsBookTickerHandler := func(event *binance.WsBookTickerEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAllBookTickerServe(wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsMarketStatStream serve websocket that push 24hr statistics for single market every second
func WsMarketStatStream(symbol string) {
	wsMarketStatHandler := func(event *binance.WsMarketStatEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsMarketStatServe(symbol, wsMarketStatHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

// WsCombinedMarketStatStream is similar to WsMarketStatServe, but it handles multiple symbol
func WsCombinedMarketStatStream(symbol []string) {
	wsMarketStatHandler := func(event *binance.WsMarketStatEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsCombinedMarketStatServe(symbol, wsMarketStatHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

//WsAllMarketsStatsStream serve websocket that push 24hr statistics for all market every second
/* func WsAllMarketsStatsStream() {
	wsAllMarketsStatServeHandler := func(event *binance.WsAllMarketsStatEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAllMarketsStatServe(wsAllMarketsStatServeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
} */

// WsAllMiniMarketsStatsStream serve websocket that push mini version of 24hr statistics for all market every second
/* func WsAllMiniMarketsStatsStream() {
	wsAllMiniMarketsStatServeHandler := func(event *binance.WsAllMiniMarketsStatEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAllMiniMarketsStatServe(wsAllMiniMarketsStatServeHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
} */
