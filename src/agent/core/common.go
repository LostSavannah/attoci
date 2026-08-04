package core

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type HttpClient struct {
	BaseUrl        string
	ApiKey         string
	TimeoutSeconds int
}

func (client HttpClient) Send(method string, url string, data interface{}) (response []byte, err error) {
	err = nil
	var body []byte = nil
	if data != nil {
		body, err = json.Marshal(data)
		if err != nil {
			return
		}
	}
	req, err := http.NewRequest(
		method,
		client.BaseUrl+url,
		bytes.NewReader(body),
	)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", client.ApiKey)
	innerClient := &http.Client{Timeout: time.Second * time.Duration(client.TimeoutSeconds)}
	res, err := innerClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	response, err = io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
