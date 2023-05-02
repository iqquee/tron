package tron

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Http   *http.Client
	ApiKey string
}


const (
	MainNet       = "https://api.trongrid.io"
	ShastaTestNet = "https://api.shasta.trongrid.io"
	NileTestNet   = "https://nile.trongrid.io"
)

func New(apiKey string) *Client {
	return &Client{
		ApiKey: apiKey,
	}
}


func (c *Client) NewRequest(method, url string, isPayload bool, payload interface{}) ([]byte, int, error) {
	var newPayload []byte
	if isPayload {
		jsonReq, jsonReqErr := json.Marshal(&payload)
		if jsonReqErr != nil {
			return nil, 0, jsonReqErr
		}
		newPayload = jsonReq
	}

	req, reqErr := http.NewRequest(method, url, bytes.NewBuffer(newPayload))
	if reqErr != nil {
		return nil, 0, reqErr
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TRON-PRO-API-KEY", c.ApiKey)

	resp, respErr := c.Http.Do(req)
	if respErr != nil {
		return nil, 0, respErr
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp_body, resp.StatusCode, nil
}
