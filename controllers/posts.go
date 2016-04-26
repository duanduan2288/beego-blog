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
	beego.ReadFromRequest(&this.Controller)
	this.Data["list"] = list
	this.Layout = "layouts/main.tpl"
	this.TplName = "posts/index.tpl"
}

/**
*
*创建文章
 */
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

//文章详情
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

//编辑文章
func (this *PostsController) Edit() {
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
	this.Data["id"] = id
	this.Data["title"] = title
	this.Data["content"] = content
	this.Data["create_time"] = create_time
	this.Layout = "layouts/main.tpl"
	this.TplName = "posts/edit.tpl"

}

func (this *PostsController) SaveEdit() {
	flash := beego.NewFlash()
	id := this.GetString("id")
	title := this.GetString("title")
	content := this.GetString("content")

	db, err := sql.Open("mysql", "root:@/blog")

	if nil != err {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/posts/edit/"+id, 302)
	}
	stmt, err1 := db.Prepare("UPDATE posts SET title=?,content=? WHERE id=?")
	if nil != err1 {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/posts/edit/"+id, 302)
	}

	defer stmt.Close()

	_, err2 := stmt.Exec(title, content, id)
	if nil != err2 {
		flash.Error(err2.Error())
		flash.Store(&this.Controller)
		this.Redirect("/posts/edit/"+id, 302)
	}

	this.Redirect("/posts/index", 302)

}

//删除文章
func (this *PostsController) Delete() {

	mystruct := make(map[string]string)
	id := this.GetString("id")
	status := this.GetString("status")
	if "" == id || "" == status {

		mystruct["info"] = "error"
		mystruct["message"] = "参数错误"
		this.Data["json"] = &mystruct
		this.ServeJSON()
	}
	db, err := sql.Open("mysql", "root:@/blog")
	if nil != err {

		mystruct["info"] = "error"
		mystruct["message"] = "数据库连接出错"
		this.Data["json"] = &mystruct
		this.ServeJSON()
	}
	stmt, err1 := db.Prepare("UPDATE posts SET status=? where id=?")
	if nil != err1 {

		mystruct["info"] = "error"
		mystruct["message"] = "操作失败"
		this.Data["json"] = &mystruct
		this.ServeJSON()
	}
	_, err3 := stmt.Exec(status, id)

	if nil != err3 {

		mystruct["info"] = "error"
		mystruct["message"] = "操作失败"
		this.Data["json"] = &mystruct
		this.ServeJSON()
	}
	mystruct["info"] = "ok"
	mystruct["message"] = "操作成功"
	this.Data["json"] = &mystruct
	this.ServeJSON()
}
