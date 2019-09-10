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
	urlStr := "http://127.0.0.1:2001/login"
	client := http.Client{}
	param := url.Values{
		"username": {"zhou123"},
		"password": {"123456"},
	}

	//jsonStr, _ := json.Marshal(param)
	//requestBody := bytes.NewBufferString(string(jsonStr))

	requestBody := bytes.NewBufferString(param.Encode())
	request, err := http.NewRequest("POST", urlStr, requestBody)
	//response, err := client.Post(urlStr, "application/x-www-form-urlencoded", requestBody)
	//response, err := http.Post(urlStr, "application/x-www-form-urlencoded", requestBody)
	//response, err := http.PostForm(urlStr, param)

	if err != nil {
		log.Fatal(err)
	}

	/*
	   cookie的添加
	   方式一：
	       request.Header.Set("Cookie", "name=hanru")
	   方式二：
	       request.AddCookie(Cookie)
	*/
	cookieId := &http.Cookie{Name: "userId", Value: "1234"}
	cookieName := &http.Cookie{Name: "name", Value: "golang"}
	request.AddCookie(cookieId)
	request.AddCookie(cookieName)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("response:%v\n", response)
	fmt.Printf("response.Header:%v\n", response.Header)
	fmt.Printf("response.Cookies:%v\n", response.Cookies())
	fmt.Printf("request.Header:%v\n", request.Header)
	fmt.Printf("request.Cookies:%v\n", request.Cookies())

}
