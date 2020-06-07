package routers

import (
	"strings"
	"github.com/physpeach/bbs/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	var FilterMethod = func(ctx *context.Context) {
		if ctx.Input.Query("_method")!="" && ctx.Input.IsPost(){
			  ctx.Request.Method = strings.ToUpper(ctx.Input.Query("_method"))
		}
	}
	beego.InsertFilter("*", beego.BeforeRouter, FilterMethod)

	beego.Router("/", &controllers.ThreadsController{}, "get:Index")
	accountsNs := beego.NewNamespace("/:accountname",
		beego.NSRouter("/", &controllers.AccountsController{}, "get:Show"),
		beego.NSRouter("/", &controllers.AccountsController{}, "put:Update"),
		beego.NSRouter("/", &controllers.AccountsController{}, "delete:Destroy"),
		beego.NSRouter("/edit", &controllers.AccountsController{}, "get:Edit"),
		
		beego.NSRouter("/", &controllers.ThreadsController{}, "post:Create"),
		beego.NSNamespace("/:threadid",
			beego.NSRouter("/", &controllers.ThreadsController{}, "get:Show"),
			beego.NSRouter("/", &controllers.ThreadsController{}, "delete:Destroy"),

			beego.NSRouter("/", &controllers.CommentsController{}, "post:Create"),
			beego.NSNamespace("/:commentid",
				beego.NSRouter("/", &controllers.CommentsController{}, "get:Show"),
				beego.NSRouter("/", &controllers.CommentsController{}, "delete:Destroy"),
			),
		),
	)
	beego.AddNamespace(accountsNs)
	signupNs := beego.NewNamespace("/signup",
		beego.NSRouter("/", &controllers.AccountsController{},"post:Create"),
		beego.NSRouter("/new", &controllers.AccountsController{},"get:New"),
	)
	beego.AddNamespace(signupNs)
	loginNs := beego.NewNamespace("/login",
		beego.NSRouter("/", &controllers.SessionsController{}, "post:Create"),
		beego.NSRouter("/", &controllers.SessionsController{}, "delete:Destroy"),
		beego.NSRouter("/new", &controllers.SessionsController{},"get:New"),
	)
	beego.AddNamespace(loginNs)
}
