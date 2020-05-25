package routers

import (
	"github.com/physpeach/bbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ThreadsController{})
	signupNs := beego.NewNamespace("/signup",
		beego.NSRouter("", &controllers.AccountsController{},"post:Create"),
		beego.NSRouter("/new", &controllers.AccountsController{},"get:New"),
	)
	beego.AddNamespace(signupNs)
	loginNs := beego.NewNamespace("/login",
		beego.NSRouter("", &controllers.SessionsController{}, "post:Create"),
		beego.NSRouter("/new", &controllers.SessionsController{},"get:New"),
	)
	beego.AddNamespace(loginNs)
}
