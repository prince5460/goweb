package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()

	//向redis写入数据string[k-v]
	_, err = conn.Do("Set", "name", "tom&jerry")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	//从redis中读取数据string[k-v]
	reply, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("success:", reply)

	//批量操作string
	_, err = conn.Do("MSet", "name", "James", "age", 20, "address", "beijing")
	if err != nil {
		fmt.Println("MSet err=", err)
		return
	}

	strings, err := redis.Strings(conn.Do("MGet", "name", "age", "address"))
	if err != nil {
		fmt.Println("MGet err=", err)
		return
	}
	for _, v := range strings {
		fmt.Println(v)
	}
}
