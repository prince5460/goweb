package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//MULTI：开启事务
//EXEC：执行事务
//DISCARD：取消事务
//WATCH：监视事务中的键变化，一旦有改变则取消事务。

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	defer conn.Close()

	conn.Send("MULTI")
	conn.Send("INCR", "foo")
	conn.Send("INCR", "bar")
	r, err := conn.Do("EXEC")
	fmt.Println(r)
}
