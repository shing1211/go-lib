package binance

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2"
)

func GetAllSymbolList(binanceClient *binance.Client) ([]string, []float64) {
	// Get Exchange information
	fmt.Println("Get all binance Symbol List")
	symbolList := []string{}
	lotsizeList := []float64{}
	res, err := binanceClient.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	for i := range res.Symbols {
		symbolList = append(symbolList, res.Symbols[i].Symbol)

		lotsize, err := strconv.ParseFloat(res.Symbols[i].LotSizeFilter().MinQuantity, 64)
		if err != nil {
			fmt.Println(err)
			return nil, nil
		}
		lotsizeList = append(lotsizeList, lotsize)
	}
	return symbolList, lotsizeList
}

func GetAllUSDTSymbolList(binanceClient *binance.Client) ([]string, []float64) {
	// Get Exchange information
	fmt.Println("Get all Tradeble USDT based Symbol List")
	symbolList := []string{}
	lotsizeList := []float64{}
	res, err := binanceClient.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	for i := range res.Symbols {
		if res.Symbols[i].Symbol != "USDCUSDT" && strings.HasSuffix(res.Symbols[i].Symbol, "USDT") && !strings.HasSuffix(res.Symbols[i].Symbol, "UPUSDT") && !strings.HasSuffix(res.Symbols[i].Symbol, "DOWNUSDT") {
			symbolList = append(symbolList, res.Symbols[i].Symbol)

			lotsize, err := strconv.ParseFloat(res.Symbols[i].LotSizeFilter().MinQuantity, 64)
			if err != nil {
				fmt.Println(err)
				return nil, nil
			}
			lotsizeList = append(lotsizeList, lotsize)
		}
	}
	return symbolList, lotsizeList
}
