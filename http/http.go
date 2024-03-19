package http

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	myjsonlib "github.com/shing1211/go-lib/json"
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
	retbody, err := io.ReadAll(res.Body)
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
	bodyBytes, err := io.ReadAll(resp.Body)
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
	bodyBytes, err := io.ReadAll(resp.Body)
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
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 500, nil, err
	}
	return resp.StatusCode, bodyBytes, err
}

var (
	username = "elastic"
	password = "TMgmilyMmFeeHWEHiLRP"
	//url      = "http://192.168.1.86:9200/_license"
	//geturl = "http://192.168.1.86:9200/logstash-2021-12-09/_search"
	geturl = "http://192.168.1.86:9200/golanglog-2021-12-14/_doc/HdNcuH0BqZrXMDWgtdle"
	//url = "https://pokeapi.co/api/v2/pokedex/kantso/"
	//url = "http://192.168.1.86:9200/logstash-*/_search"
	//posturl = "http://192.168.1.86:9200/logstash-2021-12-14/_doc/"
	posturl = "http://192.168.1.86:9200/golanglog-2021-12-14/_doc/"
)

var exit = make(chan bool)

func httpBasicAuthRequest(method string, url string, username string, password string, jsonData []byte) (responsebody []byte, err error) {
	// Concatenate customer key and customer secret and use base64 to encode the concatenated string
	plainCredentials := username + ":" + password
	base64Credentials := base64.StdEncoding.EncodeToString([]byte(plainCredentials))

	payload := bytes.NewBuffer(jsonData)
	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Add Authorization header
	req.Header.Add("Authorization", "Basic "+base64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	//fmt.Println("response Status:", res.Status)
	//fmt.Println("response Headers:", res.Header)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}

func loopHttpPostDocReq(routine string) {
	sum := 0
	for {
		sum++
		count := strconv.FormatInt(int64(sum), 10)

		// Http POST request with basic auth and request body
		postBody, _ := json.Marshal(map[string]string{
			"@timestamp": time.Now().UTC().Format("2006-01-02T15:04:05.000"),
			"message":    "tchan golang test post request " + routine + " " + count,
			"tag":        "golang-elk-test",
			"user":       "tchan",
		})

		newPostRes, err := httpBasicAuthRequest("POST", posturl, username, password, postBody)

		if err != nil {
			log.Fatal(err)
		}

		//newPostRes, err = formatJSON(newPostRes)

		//if err != nil {
		//	log.Fatal(err)
		//}

		fmt.Println(string(newPostRes))
	}
}

func testMain() {
	// Http GET request with basic auth
	newRes, err := httpBasicAuthRequest("GET", geturl, username, password, nil)

	if err != nil {
		log.Fatal(err)
	}

	newRes, err = myjsonlib.FormatJSON(newRes)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(newRes))

	go loopHttpPostDocReq("routine 1")
	go loopHttpPostDocReq("routine 2")
	go loopHttpPostDocReq("routine 3")
	<-exit
	fmt.Println("Execution completed! Program exit.")
}
