package routers

import (
	"github.com/astaxie/beego"
	"yaoqianshu/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
