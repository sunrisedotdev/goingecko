package coins

import "github.com/sunrisedotdev/goingecko/types"

type Tickers struct {
	Name    string         `json:"name"`
	Tickers []types.Ticker `json:"tickers"`
}
