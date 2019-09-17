package main

import (
	"github.com/astaxie/beego"
	_ "goweb/beegoDemo/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
