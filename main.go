package main

import (
	"github.com/astaxie/beego"
	_ "yaoqianshu/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true //bug for beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
