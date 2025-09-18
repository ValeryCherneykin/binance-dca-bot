package binance

import (
	"github.com/ValeryCherneykin/binance-dca-bot/internal/config"
	"go.uber.org/zap"
)

const (
	defaultAPIBase = "https://api.binance.com"
	testnetAPIBase = "https://testnet.binance.vision"
	defaultWSHost  = "wss://stream.binance.com:9443"
	testnetWSHost  = "wss://testnet.binance.vision"
)

type binanceClient struct {
	config  config.BinanceConfig
	logger  *zap.Logger
	apiBase string
	wsHost  string
}
