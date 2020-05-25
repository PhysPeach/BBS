package controllers

import (
	"time"
	"github.com/physpeach/bbs/models"
	"github.com/astaxie/beego"
)

// AccountsController operations for Accounts
type AccountsController struct {
	beego.Controller
}

// URLMapping ...
func (c *AccountsController) URLMapping() {
	c.Mapping("Create", c.Create)
	c.Mapping("New", c.New)
}

func (c *AccountsController) New(){
	c.Layout = "layouts/application.tpl"
	c.TplName = "accounts/new.tpl"
}

// Post ...
// @Title Create
// @Description create Accounts
// @Param	Name {string} string true
// @router / [post]
func (c *AccountsController) Create() {
	account := models.Account{
		Name: c.GetString("Name"),
		CreatedAt: time.Now()}
	//avoid same name resistration
	if models.ExistSameName(&account) {
		c.Abort("400")
	}else{
		if _, err := models.AddAccount(&account); err != nil {
			c.Abort("500")
		}
	}
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/index.tpl"
}