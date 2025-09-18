package config

import (
	"errors"
	"os"

	"github.com/ValeryCherneykin/binance-dca-bot/internal/logger"
	"go.uber.org/zap"
)

const (
	binanceAPIKey = "BINANCE_API_KEY"
	binanceSecret = "BINANCE_SECRET"
)

type BinanceConfig interface {
	APIKey() string
	Secret() string
}

type binanceConfig struct {
	apiKey string
	secret string
}

func NewBinanceConfig() (BinanceConfig, error) {
	apiKey := os.Getenv(binanceAPIKey)
	if len(apiKey) == 0 {
		logger.Error("binance api key not found", zap.String("env_var", binanceAPIKey))
		return nil, errors.New("binance api key not found")
	}

	secret := os.Getenv(binanceSecret)
	if len(secret) == 0 {
		logger.Error("binance secret not found", zap.String("env_var", binanceSecret))
		return nil, errors.New("binance secret not found")
	}

	cfg := &binanceConfig{
		apiKey: apiKey,
		secret: secret,
	}

	logger.Info("binance config loaded")

	return cfg, nil
}

func (cfg *binanceConfig) APIKey() string {
	return cfg.apiKey
}

func (cfg *binanceConfig) Secret() string {
	return cfg.secret
}
