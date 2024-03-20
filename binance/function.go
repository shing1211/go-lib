package binance

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/adshao/go-binance/v2"
	log "github.com/sirupsen/logrus"
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
	symbolList := []string{}
	lotsizeList := []float64{}
	res, err := binanceClient.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		log.Warn(err)
		return nil, nil
	}
	for i := range res.Symbols {
		if res.Symbols[i].Symbol != "USDCUSDT" && strings.HasSuffix(res.Symbols[i].Symbol, "USDT") && !strings.HasSuffix(res.Symbols[i].Symbol, "UPUSDT") && !strings.HasSuffix(res.Symbols[i].Symbol, "DOWNUSDT") {
			symbolList = append(symbolList, res.Symbols[i].Symbol)
			lotsize, err := strconv.ParseFloat(res.Symbols[i].LotSizeFilter().MinQuantity, 64)
			if err != nil {
				log.Warn(err)
				return nil, nil
			}
			lotsizeList = append(lotsizeList, lotsize)
		}
	}
	return symbolList, lotsizeList
}

func ListRecentTradeBySymbol(binanceClient *binance.Client, symbol string) {
	// List Aggregate Trades
	trades, err := binanceClient.NewRecentTradesService().Symbol(symbol).Do(context.Background())

	fmt.Println("")
	fmt.Println("List Recent Trades")

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, t := range trades {
		fmt.Println(t)
	}
}

func ListAggregateTradeBySymbol(binanceClient *binance.Client, symbol string) {
	// List Aggregate Trades
	trades, err := binanceClient.NewAggTradesService().Symbol(symbol).StartTime(1634366367000).EndTime(1634369967000).Do(context.Background())

	fmt.Println("")
	fmt.Println("List Aggregate Trades")

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, t := range trades {
		fmt.Println(t)
	}
}

func ShowPriceDepthBySymbol(binanceClient *binance.Client, symbol string) {
	// Show Price Depth
	fmt.Println("")
	fmt.Println("Show Price Depth for:", symbol)
	res, err := binanceClient.NewDepthService().Symbol(symbol).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func ListKlinesBySymbolByInterval(binanceClient *binance.Client, symbol string, interval string) {
	// List Klines
	fmt.Println("")
	fmt.Println("List Klines for:", symbol, interval)
	klines, err := binanceClient.NewKlinesService().Symbol(symbol).Interval(interval).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, k := range klines {
		fmt.Println(k)
	}
}

func ListTicketPricesChangeStatsBySymbol(binanceClient *binance.Client, symbol string) {
	// List Ticker Prices
	fmt.Println("")
	fmt.Println("List Ticker Price change Stats for:", symbol)
	prices, err := binanceClient.NewListPriceChangeStatsService().Symbol(symbol).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}
}

func GetSymbolPrice(binanceClient *binance.Client, symbol string) float64 {
	// Get Symbol Price
	prices, err := binanceClient.NewListPricesService().Symbol(symbol).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return 0
	}

	for _, p := range prices {
		price, err := strconv.ParseFloat(p.Price, 64)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		return price
	}
	return 0
}

func GetBookTicketPricesBySymbol(binanceClient *binance.Client, symbol string) {
	// List Ticker Prices
	fmt.Println("List Book Ticker Price for:", symbol)
	prices, err := binanceClient.NewListBookTickersService().Symbol(symbol).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, p := range prices {
		fmt.Println(p)
	}
}

func ListAllTicketPrices(binanceClient *binance.Client) {
	// List Ticker Prices
	fmt.Println("")
	fmt.Println("List All Ticker Price for:")
	prices, err := binanceClient.NewListPricesService().Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}
}

func ListAllDeposit(binanceClient *binance.Client) {
	fmt.Println("")
	fmt.Println("List All Deposit for:")
	// List Orders
	orders, err := binanceClient.NewListDepositsService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}
}

func ListDepositAddress(binanceClient *binance.Client) {
	fmt.Println("")
	fmt.Println("List All Deposit Address:")
	// List Orders
	orders, err := binanceClient.NewGetDepositAddressService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(orders)
}

func ListAllMktFilledTradesBySymbol(binanceClient *binance.Client, symbol string) {
	fmt.Println("")
	fmt.Println("List All Market Filled Trades for:", symbol)
	// List Orders
	orders, err := binanceClient.NewHistoricalTradesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}
}

func ListAllTradesBySymbol(binanceClient *binance.Client, symbol string) {
	fmt.Println("")
	fmt.Println("List All Trades for:", symbol)
	// List Orders
	orders, err := binanceClient.NewListTradesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}
}

func ListAllOrdersBySymbol(binanceClient *binance.Client, symbol string) {
	fmt.Println("")
	fmt.Println("List All History Order for:", symbol)
	// List Orders
	orders, err := binanceClient.NewListOrdersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range orders {
		fmt.Println(o)
	}
}

func ListAllOpenOrdersBySymbol(binanceClient *binance.Client, symbol string) {
	// List Open Orders
	fmt.Println("")
	fmt.Println("List All Open Order for:", symbol)
	openOrders, err := binanceClient.NewListOpenOrdersService().Symbol(symbol).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, o := range openOrders {
		fmt.Println(o)
	}
}

func ListAllAssetDetails(binanceClient *binance.Client) {
	// List Open Orders
	fmt.Println("")
	fmt.Println("List Asset Details:")
	openOrders, err := binanceClient.NewGetAssetDetailService().Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, o := range openOrders {
		fmt.Println(o)
	}
}

func CreateOrder(binanceClient *binance.Client, symbol string, sidetype binance.SideType, ordertype binance.OrderType, quantity string) (*binance.CreateOrderResponse, error) {
	// Create order
	fmt.Println("Create Order for:", symbol)
	order, err := binanceClient.NewCreateOrderService().
		Symbol(symbol).
		Side(sidetype).
		Type(ordertype).
		//TimeInForce(binance.TimeInForceTypeGTC).
		Quantity(quantity).
		//Price("5.000000").
		Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(order)
	return order, nil
}

func ListOrderbyID(binanceClient *binance.Client, orderID int64, symbol string) {
	fmt.Println("")
	fmt.Println("Order information of OrderID:", orderID)
	order, err := binanceClient.NewGetOrderService().Symbol(symbol).OrderID(orderID).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
	// fmt.Println(order.Symbol)
	// fmt.Println(order.OrderID)
	// fmt.Println(order.OrderListId)
	// fmt.Println(order.ClientOrderID)
	// fmt.Println(order.Price)
	// fmt.Println(order.OrigQuantity)
	// fmt.Println(order.ExecutedQuantity)
	// fmt.Println(order.IcebergQuantity)
	// fmt.Println(order.Status)
	// fmt.Println(order.TimeInForce)
	// fmt.Println(order.Type)
	// fmt.Println(order.Side)
	// fmt.Println(order.CummulativeQuoteQuantity)
	// fmt.Println(order.StopPrice)
	// fmt.Println(order.Time)
	// fmt.Println(order.UpdateTime)
	// fmt.Println(order.IsWorking)
	// fmt.Println(order.IsIsolated)
}

func CancelOrderByID(binanceClient *binance.Client, orderID int64, symbol string) {
	// Cancel order
	fmt.Println("")
	fmt.Println("Cancel Order for:", orderID)
	//_, err := client.NewCancelOrderService().Symbol(symbol).OrderID(orderID).Do(context.Background())
	//cancelOrder, err := client.NewCancelOrderService().Symbol(symbol).OrderID(orderID).Do(context.Background())
	cancelOrder, err := binanceClient.NewCancelOrderService().Symbol(symbol).OrderID(orderID).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cancelOrder)
}

func CancelAllOpenOrderBySymbol(binanceClient *binance.Client, symbol string) {
	// Cancel order
	fmt.Println("")
	fmt.Println("Cancel All Open Order for:", symbol)
	cancelOrder, err := binanceClient.NewCancelOpenOrdersService().Symbol(symbol).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cancelOrder)
}

func GetExchangeInfo(binanceClient *binance.Client) {
	// Get Exchange information
	fmt.Println("Get Exchange Information")
	res, err := binanceClient.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func GetAccount(binanceClient *binance.Client) {
	fmt.Print("Get Account")
	account, err := binanceClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(account)
}

func ListAvgPriceBySymbol(binanceClient *binance.Client, symbol string) {
	// List Average Price by symbol
	fmt.Println("")
	fmt.Println("List Average Price for:", symbol)
	avgPrice, err := binanceClient.NewAveragePriceService().Symbol(symbol).Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(avgPrice)
}

func PingBinance(binanceClient *binance.Client) {
	if err := binanceClient.NewPingService().Do(context.Background()); err != nil {
		log.Warn(err)
		return
	}
	log.Info("Binance connection is Alive!")
}

func GetBinanceServerTime(binanceClient *binance.Client) {
	time, err := binanceClient.NewServerTimeService().Do(context.Background())
	if err != nil {
		log.Warn(err)
		return
	}
	log.Info("Binance Server Time:", time)
}

func SetBinanceServerTime(binanceClient *binance.Client) {
	binanceClient.NewSetServerTimeService().Do(context.Background())
	log.Info("Syncronzied Binance Server Time")
}

func GetSymbolLotSize(binanceClient *binance.Client, symbol string) float64 {
	res, err := binanceClient.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return 0
	}
	for i := range res.Symbols {
		if res.Symbols[i].Symbol == symbol {
			lotsize, err := strconv.ParseFloat(res.Symbols[i].LotSizeFilter().MinQuantity, 64)
			if err != nil {
				return 0
			}
			return lotsize
		}
	}
	return 0
}

func PrintAllAssetBalance(binanceClient *binance.Client) {
	account, err := binanceClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	for i := range account.Balances {
		p := account.Balances[i]
		a := p.Asset
		var f, l float64

		f, err1 := strconv.ParseFloat(p.Free, 64)
		if err1 != nil {
			fmt.Println(err1)
		}

		l, err2 := strconv.ParseFloat(p.Locked, 64)
		if err2 != nil {
			fmt.Println(err2)
		}

		if f > 0 || l > 0 {
			fmt.Printf("%v %v %v\n", a, f, l)
		}

	}
}

func GetAllAssetBalance(binanceClient *binance.Client) []binance.Balance {
	account, err := binanceClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return account.Balances
}

func GetAssetBalance(binanceClient *binance.Client, coin string) (free float64, locked float64) {
	account, err := binanceClient.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println("Account Balance:")
	//fmt.Printf("\nAccount Type: %v", account.AccountType)
	//fmt.Printf("\nBuyer Commission: %v", account.BuyerCommission)
	//fmt.Printf("\nSeller Commission: %v", account.SellerCommission)
	//fmt.Printf("\nMaker Commission: %v", account.MakerCommission)
	//fmt.Printf("\nTaker Commission: %v", account.TakerCommission)
	//fmt.Printf("\nUpdate Time: %v", account.UpdateTime)
	//fmt.Printf("\nCan Trade: %v", account.CanTrade)
	//fmt.Printf("\nCan Deposit: %v", account.CanDeposit)
	//fmt.Printf("\nCan Withdraw: %v", account.CanWithdraw)
	//fmt.Printf("\nPermission: %v", account.Permissions)
	//fmt.Printf("\nBalance: \n")

	for i := range account.Balances {
		p := account.Balances[i]
		var f, l float64
		a := p.Asset

		f, err1 := strconv.ParseFloat(p.Free, 64)
		if err1 != nil {
			fmt.Println(err1)
		}

		l, err2 := strconv.ParseFloat(p.Locked, 64)
		if err2 != nil {
			fmt.Println(err2)
		}

		if a == coin {
			//
			return f, l

			//json, err := json.Marshal(models.AccBal{Coin: a, Free: f, Locked: l})
			//if err != nil {
			//	fmt.Println(err)
			//}
			//return json
		}
	}
	return 0, 0
}
