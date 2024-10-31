package test

import (
	"fmt"
	"testing"

	"github.com/sunrisedotdev/goingecko"
)

func TestCoins(t *testing.T) {
	coin := "bitcoin"

	cgClient := goingecko.NewClient(nil, "")

	coinData, _ := cgClient.CoinsId(coin, true, true, true, true, true, true)
	fmt.Println(coinData.MarketData.CurrentPrice.Usd)

}
