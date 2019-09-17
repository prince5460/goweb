package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type TestInputController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string `form:"pwd"`
}

func (c *TestInputController) TestInputGet() {
	//id := c.GetString("id")
	//c.Ctx.WriteString(id)

	//name := c.GetString("name")
	//c.Ctx.WriteString(name)
	//
	//idStr := c.Input().Get("id")
	//c.Ctx.WriteString(idStr)

	//读取session
	username := c.GetSession("username")
	password := c.GetSession("password")
	if nameString, ok := username.(string); ok && nameString != "" {
		c.Ctx.WriteString("Username:" + username.(string) + ",Password:" + password.(string))
	} else {
		c.Ctx.WriteString(
			`<html>
					<form action="http://127.0.0.1:9527/testinput" method="post">
						用户名：<input type ="text" name="Username" />
						<br/>
						密&nbsp&nbsp&nbsp码：<input type="password" name="pwd">
						<br/>
						<input type="submit" value="提交">
					</form></html>`)
	}
}

func (c *TestInputController) TestInputPost() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		log.Panic(err)
	}
	c.Ctx.WriteString("Username:" + u.Username + ",Password:" + u.Password)
}
