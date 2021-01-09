package market

import (
	"github.com/codingpop/testable-go/cryptocurrency"
)

// CoinGetter defines an interface for the cryptocurrency dependency
type CoinGetter interface {
	Get(int) (cryptocurrency.Result, error)
}

// Coin represents the cryptocurrency data
type Coin struct {
	ID        int
	Symbol    string
	Name      string
	Price     float64
	MarketCap float64
}

// GetCoinData fetches a cryptocurrency info
func GetCoinData(symbol int, cg CoinGetter) (Coin, error) {
	crypto, err := cg.Get(symbol)
	if err != nil {
		return Coin{}, err
	}

	return Coin(crypto), nil
}
