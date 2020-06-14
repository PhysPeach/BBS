package controllers

import (
	"fmt"
	"regexp"
	"time"
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
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
	passSalt := beego.AppConfig.String("passSalt")
	unhashed := c.GetString("Password")
	passRegex := regexp.MustCompile(`^[a-zA-Z\d]{8,32}$`)
	if !passRegex.MatchString(unhashed) {
		c.Abort("400")
	}
	hashed, err := scrypt.Key([]byte(unhashed), []byte(passSalt), 32768, 8, 1, 32)
	if err != nil {
		c.Abort("500")
	}
	password := hex.EncodeToString(hashed[:])
	account := models.Account{
		Name: c.GetString("Name"),
		Password: password,
		CreatedAt: time.Now()}
	nameRegex := regexp.MustCompile(`^[a-z\d]{8,32}$`)
	if !nameRegex.MatchString(account.Name) {
		c.Abort("400")
	}
	//avoid same name resistration
	if models.ExistSameName(&account) {
		c.Abort("400")
	}
	if _, err := models.AddAccount(&account); err != nil {
		fmt.Println(err)
		c.Abort("500")
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

	threads, err := models.GetAllThreadByHostAccountId(account.ID)
	if err != nil {
		c.Abort("500")
	}
	c.Data["account"] = account
	c.Data["threads"] = threads
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
	nameRegex := regexp.MustCompile(`^[a-z\d]{8,32}$`)
	if !nameRegex.MatchString(account.Name) {
		c.Abort("400")
	}
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
	c.DestroySession()
	c.Ctx.Redirect(302, "/")
}