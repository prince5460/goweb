package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	for i := 0; i < 10; i++ {
		file.WriteString("just a test!\r\n")
		file.Write([]byte("some words!\r\n"))
	}

}
