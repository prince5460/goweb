package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//使用postform
func main() {
	urlStr := "http://127.0.0.1:8080/login"
	param := url.Values{
		"username": {"zhou"},
		"password": {"123456"},
	}
	resp, err := http.PostForm(urlStr, param)
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
	fmt.Printf("response:%+v\n", resp)
	fmt.Printf("response.Header:%v\n", resp.Header)
	fmt.Printf("response.Cookies:%v\n", resp.Cookies())
}
