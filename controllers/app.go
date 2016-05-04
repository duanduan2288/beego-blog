package controllers

import (
	"github.com/astaxie/beego"
)

func init() {

}

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {
	this.Data["lang"] = "zh-CN"
}

func Parse(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case []string:
		data := i.([]string)
		return data[0]
	default:
		panic("类型错误")
	}
}

type AppController struct {
	baseController
}

func (this *AppController) Get() {

	usrename := this.GetSession("username")

	if nil == usrename {
		this.Redirect("/login", 302)
	}

	this.Data["username"] = usrename
	this.Layout = "layouts/main.tpl"
	this.TplName = "index.tpl"
}

//join chatroom
func (this *AppController) Join() {

	uname := this.GetSession("username")

	tech := this.GetString("tech")

	if nil == uname {
		this.Redirect("/", 302)
		return
	}

	switch tech {
	case "longpolling":
		this.Redirect("/lp", 302)

	case "websocket":
		this.Redirect("/ws", 302)
	default:
		this.Redirect("/", 302)

	}
	return
}
