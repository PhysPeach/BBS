package controllers

import (
	"regexp"
	"encoding/hex"
	"html/template"
	"golang.org/x/crypto/bcrypt"
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
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.tpl"
	c.TplName = "accounts/new.tpl"
}

// Post ...
// @Title Create
// @Description create Accounts
// @Param      Name {string} string true
// @router / [post]
func (c *AccountsController) Create() {
	unhashed := c.GetString("Password")
	if unhashed != c.GetString("PasswordConfirmation"){
		c.Abort("400")
	}
	passRegex := regexp.MustCompile(`^[a-zA-Z\d]{8,32}$`)
	if !passRegex.MatchString(unhashed) {
		c.Abort("400")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(unhashed), 12)
	if err != nil {
		c.Abort("500")
	}
	password := hex.EncodeToString(hashed[:])
	account := models.Account{
		Name: c.GetString("Name"),
		Password: password}
	isValid := ConfirmAccountName(account.Name)
	if !isValid {
		c.Abort("400")
	}
	if _, err := models.AddAccount(&account); err != nil {
		c.Abort("500")
	}
	c.SetSession("sessAccountID", account.ID)
	c.Ctx.Redirect(302, "/" + account.Name)
}

func (c *AccountsController) Show() {
	var sessAccountName string
	if sessAccountID := c.GetSession("sessAccountID"); sessAccountID != nil{
		sessAccount, _ := models.GetAccountById(sessAccountID.(int64))
		sessAccountName = sessAccount.Name
	}
	account, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil{
		c.Abort("404")
	}
	c.Data["sessAccountName"] = sessAccountName
	c.Data["editable"] = (sessAccountName == account.Name)
	c.Data["account"] = account
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.tpl"
	c.TplName = "accounts/show.tpl"
}

func (c *AccountsController) Edit() {
	account, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil{
		c.Abort("404")
	}
	sessAccountID := c.GetSession("sessAccountID")
	if sessAccountID == nil{
		c.Abort("401")
	}
	sessAccount, _ := models.GetAccountById(sessAccountID.(int64))
	if sessAccount.Name != account.Name {
		c.Abort("401")
	}
	c.Data["sessAccountName"] = sessAccount.Name
	c.Data["accountname"] = account.Name
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.tpl"
	c.TplName = "accounts/edit.tpl"
}

func(c *AccountsController) Update() {
	account, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil {
		c.Abort("404")
	}
	if sessAccountID := c.GetSession("sessAccountID"); sessAccountID != account.ID {
		c.Abort("401")
	}
	updatingAccount := models.Account{
		Name: c.GetString("Name")}
	isValid := ConfirmAccountName(updatingAccount.Name)
	if !isValid {
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
		c.Abort("500")
	}
	if sessAccountID := c.GetSession("sessAccountID"); sessAccountID != account.ID {
		c.Abort("401")
	}
	if err = models.DeleteAccount(account.ID); err != nil {
		c.Abort("500")
	}
	c.DestroySession()
	c.Ctx.Redirect(302, "/")
}

func ConfirmAccountName(name string) (bool) {
	nameRegex := regexp.MustCompile(`^[a-z\d]{1,32}$`)
	if !nameRegex.MatchString(name) {
		return false
	}
	if exist := models.ExistSameAccountName(name); exist {
		return false
	}
	//blacklist
	if name == "login" || name == "signup"{
		return false
	}
	return true
}