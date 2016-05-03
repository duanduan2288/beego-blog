package controllers

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {

	db, _ := sql.Open("mysql", "root:@/blog")

	rows, _ := db.Query("select * from user")

	defer rows.Close()

	columns, _ := rows.Columns()

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//	list := make([]interface{}, len(columns))
	list := make([]map[string]string, 0)

	for rows.Next() {
		//将行数据保存到record字典
		rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {

				record[columns[i]] = string(col.([]byte))
				//this.Ctx.WriteString(string(col.([]byte)))
			}
		}
		list = append(list, record)
	}
	this.Data["list"] = list
	this.Layout = "layouts/main.tpl"
	this.TplName = "user/index.tpl"
}

//context.Input.IsPost
func (this *UserController) CreateUser() {

	if this.Ctx.Input.IsGet() {

		flash := beego.ReadFromRequest(&this.Controller)

		if n, ok := flash.Data["error"]; ok {
			this.Data["error"] = n
		} else {
			this.Data["error"] = ""
		}
		this.Layout = "layouts/main.tpl"
		this.TplName = "user/create.tpl"
	} else {
		flash := beego.NewFlash()

		email := this.GetString("username")
		if "" == email {

			flash.Error("请输入用户名")
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}
		password := this.GetString("password")

		if "" == password {
			flash.Error("请设置密码")
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}
		rpassword := this.GetString("rpassword")

		if "" == rpassword {
			flash.Error("请确认密码")
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}

		if rpassword != password {
			flash.Error("两次输入密码不一致")
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}

		mobile := this.GetString("mobile")

		db, err := sql.Open("mysql", "root:@/blog")

		if err != nil {
			flash.Error("连接数据库出错")
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}
		var dbpassword string
		db.QueryRow("select password as dbpassword from user where email=?", email).Scan(&dbpassword)

		if len(dbpassword) > 0 {
			flash.Error("用户名已存在")
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}
		h := md5.New()
		h.Write([]byte(password))
		md5password := hex.EncodeToString(h.Sum(nil))

		stmt, err := db.Prepare("INSERT INTO user(email,password,create_time,create_ip,mobile,last_ip) VALUES(?,?,?,?,?,?)")

		defer stmt.Close()

		if err != nil {

			flash.Error(err.Error())
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}
		create_time := time.Now()
		create_ip := this.Ctx.Request.Header.Get("X-Forwarded-For")

		_, err3 := stmt.Exec(email, md5password, create_time, create_ip, mobile, create_ip)
		if err3 != nil {
			flash.Error(err3.Error())
			flash.Store(&this.Controller)
			this.Redirect("/create", 302)
		}
		this.Redirect("/user/index", 302)

	}
}
func (this *UserController) Edit() {

}
