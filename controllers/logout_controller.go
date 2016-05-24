package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Get() {
	path := this.Ctx.Request.URL.Path
	if path == "/" {
		path = "/index"
	}

	session := this.GetSession("username")
	session_status := this.GetSession("login")

	if session != nil && session_status != nil {
		this.DelSession("username")
		this.DelSession("login")
	} else {
		this.Ctx.WriteString("<html><body>您还没有登陆，点击登陆摇钱木又寸&nbsp;<a href=\"./login\">登陆</a></body></html>")
	}

	this.Data["Path"] = path
	this.Ctx.Redirect(302, "./index")
	//this.Ctx.WriteString(path)
}
