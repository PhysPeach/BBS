package routers

import (
	"github.com/physpeach/bbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ThreadsController{})
	beego.Router("/signup", &controllers.AccountsController{})
	beego.Router("/login", &controllers.SessionsController{})
}
