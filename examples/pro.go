package main

import (
	"fmt"

	"github.com/sunrisedotdev/goingecko"
)

func main() {
	cgClient := goingecko.NewClient(nil, "pro api key", true)
	defer cgClient.Close()

	data, err := cgClient.CoinsId("bitcoin", true, true, true, false, false, false)
	if err != nil {
		fmt.Print("Somethig went wrong...")
		return
	}
	fmt.Printf("Bitcoin price is: %f$", data.MarketData.CurrentPrice.Usd)
}
