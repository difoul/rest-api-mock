package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, auth string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if len(auth) > 0 {
		req.Header.Set("Authorization", auth)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("GET %s returned %d status", url, resp.StatusCode))
	}
	return body, nil
}

func HttpPost(url string, bodyJson string, auth string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyJson)))
	req.Header.Set("Authorization", auth)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("POST %s returned %d status", url, resp.StatusCode))
	}
	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}
