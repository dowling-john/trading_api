package instrument

import (
	"trading_212_api/pricing"
)

type (
	Instrument interface {
		CreateBuyOrder() (bool, error)
		CreateSellOrder() (bool, error)

		GetPricingHistory(interval pricing.Interval) ([]*pricing.CandleStick, error)
	}

	Position interface {
	}
)
