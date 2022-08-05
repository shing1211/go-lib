package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGetRequest(url string, timeout time.Duration) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 500, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{Timeout: time.Second * timeout}
	res, err := client.Do(req)
	if err != nil {
		return 500, nil, err
	}
	defer res.Body.Close()
	retbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 500, nil, err
	}
	return res.StatusCode, retbody, nil
}

func HttpPostRequest(url string, timeout time.Duration, jsonReq []byte) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return 500, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{Timeout: time.Second * timeout}
	resp, err := client.Do(req)
	if err != nil {
		return 500, nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}
	return resp.StatusCode, bodyBytes, nil
}

func HttpPutRequest(url string, timeout time.Duration, jsonReq []byte) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return 500, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{Timeout: time.Second * timeout}
	resp, err := client.Do(req)
	if err != nil {
		return 500, nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}
	return resp.StatusCode, bodyBytes, err
}

func HttpDeleteRequest(url string, timeout time.Duration, jsonReq []byte) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return 500, nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{Timeout: time.Second * timeout}
	resp, err := client.Do(req)
	if err != nil {
		return 500, nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}
	return resp.StatusCode, bodyBytes, err
}
