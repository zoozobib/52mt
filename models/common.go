package models

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"time"
)

type User struct { //用户表
	Id           int    `orm:"pk;auto"`
	Name         string `orm:"size(100)"`
	Password     string `orm:"size(100)"`
	Account      string `orm:"size(100)"`
	Account_name string `orm:"size(100)"`
	Mp           string `orm:"size(100)"`
	Email        string `orm:"size(100)"`
}

type Goods_tip struct { //商品任务表
	Gid            int    `orm:"pk;auto"`
	Uid            int    `orm:"size(4)"`
	Site_username  string `orm:"size(200)"`
	Site_name      string `orm:"size(200)"`
	Site_mp        string `orm:"size(200)"`
	Product_name   string `orm:"size(200);null"`
	Screenshot_url string `orm:"size(200)"`
	Invest_amount  string `orm:"size(50)"`
	Invest_time    string `orm:"size(50)"`
	Status         string `orm:"size(1);default(0)"`
}

/*
type Cus_account struct {
	Id int `orm:"pk;auto"`
  C_id string `orm:"size(100)"`
  M_user_id string `orm:"size(100)"`
  C_user_id string `orm:"size(100)"`
  User_name string `orm:"size(100)"`
  Mobile_phone_number string `orm:"size(100)"`
  Create_time time.Time
  Reg_time string `orm:"size(100)"`
  Is_name_auth string `orm:"size(100)"`
  Is_mobile_auth string `orm:"size(100)"`
  First_invest_type string `orm:"size(100)"`
  First_buy_time string `orm:"size(100)"`
  First_buy_amount string `orm:"size(100)"`
  Total_buy_current_amount string `orm:"size(100)"`
  Total_buy_regular_amount string `orm:"size(100)"`
  Protocol_des string `orm:"size(100)"`
  Integral int `orm:"size(4)"`
  Income float64 `orm:"size(8)"`
}
*/
/*
type Cus_account struct {
	M_user_id                string `orm:"size(100)"`
	Mobile_phone_number      string `orm:"size(100)"`
	C_user_id                string `orm:"size(100)"`
	User_name                string `orm:"size(100)"`
	Reg_time                 string `orm:"size(100)"`
	Is_name_auth             string `orm:"size(100)"`
	Is_mobile_auth           string `orm:"size(100)"`
	First_invest_type        string `orm:"size(100)"`
	First_buy_time           string `orm:"size(100)"`
	First_buy_amount         string `orm:"size(100)"`
	Total_buy_current_amount string `orm:"size(100)"`
	Total_buy_regular_amount string `orm:"size(100)"`
	Protocol_des             string `orm:"size(100)"`
	Create_time              time.Time
	Id                       int `orm:"pk;auto"`
	C_id                     int `orm:"size(4)"`
}
*/

type Cus_account struct {
	Id                       int     `orm:"pk;auto"`
	C_id                     string  `orm:"size(100)"`
	M_user_id                string  `orm:"size(100)"`
	C_user_id                string  `orm:"size(100)"`
	User_name                string  `orm:"size(100)"`
	Mobile_phone_number      string  `orm:"size(100)"`
	Create_time              string  `orm:"size(100)"`
	Reg_time                 string  `orm:"size(100)"`
	Is_name_auth             string  `orm:"size(100)"`
	Is_mobile_auth           string  `orm:"size(100)"`
	First_invest_type        string  `orm:"size(100)"`
	First_buy_time           string  `orm:"size(100)"`
	First_buy_amount         string  `orm:"size(100)"`
	Total_buy_current_amount string  `orm:"size(100)"`
	Total_buy_regular_amount string  `orm:"size(100)"`
	Protocol_des             string  `orm:"size(100)"`
	Integral                 int     `orm:"size(4)"`
	Income                   float64 `orm:"size(8)"`
}

type Integral_detail struct {
	Id              int    `orm:"pk;auto"`
	Uid             int    `orm:"size(4)"`
	Mp              string `orm:"size(100)"`
	Total_integral  int    `orm:"size(4)"`
	Cash_integral   int    `orm:"size(4)"`
	Remain_integral int    `orm:"size(4)"`
	Create_time     string `orm:"size(100)"`
	Status          int    `orm:"size(1)"`
}

/*数据库默认注册*/
func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "yaoqianshu:yaoqianshu@/yaoqianshu?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Goods_tip))
	orm.RegisterModel(new(Cus_account))
	orm.RegisterModel(new(Integral_detail))

	// create table
	orm.RunSyncdb("default", false, true)

	orm.Debug = true

}

/*end by zoozobib*/

/*md5加密方法*/
func Md5(str string) string {
	if str != "" {
		s := md5.New()
		s.Write([]byte(str))
		ss := s.Sum(nil)
		return hex.EncodeToString(ss)
	} else {
		return ""
	}
}

/*end by zoozobib*/

/**/
func Change_strings(str string) string {
	var left_p, right_p string

	if len(str) != 11 {
		return ""
	}

	if str != "" {
		for i := 0; i < 3; i++ {
			left_p += string(str[i])
		}

		for i := 7; i < 11; i++ {
			right_p += string(str[i])
		}

		return left_p + fmt.Sprintf("%s", "%%") + right_p

	} else {
		return ""
	}
}

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

/*base64加密*/

func Base64_encode(src []byte) []byte {
	var coder = base64.NewEncoding(base64Table)
	return []byte(coder.EncodeToString(src))
}

/*end*/

/*base64解密*/
func Base64_decode(src []byte) ([]byte, error) {
	var coder = base64.NewEncoding(base64Table)
	return coder.DecodeString(string(src))
}

/*end*/
