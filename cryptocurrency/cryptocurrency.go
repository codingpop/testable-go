package cryptocurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	ID        int
	Symbol    string
	Name      string
	Price     float64
	MarketCap float64
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (r *Result) UnmarshalJSON(data []byte) error {
	var slice []struct {
		ID        string `json:"id"`
		Symbol    string `json:"symbol"`
		Name      string `json:"name"`
		Price     string `json:"price_usd"`
		MarketCap string `json:"market_cap_usd"`
	}

	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}

	id, err := strconv.Atoi(slice[0].ID)
	if err != nil {
		return err
	}
	price, err := strconv.ParseFloat(slice[0].Price, 64)
	if err != nil {
		return err
	}
	marketCap, err := strconv.ParseFloat(slice[0].MarketCap, 64)
	if err != nil {
		return err
	}
	symbol := slice[0].Symbol
	name := slice[0].Name

	r.ID = id
	r.Symbol = symbol
	r.Name = name
	r.Price = price
	r.MarketCap = marketCap

	return nil
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

	var r Result

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return Result{}, err
	}

	return r, nil
}
