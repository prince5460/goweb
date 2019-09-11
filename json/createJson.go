package main

import (
	"encoding/json"
	"fmt"
)

type Server2 struct {
	ServerName string
	ServerIP   string
}
type Serverslice struct {
	Servers []Server2
}

func main() {
	var s Serverslice
	s.Servers = append(s.Servers, Server2{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server2{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
