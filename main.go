package main

import (
	"fmt"

	"github.com/codingpop/testable-go/cryptocurrency"
	"github.com/codingpop/testable-go/market"
)

const (
	eth = 80
	btc = 90
)

func main() {
	c := cryptocurrency.New()
	coin, err := market.GetCoinData(eth, c)
	if err != nil {
		return
	}

	fmt.Println(coin)
}
