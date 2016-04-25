package controllers

import (
	"database/sql"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type PostsController struct {
	beego.Controller
}

//列表
func (this *PostsController) Index() {

	list := make([]map[string]string, 0)
	flash := beego.NewFlash()

	db, err := sql.Open("mysql", "root:@/blog")
	if nil != err {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
	} else {
		rows, _ := db.Query("select * from posts")

		defer rows.Close()

		columns, _ := rows.Columns()

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		//	list := make([]interface{}, len(columns))

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
	}
	this.Data["list"] = list
	this.Layout = "layouts/main.tpl"
	this.TplName = "posts/index.tpl"
}

func (this *PostsController) Create() {
	if this.Ctx.Input.IsGet() {
		this.Layout = "layouts/main.tpl"
		this.TplName = "posts/create.tpl"
	} else {
		flash := beego.NewFlash()

		title := this.GetString("title")
		content := this.GetString("content")

		db, err := sql.Open("mysql", "root:@/blog")

		if nil != err {
			flash.Error(err.Error())
			flash.Store(&this.Controller)
			this.Redirect("/posts/create", 302)
		}
		stmt, err1 := db.Prepare("INSERT INTO posts (title,content,create_time) VALUES(?,?,?)")
		if nil != err1 {
			flash.Error(err.Error())
			flash.Store(&this.Controller)
			this.Redirect("/posts/create", 302)
		}
		create_time := time.Now()

		defer stmt.Close()

		_, err2 := stmt.Exec(title, content, create_time)
		if nil != err2 {
			flash.Error(err2.Error())
			flash.Store(&this.Controller)
			this.Redirect("/posts/create", 302)
		}

		this.Redirect("/posts/index", 302)
	}
}

func (this *PostsController) Detail() {

	flash := beego.NewFlash()
	id := this.Ctx.Input.Param(":id")
	if "" == id {
		flash.Error("参数为空")
		flash.Store(&this.Controller)
		this.Redirect("/posts/index", 302)
	}
	db, err := sql.Open("mysql", "root:@/blog")

	if nil != err {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/posts/index", 302)
	}
	var title string
	var content string
	var create_time string

	err1 := db.QueryRow("select title,content,create_time from posts where id=?", id).Scan(&title, &content, &create_time)
	if nil != err1 {
		flash.Error(err1.Error())
		flash.Store(&this.Controller)
		this.Redirect("/posts/index", 302)
	}
	this.Data["title"] = title
	this.Data["content"] = content
	this.Data["create_time"] = create_time
	this.Layout = "layouts/main.tpl"
	this.TplName = "posts/detail.tpl"
}
func (this *PostsController) Edit() {

}

func (this *PostsController) Delete() {

}
