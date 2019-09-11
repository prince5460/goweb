package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

type RecurlyServers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []Server `xml:"server"`
	Description string   `xml:",innerxml"`
}

func main() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	v := RecurlyServers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("v:", v)
}
