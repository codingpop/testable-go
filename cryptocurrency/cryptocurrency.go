package cryptocurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Cryptocurrency holds the cryptocurrency data
type Cryptocurrency struct {
	ID        int
	Symbol    string
	Name      string
	Price     float64
	MarketCap float64
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (c *Cryptocurrency) UnmarshalJSON(data []byte) error {
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

	c.ID = id
	c.Symbol = symbol
	c.Name = name
	c.Price = price
	c.MarketCap = marketCap

	return nil
}

// Get gets cryptocurrency info
func Get(id int) (_ Cryptocurrency, retErr error) {
	resp, err := http.Get(fmt.Sprintf("https://api.coinlore.net/api/ticker/?id=%d", id))
	if err != nil {
		return Cryptocurrency{}, err
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			retErr = err
		}
	}()

	var c Cryptocurrency

	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return Cryptocurrency{}, err
	}

	return c, nil
}
