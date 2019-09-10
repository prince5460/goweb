package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targetUrl string) error {
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	file, err := writer.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer:", err)
		return err
	}
	//打开文件句柄
	openfile, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}

	//iocopy
	_, err = io.Copy(file, openfile)
	if err != nil {
		fmt.Println("error copy file:", err)
		return err
	}
	contentType := writer.FormDataContentType()
	writer.Close()

	resp, err := http.Post(targetUrl, contentType, buffer)
	if err != nil {
		fmt.Println("post err:", err)
		return err
	}
	defer resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err:", err)
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main() {
	target_url := "http://localhost:8080/upload"
	filename := "./upload.html"
	postFile(filename, target_url)
}
