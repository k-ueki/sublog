package util

import (
	"bytes"
	"log"
	"net/http"
)

func HttpGet(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	return res, nil
}

func HttpPost(url, jsonContent string) error {
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(jsonContent)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
