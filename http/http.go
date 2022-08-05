package http

import (
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGetRequest(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 5, // Timeout after 5 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, err
	}
	defer res.Body.Close()
	retbody, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}
	return retbody, nil
}
