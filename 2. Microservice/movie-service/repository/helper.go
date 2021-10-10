package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GenerateApiRequest(method, url, endpoint string, body interface{}, queryParam map[string]string) (*http.Request, error) {
	marshaled, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url+endpoint, bytes.NewReader(marshaled))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Host", url)
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	for key, paramValue := range queryParam {
		q.Add(key, paramValue)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}
