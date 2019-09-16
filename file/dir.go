package main

import (
	"fmt"
	"os"
)

//操作目录

func main() {
	os.Mkdir("file/tom", 0777)
	os.MkdirAll("file/tom/test1/test2", 0777)

	//删除单个目录a
	err := os.Remove("file/tom")
	if err != nil {
		fmt.Println("err:", err)
	}

	//删除多级子目录
	os.RemoveAll("file/tom")
}
