package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool
var Pool *redis.Pool

//当程序启动时初始化连接池
func init() { //init 用于初始化一些参数，先于main执行
	Pool = &redis.Pool{
		MaxIdle:     16,  //最大空闲连接数
		MaxActive:   32,  //客户端与数据库的最大连接数,0表示没有限制
		IdleTimeout: 120, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	conn := Pool.Get()
	defer conn.Close()

	res, err := conn.Do("HSET", "user", "name", "Lucy")
	fmt.Println(res, err)
	res1, err := redis.String(conn.Do("HGET", "user", "name"))
	fmt.Printf("res:%s,error:%v", res1, err)

	_, err = conn.Do("Set", "job", "coder")
	if err != nil {
		fmt.Println("Set err=", err)
		return
	}
	r, err := redis.String(conn.Do("Get", "job"))
	if err != nil {
		fmt.Println("Get err=", err)
		return
	}
	fmt.Println("r=", r)

}
