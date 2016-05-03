package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.AppController{})
	beego.Router("/login", &controllers.LoginController{})
	//用户管理
	beego.Router("/create", &controllers.UserController{}, "get,post:CreateUser")
	beego.Router("/user/index", &controllers.UserController{}, "get,post:Index")
	beego.Router("/user/edit/:id", &controllers.UserController{}, "get:Edit")
	//文章管理
	beego.Router("/posts/create", &controllers.PostsController{}, "get,post:Create")
	beego.Router("/posts/index", &controllers.PostsController{}, "get:Index")
	beego.Router("/posts/detail/:id", &controllers.PostsController{}, "get:Detail")
	beego.Router("/posts/edit/:id", &controllers.PostsController{}, "get:Edit")
	beego.Router("/posts/save", &controllers.PostsController{}, "post:SaveEdit")
	beego.Router("/posts/delete", &controllers.PostsController{}, "post:Delete")
	//聊天室
	beego.Router("/join", &controllers.AppController{}, "get,post:Join")
	beego.Router("/ws", &controllers.WebSocketController{}, "get:Get")
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get,post:Join")
}
