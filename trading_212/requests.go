package trading_212

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("User-Agent", "PostmanRuntime/7.35.0")
}

func apiPut(url string, payload []byte) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}

	setHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	return body
}
