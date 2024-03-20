package connector

import (
	"errors"

	"github.com/adshao/go-binance/v2"
	"github.com/shing1211/go-lib/config"
	log "github.com/sirupsen/logrus"
)

// Redis global variable for access
var BinanceClient *binance.Client

func InitBinanceConn(config config.BinanceConfig) error {
	apikey := config.APIKey
	secretkey := config.SecretKey

	log.Info("Connecting to Binance")
	binanceClient := binance.NewClient(apikey, secretkey)

	if binanceClient == nil {
		return errors.New("connect to Binance failed")
	}

	BinanceClient = binanceClient

	log.Info("Connected to Binance successfully!")
	return nil
}
