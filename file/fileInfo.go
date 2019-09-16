package main

import (
	"fmt"
	"os"
)

func main() {
	/*
	   fileinfo：文件的信息
	       interface：
	           Name()
	           IsDir()
	           Mode()
	           Size()
	           ModTime()

	*/
	fileInfo, err := os.Stat("hello.txt")
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	fmt.Printf("%T\n", fileInfo)
	//文件名
	fmt.Println(fileInfo.Name()) //hello.txt
	//是否是目录
	fmt.Println(fileInfo.IsDir()) //false
	//大小
	fmt.Println(fileInfo.Size()) //1
	//权限
	fmt.Println(fileInfo.Mode()) //-rw-rw-rw-
	//修改时间
	fmt.Println(fileInfo.ModTime())

}
