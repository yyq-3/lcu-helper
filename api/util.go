package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

func get(url string) (data []byte, err error) {
	return request("GET", url, nil)
}

func post(url string, body any) (data []byte, err error) {
	return request("POST", url, body)
}

func request(method, url string, reqBody any) (data []byte, err error) {
	var body io.Reader
	if reqBody != nil {
		bts, err := json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(bts)
	}
	req, _ := http.NewRequest(method, apiAddr+url, body)
	if req.Body != nil {
		req.Header.Add("ContentType", "application/json")
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
