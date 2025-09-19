package cryptorg

type CryptorgClient interface {
	PlaceMarketBuy(symbol string, quantity float64) (string, error)
	PlaceLimitSell(symbol string, quantity, price float64) (string, error)
	CancelOrder(orderID string) error
}
