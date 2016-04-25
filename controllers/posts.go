package controllers

import (
	"github.com/astaxie/beego"
)

type PostsController struct {
	beego.Controller
}

//列表
func (this *PostsController) Index() {

}

func (this *PostsController) Create() {
	if this.Ctx.Input.IsGet() {
		this.Layout = "layouts/main.tpl"
		this.TplName = "posts/index.tpl"
	} else {
		title := this.GetString("title")
		content := this.GetString("content")
	}
}

func (this *PostsController) Edit() {

}

func (this *PostsController) Delete() {

}
