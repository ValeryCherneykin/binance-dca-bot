package config

import (
	"errors"
	"os"

	"github.com/ValeryCherneykin/cryptorg-dca-bot/internal/logger"
	"go.uber.org/zap"
)

const (
	cryptorgAccessID = "CRYPTORG_ACCESS_ID"
	cryptorgAPIKey   = "CRYPTORG_API_KEY"
	cryptorgSecret   = "CRYPTORG_SECRET"
	binanceTestnet   = "BINANCE_TESTNET"
	dryRun           = "DRY_RUN"
)

type CryptorgConfig interface {
	AccessID() string
	APIKey() string
	Secret() string
	IsTestnet() bool
	IsDryRun() bool
}

type cryptorgConfig struct {
	accessID  string
	apiKey    string
	secret    string
	isTestnet bool
	isDryRun  bool
}

func NewCryptorgConfig() (CryptorgConfig, error) {
	accessID := os.Getenv(cryptorgAccessID)
	if accessID == "" {
		logger.Error("Cryptorg Access ID not found", zap.String("env_var", cryptorgAccessID))
		return nil, errors.New("cryptorg access id not found")
	}

	apiKey := os.Getenv(cryptorgAPIKey)
	if apiKey == "" {
		logger.Error("Cryptorg API key not found", zap.String("env_var", cryptorgAPIKey))
		return nil, errors.New("cryptorg api key not found")
	}

	secret := os.Getenv(cryptorgSecret)
	if secret == "" {
		logger.Error("Cryptorg secret not found", zap.String("env_var", cryptorgSecret))
		return nil, errors.New("cryptorg secret not found")
	}

	cfg := &cryptorgConfig{
		accessID:  accessID,
		apiKey:    apiKey,
		secret:    secret,
		isTestnet: os.Getenv(binanceTestnet) == "1" || os.Getenv(binanceTestnet) == "true",
		isDryRun:  os.Getenv(dryRun) == "1" || os.Getenv(dryRun) == "true",
	}

	logger.Info("Cryptorg config loaded",
		zap.String("access_id", cfg.accessID),
		zap.Bool("testnet", cfg.isTestnet),
		zap.Bool("dry_run", cfg.isDryRun))

	return cfg, nil
}

func (c *cryptorgConfig) AccessID() string {
	return c.accessID
}

func (c *cryptorgConfig) APIKey() string {
	return c.apiKey
}

func (c *cryptorgConfig) Secret() string {
	return c.secret
}

func (c *cryptorgConfig) IsTestnet() bool {
	return c.isTestnet
}

func (c *cryptorgConfig) IsDryRun() bool {
	return c.isDryRun
}
