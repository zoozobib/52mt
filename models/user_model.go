package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
	"time"
	//"strings"
)

func Get_account(username string) (string, string) {
	if username != "" {
		user := User{Name: username}
		o := orm.NewOrm()
		if o.Read(&user, "Name") == nil {
			return user.Account, user.Account_name
		} else {
			return "", ""
		}
	} else {
		return "", ""
	}
}

func Save_alipay_account(username, account, account_name string) bool {
	if username != "" && account != "" && account_name != "" {
		o := orm.NewOrm()
		user := User{Name: username}
		if o.Read(&user, "Name") == nil {
			user.Account = account
			user.Account_name = account_name
			if _, err := o.Update(&user, "Account", "Account_name"); err == nil {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func Save_profile(username string, args ...string) bool {
	if len(args) > 0 {
		o := orm.NewOrm()
		user := User{Name: username}
		if o.Read(&user, "Name") == nil {
			user.Password = Md5(args[0])
			user.Mp = args[1]
			user.Email = args[2]
			if _, err := o.Update(&user, "Password", "Mp", "Email"); err == nil {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func Save_profile_no_password(username string, args ...string) bool {
	if len(args) > 0 {
		o := orm.NewOrm()
		user := User{Name: username}
		if o.Read(&user, "Name") == nil {
			user.Mp = args[0]
			user.Email = args[1]
			if _, err := o.Update(&user, "Mp", "Email"); err == nil {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func Get_profile(username string) map[string]string {
	if username != "" {
		user := User{Name: username}
		o := orm.NewOrm()
		if o.Read(&user, "Name") == nil {
			users := map[string]string{"id": strconv.Itoa(user.Id), "mp": user.Mp, "email": user.Email}
			return users
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func Save_invest_tips(username string, args ...string) int64 {
	if len(args) > 0 {
		o := orm.NewOrm()
		user := User{Name: username}
		if o.Read(&user, "Name") == nil {
			user_id := user.Id

			site_username := args[0]
			site_name := args[1]
			site_mp := args[2]
			invest_time := args[3]
			product_name := args[4]
			invest_amount := args[5]
			invest_screen_shot := args[6]
			status := args[7]

			invest_tips := Goods_tip{Site_username: site_username, Site_name: site_name, Site_mp: site_mp, Invest_time: invest_time, Product_name: product_name, Invest_amount: invest_amount, Screenshot_url: invest_screen_shot, Uid: user_id, Status: status}

			id, err := o.Insert(&invest_tips)
			if err != nil {
				fmt.Println(err)
				return 0
			} else {
				return id
			}

		} else {
			return 0
		}
	} else {
		return 0
	}
}

func Get_invest_tips(username string) []Goods_tip {
	if username != "" {
		o := orm.NewOrm()
		user := User{Name: username}
		if o.Read(&user, "Name") == nil {

			user_id := user.Id

			var gt []Goods_tip
			o.QueryTable(new(Goods_tip)).Filter("Uid", user_id).All(&gt)
			return gt
		} else {
			return nil
		}
	} else {
		return nil
	}
}

/*最近一次收益金额和积分*/
func Get_current_money(mp string) []orm.Params {
	if mp != "" {
		var cus []orm.Params
		o := orm.NewOrm()
		num, err := o.Raw("SELECT round(sum(income), 2) income, round(sum(integral), 2) integral FROM `cus_account` T0 WHERE T0.`mobile_phone_number` LIKE ?", Change_strings(mp)).Values(&cus)
		if err == nil && num > 0 {
			//fmt.Println(cus)
			return cus
		} else {
			return nil
		}
	} else {
		return nil
	}
}

/*end*/

/*今天收益和积分 SELECT * FROM cus_account WHERE	mobile_phone_number LIKE '137%%5491' AND to_days(first_buy_time) = to_days(now())*/
func Get_today_money(mp string) []orm.Params {
	if mp != "" {
		var cus []orm.Params
		o := orm.NewOrm()
		num, err := o.Raw("SELECT round(sum(income), 2) income, round(sum(integral), 2) integral FROM cus_account where mobile_phone_number LIKE ? AND to_days(first_buy_time) = to_days(now())", Change_strings(mp)).Values(&cus)
		if err == nil && num > 0 {
			return cus
		} else {
			fmt.Println(err, num, cus)
			return nil
		}
	} else {
		return nil
	}
}

/*end*/

/*7日 SELECT c_id,mobile_phone_number,reg_time,first_buy_time,first_buy_amount,income,integral FROM cus_account WHERE	mobile_phone_number = '189****3704' DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(first_buy_time);*/
func Get_weekend_money(mp string) []orm.Params {
	if mp != "" {
		var cus []orm.Params
		o := orm.NewOrm()
		num, err := o.Raw("SELECT round(sum(income), 2) income, round(sum(integral), 2) integral FROM cus_account WHERE mobile_phone_number LIKE ? AND DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(first_buy_time);", Change_strings(mp)).Values(&cus)
		if err == nil && num > 0 {
			return cus
		} else {
			return nil
		}
	} else {
		return nil
	}
}

/*end*/

/*月 AND DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(first_buy_time);*/

func Get_month_money(mp string) []orm.Params {
	if mp != "" {
		var cus []orm.Params
		o := orm.NewOrm()
		num, err := o.Raw("SELECT round(sum(income), 2) income, round(sum(integral), 2) integral FROM cus_account WHERE	mobile_phone_number LIKE ? AND DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(first_buy_time);", Change_strings(mp)).Values(&cus)
		if err == nil && num > 0 {
			return cus
		} else {
			return nil
		}
	} else {
		return nil
	}
}

/**/

/*取当前用户所有的收益记录*/
func Get_user_money_list(mp string) []orm.Params {
	if mp != "" {
		var cus []orm.Params
		o := orm.NewOrm()
		num, err := o.Raw("select * from cus_account where mobile_phone_number like ?", Change_strings(mp)).Values(&cus)
		if err == nil && num > 0 {
			return cus
		} else {
			return nil
		}
	} else {
		return nil
	}
}

/*end by zoozobib*/

/*取个人可取现记录 select round(sum(income), 2) income, round(sum(integral), 2) integral from cus_account where mobile_phone_number like '137%%5491'*/
func Get_user_total_money(mp string) string {
	if mp != "" {
		var money []orm.Params
		o := orm.NewOrm()
		_, err := o.Raw("select round(sum(integral), 2) integral from cus_account where mobile_phone_number like ?", Change_strings(mp)).Values(&money)
		if err == nil {
			if money[0]["integral"] == nil {
				return "0"
			} else {
				moneys := string(money[0]["integral"].(string))
				return moneys
			}
		} else {
			fmt.Println(err, "err")
			return "0"
		}
	} else {
		return "0"
	}
}

/**/

/*取个人可取现记录 select round(sum(income), 2) income, round(sum(integral), 2) integral from cus_account where mobile_phone_number like '137%%5491'*/
func Get_user_effect_money(mp string) string {
	if mp != "" {
		var money []orm.Params
		o := orm.NewOrm()
		//SELECT sum(invest_amount) FROM	goods_tip a,	USER b WHERE	b.mp LIKE '152%1741' AND b.id = a.uid AND a. STATUS = '1';
		_, err := o.Raw("SELECT sum(invest_amount) integral FROM	goods_tip a,	USER b WHERE	b.mp LIKE ? AND b.id = a.uid AND a. STATUS = '1'", Change_strings(mp)).Values(&money)
		if err == nil {
			if money[0]["integral"] == nil {
				return "0"
			} else {
				moneys := string(money[0]["integral"].(string))
				return moneys
			}
		} else {
			fmt.Println(err, "err")
			return "0"
		}
	} else {
		return "0"
	}
}

/**/

//处理积分问题
func Cash_exchange_integral(uid int, mp string, in_ string) bool {
	if len(in_) == 0 || len(mp) < 11 {
		return false
	}

	mp1 := mp[:3] + "%" + mp[7:11]
	//mp2 := mp[:3] + "****" + mp[7:11]

	c_in, err := strconv.Atoi(in_)
	if err != nil {
		fmt.Println("string to int fail, ", err)
		return false
	}

	fmt.Println(c_in)

	o := orm.NewOrm()

	//查询总积分

	var data1 []orm.Params
	num, MQErr1 := o.Raw("SELECT  sum(integral) integral FROM cus_account WHERE mobile_phone_number like ? GROUP BY mobile_phone_number", mp1).Values(&data1)

	if MQErr1 != nil {
		fmt.Println("mysql error, msg:", MQErr1)
		return false
	}

	if num == 0 {
		fmt.Println("cus_account is not data, mp is :", mp1)
		return false
	}

	ti := data1[0]["integral"].(string)

	i_ti, err := strconv.Atoi(ti)
	if err != nil {
		fmt.Println("string to int fail, ", err)
		return false
	}

	fmt.Println("总积分:", i_ti)

	//查询历史提取积分
	var data2 []orm.Params
	num, MQErr2 := o.Raw("SELECT sum(cash_integral) cash_integral FROM integral_detail WHERE uid = ? AND status = '1'", uid).Values(&data2)

	fmt.Println("num : ", num, " data2:")
	if MQErr2 != nil {
		fmt.Println("mysql error, msg:", MQErr2)
		return false
	}

	var remain_in int
	//第一次提取积分
	if data2[0]["cash_integral"] == nil {
		remain_in = i_ti - c_in
	} else {
		ci := data2[0]["cash_integral"].(string)
		i_ci, err := strconv.Atoi(ci)
		if err != nil {
			fmt.Println("string to int fail, ", err)
			return false
		}
		fmt.Println("历史提取积分:", i_ci)
		remain_in = i_ti - c_in - i_ci
	}

	if remain_in < 0 {
		fmt.Println("提取积分大于总积分")
		return false
	}

	//插入提取积分明细

	add := Integral_detail{Uid: uid, Mp: mp, Total_integral: i_ti, Cash_integral: c_in,
		Remain_integral: remain_in, Create_time: time.Now().Format("2006-01-02 15:04:05"), Status: 0}

	id, MQErr3 := o.Insert(&add)
	if MQErr3 != nil {
		fmt.Println("mysql error, msg:", MQErr3)
		return false
	}

	fmt.Println("id : ", id)

	return true
}

//end by zoozobib

//提现记录
func Get_nums(uid int) []orm.Params {

	o := orm.NewOrm()

	//	mp1 := mp[:3] + "%" + mp[7:11]

	var data []orm.Params
	num, MQErr2 := o.Raw("SELECT * FROM integral_detail WHERE uid = ?", uid).Values(&data)

	fmt.Println("num : ", num, " data:", data)
	if MQErr2 != nil {
		fmt.Println("mysql error, msg:", MQErr2)
		return nil
	}
	return data
}

/*end by zoozobib*/

/*取可提现积分*/
func Query_remain_cash_integral(uid int) int {

	o := orm.NewOrm()
	//查询goods_tip
	var data1 []orm.Params
	num, MQErr1 := o.Raw("SELECT sum(invest_amount) invest_amount FROM goods_tip WHERE uid = ? AND status = '1'", uid).Values(&data1)

	fmt.Println("num : ", num, " data1:", data1)
	if MQErr1 != nil {
		fmt.Println("mysql error, msg:", MQErr1)
		return 0
	}
	var integral int = 0
	//	var i_invest float64 = 0.00

	if data1[0]["invest_amount"] == nil {
		return 0
	} else {
		invest := data1[0]["invest_amount"].(string)
		i_invest, err := strconv.ParseFloat(invest, 64)
		if err != nil {
			fmt.Println("string to int fail, ", err)
			return 0
		}
		//		fmt.Println("i_invest:", i_invest)
		integral = int(i_invest*0.02) * 10000
	}

	//查询剩余积分
	var data2 []orm.Params
	num, MQErr2 := o.Raw("SELECT sum(cash_integral)  cash_integral FROM integral_detail WHERE uid = ? AND status = '1' ", uid).Values(&data2)

	fmt.Println("num : ", num, " data2:", data2)
	if MQErr2 != nil {
		fmt.Println("mysql error, msg:", MQErr2)
		return integral
	}

	if data2[0]["cash_integral"] == nil {
		return integral
	} else {
		c_i := data2[0]["cash_integral"].(string)
		i_c_i, err := strconv.Atoi(c_i)
		if err != nil {
			fmt.Println("string to int fail, ", err)
			return integral
		}
		return integral - i_c_i
	}

	//return 0
}

/*end by chenliang*/
