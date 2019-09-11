package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//Send：发送命令至缓冲区
//Flush：清空缓冲区，将命令一次性发送至服务器
//Recevie：依次读取服务器响应结果，当读取的命令未响应时，该操作会阻塞。
func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	fmt.Println("connect success ...")
	conn.Send("HSET", "user", "name", "lily", "age", "30")
	conn.Send("HSET", "user", "sex", "female")
	conn.Send("HGET", "user", "age")
	conn.Flush()

	res1, err := conn.Receive()
	fmt.Printf("Receive res1:%v \n", res1)
	res2, err := conn.Receive()
	fmt.Printf("Receive res2:%v\n", res2)
	res3, err := conn.Receive()
	fmt.Printf("Receive res3:%s\n", res3)
}
