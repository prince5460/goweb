package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	httpPost()
}

func httpPost() {
	urlStr := "http://127.0.0.1:2001/login"
	contentType := "application/x-www-form-urlencoded"
	param := &url.Values{
		"username": {"zhou123"},
		"password": {"123456"},
	}
	requestBody := bytes.NewBufferString(param.Encode())
	resp, err := http.Post(urlStr, contentType, requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("response:%v\n", resp)
	fmt.Printf("response.Header:%v\n", resp.Header)
	fmt.Printf("response.Cookies:%v\n", resp.Cookies())
}
