package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
	   遍历文件夹
	*/
	dirName := "image"
	listDir(dirName)
}

func listDir(dirName string) {
	fileInfos, _ := ioutil.ReadDir(dirName)
	for _, fi := range fileInfos {
		fileName := dirName + "/" + fi.Name()
		fmt.Println(fileName)
		if fi.IsDir() {
			listDir(fileName)
		}
	}
}
