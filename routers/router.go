package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/create", &controllers.UserController{}, "get,post:CreateUser")
	beego.Router("/user/index", &controllers.UserController{}, "get,post:Index")
	beego.Router("/posts/create", &controllers.PostsController{}, "get,post:Create")
	beego.Router("/posts/index", &controllers.PostsController{}, "get:Index")
	beego.Router("/posts/detail/:id", &controllers.PostsController{}, "get:Detail")
}
