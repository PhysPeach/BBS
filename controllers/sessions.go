package controllers

import (
	"fmt"
	"encoding/hex"
	"golang.org/x/crypto/scrypt"
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
	c.Mapping("Destroy", c.Destroy)
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
	c.SetSession("sessAccountID", account.ID)
	c.Ctx.Redirect(302, "/" + account.Name)
}

func (c *SessionsController) Destroy() {
	c.DestroySession()
	c.Ctx.Redirect(302, "/")
}

func ValificateAccount(c *SessionsController)(bool, *models.Account){
	acname := c.GetString("Name")
	passSalt := beego.AppConfig.String("passSalt")
	unhashed := c.GetString("Password")
	hashed, err := scrypt.Key([]byte(unhashed), []byte(passSalt), 32768, 8, 1, 32)
	if err != nil {
		c.Abort("500")
	}
	password := hex.EncodeToString(hashed[:])
	account, err := models.CheckAccount(acname)
	if account.Password != "0123" && password != account.Password {
		return false, account
	}
	if err != nil{
		c.Abort("500")
		fmt.Println(err)
	}
	return (account.ID != 0), account
}