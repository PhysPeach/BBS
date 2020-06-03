package controllers

import (
	"fmt"
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
	c.Mapping("Show", c.Show)
	c.Mapping("Edit", c.Edit)
	c.Mapping("Update", c.Update)
	c.Mapping("Destroy", c.Destroy)
}

func (c *AccountsController) New(){
	c.Layout = "layouts/application.tpl"
	c.TplName = "accounts/new.tpl"
}

// Post ...
// @Title Create
// @Description create Accounts
// @Param      Name {string} string true
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
	c.SetSession("sessName", account.Name)
	c.Ctx.Redirect(302, "/" + account.Name)
}

func (c *AccountsController) Show() {
	sessName := c.GetSession("sessName")
	account, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil{
		fmt.Println("Nil Account")
		c.Abort("400")
	}
	c.Data["accountname"] = account.Name
	c.Data["editable"] = (sessName == account.Name)
	c.Data["sessName"] = sessName
	c.Layout = "layouts/application.tpl"
	c.TplName = "accounts/show.tpl"
}

func (c *AccountsController) Edit() {
	sessName := c.GetSession("sessName")
	account, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil{
		fmt.Println("Nil Account")
		c.Abort("400")
	}
	if sessName != account.Name {
		c.Abort("403")
	} else {
		c.Data["accountname"] = account.Name
		c.Data["sessName"] = sessName
		c.Layout = "layouts/application.tpl"
		c.TplName = "accounts/edit.tpl"
	}
}

func(c *AccountsController) Update() {
	account, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil {
		c.Abort("500")
	}

	updatingAccount := models.Account{
		Name: c.GetString("Name")}
	//avoid same name resistration
	if models.ExistSameName(&updatingAccount) {
		c.Abort("400")
	}
	account.Name = updatingAccount.Name
	if err := models.UpdateAccountById(account); err != nil {
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/" + account.Name)
}

func(c *AccountsController) Destroy() {
	account, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil {
		fmt.Println(err)
		c.Abort("500")
	}
	if err = models.DeleteAccount(account.ID); err != nil {
		fmt.Println(err)
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/")
}