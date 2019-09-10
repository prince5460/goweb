package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//通过client对象调用Post()方法
func main() {
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName"
	client := http.Client{}
	param := &url.Values{
		"theCityName": {"北京"},
	}
	requestBody := bytes.NewBufferString(param.Encode())

	response, err := client.Post(urlStr, "application/x-www-form-urlencoded", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}

	fmt.Printf("%v", response)

}
