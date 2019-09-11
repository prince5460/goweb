package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	SreverIp   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var s ServerSlice
	jsonStr := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(jsonStr), &s)
	fmt.Println(s)
}
