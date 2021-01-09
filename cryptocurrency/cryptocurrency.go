package cryptocurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Cryptocurrency defines some methods that can be use to
// fetch cryptocurrency data
type Cryptocurrency struct {
	api string
}

// New creates a new instance of Cryptocurrency
func New() Cryptocurrency {
	return Cryptocurrency{
		api: "https://api.coinlore.net/api/ticker",
	}
}

// Result holds the cryptocurrency data
type Result struct {
	ID        int     `json:"id,string"`
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Price     float64 `json:"price_usd,string"`
	MarketCap float64 `json:"market_cap_usd,string"`
}

// Get gets cryptocurrency info
func (c Cryptocurrency) Get(id int) (_ Result, retErr error) {
	resp, err := http.Get(fmt.Sprintf("%s?id=%d", c.api, id))
	if err != nil {
		return Result{}, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			retErr = err
		}
	}()

	var r []Result

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return Result{}, err
	}

	return r[0], nil
}
