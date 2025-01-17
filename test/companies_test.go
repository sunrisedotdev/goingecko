package test

import (
	"testing"

	"github.com/sunrisedotdev/goingecko"
)

func TestPublicTreasuryCoinId(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.PublicTreasuryCoinId("bitcoin")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}
