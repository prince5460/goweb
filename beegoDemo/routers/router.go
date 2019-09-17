package routers

import (
	"github.com/astaxie/beego"
	"goweb/beegoDemo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{},
		"get:Get;post:Post")

	beego.Router("/test", &controllers.TestController{},
		"get:TestGet;post:TestPost")

	beego.Router("/testinput", &controllers.TestInputController{},
		"get:TestInputGet;post:TestInputPost")

	beego.Router("/testlogin", &controllers.TestLoginController{},
		"get:Login;post:Post")
}
