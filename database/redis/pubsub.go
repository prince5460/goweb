package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

//发布订阅模式可以是多对多模式还可支持正则表达式，
// 发布者可以向一个或多个频道发送消息，
// 订阅者可订阅一个或者多个频道接受消息。

func Subs() { //订阅者
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
	psc := redis.PubSubConn{conn}
	psc.Subscribe("channel1") //订阅channel1频道
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Println(v)
			return
		}
	}
}

func Push(message string) { //发布者
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	_, err1 := conn.Do("PUBLISH", "channel1", message)
	if err1 != nil {
		fmt.Println("pub err: ", err1)
		return
	}

}

func main() {
	go Subs()
	go Push("this is jerry")
	time.Sleep(time.Second * 3)
}
