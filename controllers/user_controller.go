package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"yaoqianshu/models"
	//"reflect"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	path := this.Ctx.Request.URL.Path
	if path == "/" {
		path = "/index"
	}

	session := this.GetSession("username")
	session_status := this.GetSession("login")

	if session != nil && session_status != nil {
		this.Data["session"] = session.(string)
		this.Data["session_status"] = session_status.(string)

		function := this.GetString(":function")
		action := this.GetString(":action")

		if function != "" {
			this.Data["function"] = function
			//user_profile := models.Get_profile(session.(string))
			user_profile := map[string]string{"mp": session.(string)}
			this.Data["today_money"] = models.Get_today_money(user_profile["mp"])
			this.Data["current_money"] = models.Get_current_money(user_profile["mp"])
			this.Data["week_money"] = models.Get_weekend_money(user_profile["mp"])
			this.Data["month_money"] = models.Get_month_money(user_profile["mp"])
			this.Data["user_money_list"] = models.Get_user_money_list(user_profile["mp"])
			//fmt.Println(models.Get_month_money(user_profile["mp"]))
		} else {
			this.Data["function"] = ""
			//user_profile := models.Get_profile(session.(string))
			user_profile := map[string]string{"mp": session.(string)}
			//fmt.Println(models.Get_month_money(user_profile["mp"]))
			this.Data["today_money"] = models.Get_today_money(user_profile["mp"])
			this.Data["current_money"] = models.Get_current_money(user_profile["mp"])
			this.Data["week_money"] = models.Get_weekend_money(user_profile["mp"])
			this.Data["month_money"] = models.Get_month_money(user_profile["mp"])
			this.Data["user_money_list"] = models.Get_user_money_list(user_profile["mp"])
		}
		//开始获取账号user_money
		if function == "money" {
			this.Data["account"], this.Data["account_name"] = models.Get_account(session.(string))
			//user_profile := models.Get_profile(session.(string))
			user_profile := map[string]string{"mp": session.(string)}
			user_id, _ := strconv.Atoi(models.Get_profile(session.(string))["id"])
			this.Data["user_total_money"] = models.Get_user_total_money(user_profile["mp"])
			//this.Data["user_effect_money"] = models.Get_user_effect_money(user_profile["mp"])
			this.Data["user_effect_money"] = models.Query_remain_cash_integral(user_id)
			this.Data["user_nums"] = models.Get_nums(user_id)
			//fmt.Println(models.Get_nums(user_profile["mp"]))
		}
		//end

		//开始积分转换现金，提交积分信息
		if function == "money" && action == "get" {
			user_id, _ := strconv.Atoi(models.Get_profile(session.(string))["id"])
			user_profile := map[string]string{"mp": session.(string)}
			exchange_number := this.GetString("n")
			var myStruct map[string]interface{}
			en := models.Cash_exchange_integral(user_id, user_profile["mp"], exchange_number)
			if en {
				myStruct = map[string]interface{}{"success": "1"}
			} else {
				myStruct = map[string]interface{}{"success": "0"}
			}
			this.Data["json"] = &myStruct
			this.ServeJSON()
		}
		//end by zoozobib

		//开始获取账号user_profile
		if function == "profile" {
			/*
				id,mp,email := models.Get_profile( session.(string) )

				this.Data["id"] = id
				this.Data["mp"] = mp
				this.Data["email"] = email
			*/
			//var user interface{}
			user := models.Get_profile(session.(string))
			//v := reflect.ValueOf(users)
			this.Data["id"] = user["id"]
			this.Data["mp"] = user["mp"]
			this.Data["email"] = user["email"]
			//fmt.Println(user)
		}

		if function == "invest" {
			goods_tips := models.Get_invest_tips(session.(string))
			this.Data["tips"] = goods_tips
		}

		if function == "invest" && action == "add" {
			this.Data["action"] = "add"
		} else {
			this.Data["action"] = ""
		}
		//end
		this.TplName = "user.tpl"
	} else {
		this.Data["session"] = ""
		this.Data["session_status"] = ""
		this.Ctx.WriteString("<html><body>您还没有登陆，点击登陆摇钱木又寸&nbsp;<a href=\"/login\">登陆</a></body></html>")
	}

	this.Data["Path"] = path
	//this.Ctx.Redirect(302, "./index")
	//this.Ctx.WriteString(path)
}

func (this *UserController) Post() {
	path := this.Ctx.Request.URL.Path
	if path == "/" {
		path = "/index"
	}

	session := this.GetSession("username")
	session_status := this.GetSession("login")
	function := this.GetString(":function")
	action := this.GetString(":action")

	if session != nil && session_status != nil { //判断有没有登陆
		//fmt.Println(session.(string), account)
		if function == "money" { //user/money
			account := this.GetString("alipay_account")
			account_name := this.GetString("alipay_account_name")
			if account != "" && account_name != "" { //参数是否提交
				if models.Save_alipay_account(session.(string), account, account_name) { //是否保存正确
					this.Ctx.Redirect(302, path)
				} else {
					this.Ctx.WriteString("<html><body>更新提现帐户失败，请联系管理员！<a href=\"./money\">返回</a></body></html>")
				}
			} else {
				this.Ctx.WriteString("<html><body>账号信息不完整，请重新提交！<a href=\"./money\">返回</a></body></html>")
			}
		}

		if function == "profile" { //user/profile
			password := this.GetString("password")
			re_password := this.GetString("re_password")
			mp := this.GetString("mp")
			email := this.GetString("email")
			if password != re_password {
				this.Ctx.WriteString("参数有问题，密码不正确！")
			} else {
				if password == "" && re_password == "" {
					if models.Save_profile_no_password(session.(string), mp, email) {
						this.Ctx.Redirect(302, path)
					} else {
						this.Ctx.WriteString("参数有问题")
					}
				} else {
					if models.Save_profile(session.(string), password, mp, email) {
						this.Ctx.Redirect(302, path)
					} else {
						this.Ctx.WriteString("参数有问题")
					}
				}
			}
		}

		if function == "invest" && action == "add" { //user/invest/add
			this.Data["action"] = "add"
			site_username := this.GetString("site_username")
			site_name := this.GetString("site_name")
			site_mp := this.GetString("site_mp")
			invest_time := this.GetString("invest_time")
			product_name := this.GetString("product_name")
			if product_name == "" {
				product_name = ""
			}
			invest_amount := this.GetString("invest_amount")
			//invest_screen_shot := this.GetString("invest_screen_shot")

			if site_username != "" && site_name != "" && site_mp != "" && invest_time != "" && invest_amount != "" {
				f, h, err := this.GetFile("invest_screen_shot")
				file_name := h.Filename
				fmt.Println("file_name", file_name)
				defer f.Close()
				if err != nil {
					fmt.Println("getfile err ", err)
				} else {
					this.SaveToFile("invest_screen_shot", "./static/uploads/img/"+file_name)
					fmt.Println("file Upload ok!")
				}

				save_invest_tips_id := models.Save_invest_tips(session.(string), site_username, site_name, site_mp, invest_time, product_name, invest_amount, "/static/uploads/img/"+file_name, "0")

				if save_invest_tips_id != 0 {
					this.Ctx.Redirect(302, "./invest")
				} else {
					this.Ctx.WriteString("保存任务表单有错误，请联系管理员!!!")
				}

			} else {
				this.Ctx.WriteString("任务表单参数有错误，请联系管理员!!!")
			}

		} else {
			this.Data["action"] = ""
		}

	} else {
		this.Data["session"] = ""
		this.Data["session_status"] = ""
		this.Ctx.WriteString("<html><body>您还没有登陆，点击登陆摇钱木又寸&nbsp;<a href=\"/login\">登陆</a></body></html>")
	}

	this.Data["Path"] = path
}
