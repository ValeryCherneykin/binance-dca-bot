package config

import (
	"errors"
	"os"
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
		return nil, errors.New("binance api key not found")
	}

	secret := os.Getenv(binanceSecret)
	if len(secret) == 0 {
		return nil, errors.New("binance secret key not found")
	}

	cfg := &binanceConfig{
		apiKey: apiKey,
		secret: secret,
	}

	return cfg, nil
}

func (cfg *binanceConfig) APIKey() string {
	return cfg.apiKey
}

func (cfg *binanceConfig) Secret() string {
	return cfg.secret
}
