package controllers

import (
	"github.com/astaxie/beego"
	"yaoqianshu/models"
)

type RegisterSMSController struct {
	beego.Controller
}

func (this *RegisterSMSController) Get() {
	username := this.GetString(":username")
	referer := this.Ctx.Request.Referer()
	if username != "" && referer != "" {
		vcode := models.Send_message(username)
		if vcode != "" {
			str := models.Md5(string(models.Base64_encode([]byte(vcode))))
			this.SetSession("vcode", str)
			this.Data["json"] = map[string]interface{}{"status": "1", "referer": referer}
		} else {
			this.Data["json"] = map[string]interface{}{"status": "0", "referer": referer}
		}
	} else {
		this.Data["json"] = map[string]interface{}{"status": "0", "referer": referer}
	}
	this.ServeJSON()
}
