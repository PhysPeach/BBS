package main

import (
	_ "github.com/physpeach/bbs/routers"
	_ "github.com/physpeach/bbs/models"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}