package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	//"time"
	"yaoqianshu/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "register.tpl"
	//models.Model_test()
	username := this.GetString(":username")
	if username == "" {
		this.TplName = "register.tpl"
	} else {
		beego.Debug(fmt.Sprintf("%s", username))
		//c.Ctx.WriteString(username)
		var myStructData map[string]interface{}
		if username != "" {
			if models.Check_userName(username) {
				myStructData = map[string]interface{}{"success": 1, "message": "confirm", "username": username}
			} else {
				myStructData = map[string]interface{}{"success": 0, "message": "username in DB or WRONG", "username": username}
			}
		} else {
			myStructData = map[string]interface{}{"success": 0, "message": "username is NULL", "username": username}
		}
		this.Data["json"] = &myStructData
		this.ServeJSON()
	}
}

func (this *RegisterController) Post() {
	userName := this.GetString("username")
	passWord := this.GetString("password")
	vcode := this.GetString("vcode")

	if vcode != "" {

		code := models.Md5(string(models.Base64_encode([]byte(vcode))))
		session_code := this.GetSession("vcode")
		if code == session_code {
			if userName != "" && passWord != "" {
				reg_id := models.Add_user(userName, passWord)
				fmt.Printf("%d", reg_id)
				if reg_id == 0 {
					this.Ctx.WriteString("注册失败，请联系管理员")
				} else {
					//注册成功
					this.Ctx.WriteString("<html><head><meta http-equiv=\"Refresh\" content=\"3; url=./login\" /></head><body><h3>" + userName + "账号已注册成功！正在跳转至登陆页....." + "</h3></body></html>")
					//time.Sleep(3 * time.Second)
					//this.Ctx.Redirect(302, "./login")
				}
			} else {
				this.Ctx.WriteString("注册失败，请联系管理员")
			}
		} else {
			this.Ctx.WriteString("注册失败，验证码不正确")
		}

	} else {
		this.Ctx.WriteString("注册失败，验证码不正确")
	}

}
