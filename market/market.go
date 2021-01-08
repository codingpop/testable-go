package market

import "github.com/codingpop/testable-go/cryptocurrency"

const (
	eth = 80
	btc = 90
)

// Coin represents the cryptocurrency data
type Coin struct {
	ID        int
	Symbol    string
	Name      string
	Price     float64
	MarketCap float64
}

// GetCoinData fetches a cryptocurrency info
func GetCoinData(symbol int) (Coin, error) {
	crypto, err := cryptocurrency.Get(eth)
	if err != nil {
		return Coin{}, err
	}

	return Coin(crypto), nil
}
