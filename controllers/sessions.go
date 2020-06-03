package controllers

import (
	"fmt"
	"github.com/astaxie/beego"

	"github.com/physpeach/bbs/models"
)

// SessionsController operations for Sessions
type SessionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *SessionsController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
}

func (c *SessionsController) New() {
	c.Layout = "layouts/application.tpl"
	c.TplName = "sessions/new.tpl"
}

// Post ...
// @Title Create
// @Description create Sessions
// @Param	body		body 	models.Sessions	true		"body for Sessions content"
// @Success 201 {object} models.Sessions
// @Failure 403 body is empty
// @router / [post]
func (c *SessionsController) Create() {
	isValidate, account := ValificateAccount(c)
	if !isValidate {
		fmt.Println("Account does not exist")
		c.Abort("400")
	}
	c.SetSession("sessName", account.Name)
	fmt.Println("Success to create session")
	c.Ctx.Redirect(302, "/" + account.Name)
}

func ValificateAccount(c *SessionsController)(bool, *models.Account){
	acname := c.GetString("Name")
	account, err := models.CheckAccount(acname)
	if err != nil{
		fmt.Println(err)
	}
	return (account.ID != 0), account
}