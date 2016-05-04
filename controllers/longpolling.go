package controllers

import (
	"blog/models"
)

type LongPollingController struct {
	baseController
}

func (this *LongPollingController) Join() {
	uname := this.GetSession("username")
	if nil == uname {
		this.Redirect("/login", 302)
	}
	unameone := Parse(uname)
	Join(unameone, nil)

	this.Data["UserName"] = uname
	this.Layout = "layouts/main.tpl"
	this.TplName = "chatroom/longpolling.tpl"
}
func (this *LongPollingController) Post() {
	mystruct := make(map[string]string)
	uname := this.GetSession("username")
	if nil == uname {
		mystruct["info"] = "error"
		mystruct["message"] = "请先登录"
		this.Data["json"] = mystruct
		this.ServeJSON()
		return
	}
	unameone := Parse(uname)
	this.Data["UserName"] = unameone
	this.Layout = "layouts/main.tpl"
	this.TplName = "chatroom/longpolling.tpl"

	content := this.GetString("content")
	if len(content) == 0 {
		mystruct["info"] = "error"
		mystruct["message"] = "请输入要发送的内容"
		this.Data["json"] = mystruct
		this.ServeJSON()
		return
	}
	publish <- newEvent(models.EVENT_MESSAGE, unameone, content)
}

func (this *LongPollingController) Fetch() {
	mystruct := make(map[string]string)
	uname := this.GetSession("username")
	if nil == uname {
		mystruct["info"] = "error"
		mystruct["message"] = "请先登录"
		this.Data["json"] = mystruct
		this.ServeJSON()
		return
	}

	lastReceived, err := this.GetInt("lastReceived")

	if err != nil {
		mystruct["info"] = "error"
		mystruct["message"] = "参数错误"
		this.Data["json"] = mystruct
		this.ServeJSON()
		return
	}
	events := models.GetEvents(int(lastReceived))
	if len(events) > 0 {
		mystruct["info"] = "ok"
		mystruct["message"] = ""
		this.Data["json"] = events
		this.ServeJSON()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	this.Data["json"] = models.GetEvents(int(lastReceived))
	this.ServeJSON()
	return
}
