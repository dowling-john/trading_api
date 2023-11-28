package pricing

import (
	"time"
)

const (
	OneMinute     Interval = "ONE_MINUTE"
	FiveMinute    Interval = "FIVE_MINUTE"
	TenMinute     Interval = "TEN_MINUTE"
	FifteenMinute Interval = "FIFTEEN_MINUTE "
	ThirtyMinute  Interval = "THIRTY_MINUTE"
	OneHour       Interval = "THIRTY_MINUTE"
	FourHour      Interval = "THIRTY_MINUTE"
	OneDay        Interval = "THIRTY_MINUTE"
	OneWeek       Interval = "THIRTY_MINUTE"
	OneMonth      Interval = "THIRTY_MINUTE"
)

type (
	Interval string

	CandleStick interface {
		GetTimestamp() time.Time
		GetOpening() float64
		GetClosing() float64
		GetHigh() float64
		GetLow() float64
	}
)
