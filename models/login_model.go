package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func Login_user(username, password string) bool {
	o := orm.NewOrm()

	if username != "" && password != "" {
		user := User{Name: username, Password: Md5(password)}
		err := o.Read(&user, "Name", "Password")
		fmt.Println(err)
		if err == orm.ErrNoRows {
			fmt.Println("查询不到")
			return false
		} else if err == orm.ErrMissPK {
			fmt.Println("找不到主键", user)
			return false
		} else {
			fmt.Println(user.Id, user.Name)
			return true
		}
	} else {
		return false
	}
}
