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

	//操作hash
	_, err = conn.Do("HSET", "user", "name", "john", "age", 30)
	if err != nil {
		fmt.Println("redis mset error:", err)
	}
	res, err := redis.Int64(conn.Do("HGET", "user", "age"))
	if err != nil {
		fmt.Println("redis HGET error:", err)
	} else {
		res_type := reflect.TypeOf(res)
		fmt.Printf("res type : %s \n", res_type)
		fmt.Printf("res  : %d \n", res)
	}

	//批量操作hash
	_, err = conn.Do("HMSet", "user2", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("HMSet err=", err)
		return
	}

	r, err := redis.Strings(conn.Do("HMGet", "user2", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err=", err)
		return
	}
	for i, v := range r {
		fmt.Printf("r[%d]=%v\n", i, v)
	}

}
