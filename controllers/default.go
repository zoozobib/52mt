package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	path := this.Ctx.Request.URL.Path
	if path == "/" {
		path = "/index"
	}

	session := this.GetSession("username")
	session_status := this.GetSession("login")

	if session != nil && session_status != nil {
		this.Data["session"] = session.(string)
		this.Data["session_status"] = session_status.(string)
	} else {
		this.Data["session"] = ""
		this.Data["session_status"] = ""
	}

	this.Data["Path"] = path
	this.TplName = "index.tpl"
	//this.Ctx.WriteString(path)
}
