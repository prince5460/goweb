package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	file, err := os.Open(fileName)

	defer file.Close()

	if err != nil {
		fmt.Println(fileName, err)
		return
	}

	bytes := make([]byte, 1024)
	for {
		n, _ := file.Read(bytes)
		if 0 == n {
			break
		}
		os.Stdout.Write(bytes[:n])
	}
}
