package trading_212

import (
	"encoding/json"
	"fmt"
	"trading_212_api/pricing"
)

type (
	Trading212Instrument struct {
		Config Client
		Ticker string
		Name   string
		Market string
	}

	PricingRequest struct {
		Ticker      string `json:"ticker"`
		UseAskPrice bool   `json:"useAskPrice"`
		Period      string `json:"period"`
		Size        int    `json:"size"`
	}

	PricingRequests struct {
		Candles []*PricingRequest `json:"candles"`
	}

	PricingResponse struct {
		Candles json.RawMessage `json:"candles"`
	}
)

func (i *Trading212Instrument) getPricingInformation(url string, data json.RawMessage) []byte {
	return i.Config.RemoveOuterResponse(apiPut(url, data))
}

func (i *Trading212Instrument) handlePricingResponse(data json.RawMessage) PricingResponse {
	var c = PricingResponse{}
	_ = json.Unmarshal(data, &c)
	return c
}

func (i *Trading212Instrument) buildCandles(p PricingResponse) []*Trading212CandleStick {
	var candles []*Trading212CandleStick
	_ = json.Unmarshal(p.Candles, &candles)
	return candles
}

func (i *Trading212Instrument) GetPricingHistory(interval pricing.Interval) ([]*Trading212CandleStick, error) {
	data, _ := json.Marshal(PricingRequests{
		Candles: []*PricingRequest{
			{Ticker: i.Ticker, UseAskPrice: true, Period: string(interval), Size: 500},
		},
	})

	return i.buildCandles(
		i.handlePricingResponse(
			i.getPricingInformation(
				fmt.Sprintf("%s%s%s/%s?ticker=%s&languageCode=en", Protocol, i.Config.Env, BaseUrl, Charting+"/"+Version3+"/"+Candles, i.Ticker),
				data,
			),
		),
	), nil
}
