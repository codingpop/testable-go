package market_test

import (
	"errors"
	"testing"

	"github.com/codingpop/testable-go/cryptocurrency"
	"github.com/codingpop/testable-go/market"
	"github.com/google/go-cmp/cmp"
)

type Mock struct{}

func (Mock) Get(id int) (cryptocurrency.Result, error) {
	coins := map[int]cryptocurrency.Result{
		80: {
			ID:        80,
			Symbol:    "ETH",
			Name:      "Ethereum",
			Price:     1234.434,
			MarketCap: 328392.323,
		},
		90: {
			ID:        90,
			Symbol:    "BTC",
			Name:      "Bitcoin",
			Price:     40000.434,
			MarketCap: 2372397923749.323,
		},
	}

	coin, ok := coins[id]
	if !ok {
		return cryptocurrency.Result{}, errors.New("unsupported cryptocurrency")
	}

	return coin, nil
}

func TestGetCoinData(t *testing.T) {
	var tests = []struct {
		id     int
		symbol string
		want   market.Coin
	}{
		{80, "ETH", market.Coin{ID: 80, Symbol: "ETH", Name: "Ethereum", Price: 1234.434, MarketCap: 328392.323}},
		{90, "BTC", market.Coin{ID: 90, Symbol: "BTC", Name: "Bitcoin", Price: 40000.434, MarketCap: 2372397923749.323}},
	}

	var errTests = []struct {
		id     int
		symbol string
		want   market.Coin
	}{
		{100, "LRC", market.Coin{ID: 100, Symbol: "LRC", Name: "Loopring", Price: 0, MarketCap: 0}},
	}

	mock := Mock{}

	for _, tt := range tests {
		t.Run(tt.symbol, func(t *testing.T) {
			got, err := market.GetCoinData(tt.id, mock)
			if err != nil {
				t.Errorf("expecting nil err, got %v", err)
			}

			if !cmp.Equal(tt.want, got) {
				t.Errorf("expecting %#v, got %#v", tt.want, got)
			}
		})
	}

	for _, tt := range errTests {
		t.Run(tt.symbol, func(t *testing.T) {
			_, err := market.GetCoinData(tt.id, mock)
			if err == nil {
				t.FailNow()
			}
		})
	}
}
