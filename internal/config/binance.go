package config

import (
	"errors"
	"os"

	"github.com/ValeryCherneykin/binance-dca-bot/internal/logger"
	"go.uber.org/zap"
)

const (
	binanceAPIKey  = "BINANCE_API_KEY"
	binanceSecret  = "BINANCE_SECRET"
	binanceTestnet = "BINANCE_TESTNET"
	dryRun         = "DRY_RUN"
)

type BinanceConfig interface {
	APIKey() string
	Secret() string
	IsTestnet() bool
	IsDryRun() bool
}

type binanceConfig struct {
	apiKey    string
	secret    string
	isTestnet bool
	isDryRun  bool
}

func NewBinanceConfig() (BinanceConfig, error) {
	apiKey := os.Getenv(binanceAPIKey)
	if apiKey == "" {
		logger.Error("Binance API key not found", zap.String("env_var", binanceAPIKey))
		return nil, errors.New("binance api key not found")
	}

	secret := os.Getenv(binanceSecret)
	if secret == "" {
		logger.Error("Binance secret not found", zap.String("env_var", binanceSecret))
		return nil, errors.New("binance secret not found")
	}

	cfg := &binanceConfig{
		apiKey:    apiKey,
		secret:    secret,
		isTestnet: os.Getenv(binanceTestnet) == "1" || os.Getenv(binanceTestnet) == "true",
		isDryRun:  os.Getenv(dryRun) == "1" || os.Getenv(dryRun) == "true",
	}

	logger.Info("Binance config loaded",
		zap.Bool("testnet", cfg.isTestnet),
		zap.Bool("dry_run", cfg.isDryRun))

	return cfg, nil
}

func (cfg *binanceConfig) APIKey() string {
	return cfg.apiKey
}

func (cfg *binanceConfig) Secret() string {
	return cfg.secret
}

func (cfg *binanceConfig) IsTestnet() bool {
	return cfg.isTestnet
}

func (cfg *binanceConfig) IsDryRun() bool {
	return cfg.isDryRun
}
