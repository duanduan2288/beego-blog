package controllers

import (
	"container/list"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"

	"blog/models"
)

type Subscription struct {
	Archive []models.Event
	New     <-chan models.Event
}

func newEvent(ep models.EventType, user, msg string) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg}
}

func Join(user string, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	unsubscribe <- user
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn
}

var (
	subscribe   = make(chan Subscriber, 10)
	unsubscribe = make(chan string, 10)
	publish     = make(chan models.Event, 10)
	waitingList = list.New()
	subscribes  = list.New()
)

// This function handles all incoming chan messages.

func chatroom() {
	for {
		select {
		case sub := <-subscribe:

			if !isUserExist(subscribes, sub.Name) {
				subscribes.PushBack(sub)
				publish <- newEvent(models.EVENT_JOIN, sub.Name, "")
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			} else {
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn)
			}

		case event := <-publish:

			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}
			broadcastWebSocket(event)
			models.NewArchive(event)
			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}

		case unsub := <-unsubscribe:

			for sub := subscribes.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribes.Remove(sub)

					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub, "")
					break
				}
			}
		}
	}
}

func init() {
	go chatroom()
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
