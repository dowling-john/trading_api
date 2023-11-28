package trading_212

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	Trading212CandleStick struct {
		TimeStamp int64   `json:"time_stamp"`
		Opening   float64 `json:"opening"`
		Closing   float64 `json:"closing"`
		Low       float64 `json:"low"`
		High      float64 `json:"high"`
		Unknown   int     `json:"unknown"`
	}
)

func (c *Trading212CandleStick) UnmarshalJSON(p []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}

	//ToDO: This is so ugly it needs fixing
	err := json.Unmarshal(tmp[0], &c.TimeStamp)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = json.Unmarshal(tmp[1], &c.Opening)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	err = json.Unmarshal(tmp[2], &c.High)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = json.Unmarshal(tmp[3], &c.Low)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	err = json.Unmarshal(tmp[4], &c.Closing)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	err = json.Unmarshal(tmp[5], &c.Unknown)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

func (c *Trading212CandleStick) GetTimestamp() time.Time {
	return time.UnixMilli(c.TimeStamp)
}

func (c *Trading212CandleStick) GetOpening() float64 {
	return c.Opening
}

func (c *Trading212CandleStick) GetClosing() float64 {
	return c.Closing
}

func (c *Trading212CandleStick) GetHigh() float64 {
	return c.High
}

func (c *Trading212CandleStick) GetLow() float64 {
	return c.Low
}
