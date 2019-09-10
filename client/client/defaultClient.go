package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

//使用默认的DefaultClient对象
func main() {
	url := "https://www.baidu.com"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", dumpResponse)
}
