package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Layout = "layouts/main.tpl"
	c.TplName = "login.tpl"
}

func (this *LoginController) Post() {

	username := this.GetString("username")
	fpassword := this.GetString("password")
	if "" == username {
		this.Ctx.WriteString("请输入用户名")
		return
	}
	if "" == fpassword {
		this.Ctx.WriteString("请输入密码")
		return
	}
	db, err := sql.Open("mysql", "root:@/qiangdawei")
	if err != nil {
		this.Ctx.WriteString("连接数据库错误")
		return
	}
	var password string

	err2 := db.QueryRow("select password from dv_user where email=?", username).Scan(&password)

	if err2 == nil {
		h := md5.New()
		h.Write([]byte(fpassword))
		md5password := hex.EncodeToString(h.Sum(nil))

		if password != md5password {
			this.Ctx.WriteString("密码错误")
			return
		} else {
			//将登陆信息存到session
			this.SetSession("username", username)
			this.Redirect("/", 302)
		}
	} else {

		this.Ctx.WriteString(err2.Error())
		return
	}
}
