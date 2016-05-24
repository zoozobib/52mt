package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

/*验证用户名*/
func Check_userName(username string) bool {
	if username != "" {
		if Check_userName_length(username) && Check_userName_type(username) && Check_db_userName(username) {
			if Check_db_userName(username) {
				return true //可以注册
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

/*end by zoozobib*/

/*验证注册信息，为异步处理做准备 by zoozobib*/
func Check_db_userName(username string) bool {
	o := orm.NewOrm()
	//orm.Debug = true
	if username != "" {
		user := User{Name: username}
		err := o.Read(&user, "Name")
		fmt.Println(err)
		if err == orm.ErrNoRows {
			fmt.Println("查询不到")
			return true
		} else if err == orm.ErrMissPK {
			fmt.Println("找不到主键", user)
			return true
		} else {
			fmt.Println(user.Id, user.Name)
			return false
		}
	} else {
		return false
	}
}

/*end by zoozobib*/

/*验证注册信息，用户名是否符合规范-长度  1@1.com*/
func Check_userName_length(username string) bool {
	if username != "" {
		username_length := len(username)
		if username_length >= 7 && username_length <= 32 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

/*end by zoozobib*/

/*验证注册信息，用户名是否符合规范-手机号或email地址  regexp.MatchString("^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$", username) 13801380138 / 1@1.com*/
func Check_userName_type(username string) bool {
	if username != "" {
		if isOk, _ := regexp.MatchString("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$", username); isOk {
			return true
		} else if isOk, _ := regexp.MatchString("^0?(13[0-9]|15[012356789]|17[0-9]|18[0-9]|14[57])[0-9]{8}$", username); isOk {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

/*end by zoozobib*/

/*验证注册信息，密码是否符合规范*/
func Check_password_length(password string) bool {
	if password != "" {
		password_length := len(password)
		if password_length > 7 && password_length < 17 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

/*end by zoozobib*/

func Model_test() {
	o := orm.NewOrm()

	user := User{Name: "slene"}

	// insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

}

func Add_user(username, password string) int64 {
	if username != "" && password != "" && Check_userName(username) && Check_password_length(password) {
		o := orm.NewOrm()
		user := User{Name: username, Password: Md5(password), Mp: username}
		id, err := o.Insert(&user)
		if err != nil {
			return 0
		} else {
			return id
		}
	} else {
		return 0
	}
}

/*短信服务开启*/

type Msg struct {
	Success bool
	Id      string
}

func sms_http_post(sms_url string) {

	resp, err := http.Get(sms_url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	msg := string(body)
	fmt.Println("msg:", msg)

	var msg_s Msg
	json.Unmarshal([]byte(msg), &msg_s)
	fmt.Println(msg_s.Id)
	fmt.Println(msg_s.Success)
}

//http://222.73.117.138:7891/mt?un=star&pw=123456&da=13612345678&sm=【您的签名】您要发送的内容&dc=15&rd=1&tf=3&rf=2

func Send_message(da string) string {
	un := "N5195427"
	pw := "dfd6b011"
	dc := "15"
	rd := "1"
	tf := "3"
	rf := "2"

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	sm := "【摇钱木又寸】您的验证码为：" + vcode

	sms_url := fmt.Sprintf("http://222.73.117.138:7891/mt?un=%s&pw=%s&da=%s&sm=%s&dc=%s&rd=%s&tf=%s&rf=%s",
		un, pw, da, sm, dc, rd, tf, rf)

	fmt.Println(sms_url)
	sms_http_post(sms_url)
	return vcode
}
