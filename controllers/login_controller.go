package controllers

import (
	"github.com/astaxie/beego"
	"yaoqianshu/models"
	//"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	this.TplName = "login.tpl"
}

func (this *LoginController) Post() {
	userName := this.GetString("username")
	passWord := this.GetString("password")

	if models.Login_user(userName, passWord) {
		this.SetSession("username", userName)
		this.SetSession("login", "1")
		//fmt.Println( this.GetSession("username").(string) + "这是用户名" + this.GetSession("login").(string) )
		this.Ctx.Redirect(302, "./index")
	} else {
		this.Ctx.WriteString("<html><body>用户名或密码错误，请联系管理员！<a href=\"./login\">点击重新登陆</a></body></html>")
	}
}
