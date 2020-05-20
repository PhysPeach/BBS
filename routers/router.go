package routers

import (
	"github.com/physpeach/bbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.TopController{})
}
