package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}


func (c *ErrorController) Error400() {
	c.Ctx.ResponseWriter.WriteHeader(400)
	c.Data["Message"] = "400 Bad Request"
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error401() {
	c.Ctx.ResponseWriter.WriteHeader(401)
	c.Data["Message"] = "401 Unauthorized"
	c.TplName = "error.tpl"
}

func (c *ErrorController) Error404() {
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.Data["Message"] = "404 Not Found"
	c.TplName = "error.tpl"
}

// Error500 はサーバーエラーを表示する
func (c *ErrorController) Error500() {
	c.Ctx.ResponseWriter.WriteHeader(500)
	c.Data["Message"] = "500 Internal Server Error"
	c.TplName = "error.tpl"
}