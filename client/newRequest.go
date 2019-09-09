package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/**
  客户端实现方式一：
  1.先生成http.client ->
  2.再生成 http.request ->
  3.之后提交请求：client.Do(request) ->
  4.处理返回结果，每一步的过程都可以设置一些具体的参数。
*/
func main() {
	url := "https://www.baidu.com"

	client := http.Client{} //创建client对象

	request, err := http.NewRequest("GET", url, nil) //创建request

	response, err := client.Do(request) //发送请求
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

}
