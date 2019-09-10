package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//通过http包的Post()实现
/**
注意：使用post方法，要设置contentType，
设置成”application/x-www-form-urlencoded”，
否则post参数无法传递。
*/
func main() {
	urlStr := "http://www.webxml.com.cn/WebServices/WeatherWebService.asmx/getWeatherbyCityName"
	contentType := "application/x-www-form-urlencoded"
	requestBody := strings.NewReader("theCityName=北京")

	resp, err := http.Post(urlStr, contentType, requestBody)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}
