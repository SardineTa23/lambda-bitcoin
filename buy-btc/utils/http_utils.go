package utils

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func DoHttpRequest(method, url string, header, query map[string]string, data []byte) ([]byte, error) {
	if method != "GET" && method != "POST" {
		return nil, errors.New("method's neither GET nor POST")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	for k, v := range header {
		req.Header.Add(k, v)
	}

	httpClient := &http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	// 関数終了時の処理
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
