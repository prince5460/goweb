package main

import (
	"fmt"
	"net/http"
)

func main() {
	requestUrl := "https://www.baidu.com"
	response, err := http.Get(requestUrl)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer response.Body.Close()
	fmt.Println(response.Body)

}
