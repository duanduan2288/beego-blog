package routers

import (
	"blog/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	//用户管理
	beego.Router("/create", &controllers.UserController{}, "get,post:CreateUser")
	beego.Router("/user/index", &controllers.UserController{}, "get,post:Index")
	//文章管理
	beego.Router("/posts/create", &controllers.PostsController{}, "get,post:Create")
	beego.Router("/posts/index", &controllers.PostsController{}, "get:Index")
	beego.Router("/posts/detail/:id", &controllers.PostsController{}, "get:Detail")
	beego.Router("/posts/edit/:id", &controllers.PostsController{}, "get:Edit")
	beego.Router("/posts/save", &controllers.PostsController{}, "post:SaveEdit")
	beego.Router("/posts/delete", &controllers.PostsController{}, "post:Delete")
}
