package controllers

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) TestGet() {
	c.Ctx.WriteString("<font color='green' size='6'>TestController。。。。。this is method get!</font>")
}

func (c *TestController) TestPost() {
	c.Ctx.WriteString("TestController。。。。。this is method post!")
}
