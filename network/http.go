package network

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpDoRequest(client *http.Client, method, url string, header map[string]string, data []byte) ([]byte, error) {
	newRequest, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		newRequest.Header.Set(k, v)
	}

	resp, err := client.Do(newRequest)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
