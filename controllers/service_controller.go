package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
	_ "yaoqianshu/models"
)

type ServiceController struct {
	beego.Controller
}

func (this *ServiceController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	path := this.Ctx.Request.URL.Path

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
	this.TplName = "service.tpl"
	//fmt.Println(c.Ctx.Request.URL.Path)
}
