package binance

import "encoding/json"

type OrderResponse struct {
	Symbol        string          `json:"symbol"`
	OrderId       int64           `json:"orderId"`
	ClientOrderId string          `json:"clientOrderId"`
	TransactTime  int64           `json:"transactTime"`
	Status        string          `json:"status"`
	Fills         json.RawMessage `json:"fills,omitempty"`
}
