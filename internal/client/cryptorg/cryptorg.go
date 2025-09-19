package cryptorg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ValeryCherneykin/cryptorg-dca-bot/internal/config"
	"github.com/ValeryCherneykin/cryptorg-dca-bot/internal/logger"
	"go.uber.org/zap"
)

type cryptorgClient struct {
	cfg config.CryptorgConfig
}

func NewCryptorgClient(cfg config.CryptorgConfig) CryptorgClient {
	return &cryptorgClient{cfg: cfg}
}

func (c *cryptorgClient) doRequest(method, endpoint string, body any) ([]byte, error) {
	if c.cfg.IsDryRun() {
		logger.Info("Dry-run: Simulating request", zap.String("method", method), zap.String("endpoint", endpoint))
		return []byte(`{"id":"dry-order-id"}`), nil
	}

	url := "https://api.cryptorg.net/v1" + endpoint
	var reqBody []byte
	if body != nil {
		var err error
		reqBody, err = json.Marshal(body)
		if err != nil {
			logger.Error("Failed to marshal request body", zap.Error(err))
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		logger.Error("Failed to create HTTP request", zap.Error(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Access-ID", c.cfg.AccessID())
	req.Header.Set("X-API-Key", c.cfg.APIKey())

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("HTTP request failed", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response body", zap.Error(err))
		return nil, err
	}
	if resp.StatusCode >= 400 {
		logger.Error("HTTP request error", zap.Int("status_code", resp.StatusCode), zap.String("response", string(b)))
		return nil, fmt.Errorf("http %d: %s", resp.StatusCode, string(b))
	}
	return b, nil
}

func (c *cryptorgClient) PlaceMarketBuy(symbol string, quantity float64) (string, error) {
	payload := map[string]any{
		"symbol":   symbol,
		"side":     "BUY",
		"type":     "MARKET",
		"quantity": quantity,
	}
	b, err := c.doRequest("POST", "/orders", payload)
	if err != nil {
		return "", err
	}
	var resp map[string]any
	if err := json.Unmarshal(b, &resp); err != nil {
		logger.Error("Failed to unmarshal response", zap.Error(err))
		return "", err
	}
	if id, ok := resp["id"].(string); ok {
		return id, nil
	}
	return "", fmt.Errorf("no order id in response")
}

func (c *cryptorgClient) PlaceLimitSell(symbol string, quantity, price float64) (string, error) {
	payload := map[string]any{
		"symbol":   symbol,
		"side":     "SELL",
		"type":     "LIMIT",
		"quantity": quantity,
		"price":    price,
	}
	b, err := c.doRequest("POST", "/orders", payload)
	if err != nil {
		return "", err
	}
	var resp map[string]any
	if err := json.Unmarshal(b, &resp); err != nil {
		logger.Error("Failed to unmarshal response", zap.Error(err))
		return "", err
	}
	if id, ok := resp["id"].(string); ok {
		return id, nil
	}
	return "", fmt.Errorf("no order id in response")
}

func (c *cryptorgClient) CancelOrder(orderID string) error {
	if c.cfg.IsDryRun() {
		logger.Info("Dry-run: Simulating order cancellation", zap.String("order_id", orderID))
		return nil
	}
	_, err := c.doRequest("DELETE", "/orders/"+orderID, nil)
	if err != nil {
		logger.Error("Failed to cancel order", zap.String("order_id", orderID), zap.Error(err))
	}
	return err
}
