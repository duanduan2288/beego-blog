package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	usrename := this.GetSession("username")
	if nil == usrename {
		this.Redirect("/login", 302)
	}
	this.Data["username"] = usrename
	this.Layout = "layouts/main.tpl"
	this.TplName = "index.tpl"
}
