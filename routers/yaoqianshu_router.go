package routers

import (
	"github.com/astaxie/beego"
	"yaoqianshu/controllers"
)

func init() {
	beego.Router("/index", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/register/:username", &controllers.RegisterController{})
	beego.Router("/register/:username/sms", &controllers.RegisterSMSController{})
	beego.Router("/project", &controllers.ProjectController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/:function", &controllers.UserController{})
	beego.Router("/user/:function/:action", &controllers.UserController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/aboutus", &controllers.AboutusController{})
}
