package main

import (
	"fmt"
	"trading_212_api/pricing"
	"trading_212_api/trading_212"
)

func main() {
	inst := trading_212.Trading212Instrument{
		Config: trading_212.Client{
			Env: trading_212.Demo,
		},
		Ticker: "MCD",
		Market: "",
	}

	candles, _ := inst.GetPricingHistory(pricing.OneMinute)
	for _, candle := range candles {
		fmt.Println(candle.TimeStamp, candle.GetTimestamp())
		fmt.Println(candle.GetOpening(), candle.GetClosing(), candle.GetHigh(), candle.GetLow())
	}
}
