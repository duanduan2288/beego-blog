package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"blog/models"
)

type WebSocketController struct {
	baseController
}

func (this *WebSocketController) Get() {
	usrename := this.GetSession("username")

	if nil == usrename {
		this.Redirect("/login", 302)
	}
	this.Data["UserName"] = usrename
	this.Data["IsWebSocket"] = true
	this.Layout = "layouts/main.tpl"
	this.TplName = "chatroom/websocket.tpl"
}

func (this *WebSocketController) Join() {
	usrename := this.GetSession("username")
	if nil == usrename {
		this.Redirect("/login", 302)
	}
	uname := Parse(usrename)
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	Join(uname, ws)
	defer Leave(uname)

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
	}
}
func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		// Immediately send event to WebSocket users.
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				// User disconnected.
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
func Parse(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case []string:
		data := i.([]string)
		return data[0]
	default:
		panic("type match miss")
	}
	return ""
}
