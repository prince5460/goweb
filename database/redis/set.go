package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()

	_, err = conn.Do("SET", "name", "cat")
	if err != nil {
		fmt.Println("redis set error:", err)
	}

	name, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		fmt.Printf("Got name: %s \n", name)
	}

	_, err = conn.Do("MSET", "name", "Tom", "age", 30)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Strings(conn.Do("MGET", "name", "age"))
	if err != nil {
		fmt.Println("redis get error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("MGET name: %s \n", res)
		fmt.Println(len(res))
	}

	_, err = conn.Do("SAdd", "zoo", "cat", "dog", "pig")
	if err != nil {
		fmt.Println("SAdd err=", err)
		return
	}

	_, err = conn.Do("SRem", "zoo", "pig")
	if err != nil {
		fmt.Println("SRem err=", err)
		return
	}

	r, err := redis.Strings(conn.Do("SMembers", "zoo"))
	if err != nil {
		fmt.Println("SMembers err=", err)
		return
	}

	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}

//设置有效时间
func testExpire(err error, conn redis.Conn) {
	_, err = conn.Do("expire", "name", 10)
	if err != nil {
		fmt.Println("expire err=", err)
		return
	}
}
