package trading_212

import "encoding/json"

type (
	Client struct {
		Env      Environment
		userName string
		password string
	}

	Environment string

	Response struct {
		Req json.RawMessage `json:"request"`
		Res json.RawMessage `json:"response"`
	}
)

const (
	Live Environment = "live"
	Demo Environment = "demo"

	Protocol string = "https://"
	BaseUrl  string = ".trading212.com"
	Version2 string = "v2"
	Version3 string = "v3"
	Charting string = "charting"
	Candles  string = "candles"
)

func (c *Client) RemoveOuterResponse(response []byte) (j json.RawMessage) {
	var r []*Response
	_ = json.Unmarshal(response, &r)
	return r[0].Res
}

func (c *Client) Login() error {
	return nil
}
