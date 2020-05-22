package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"

	"github.com/physpeach/bbs/models"
)

var globalSessions *session.Manager
func init(){
	sessionConfig := &session.ManagerConfig{
		CookieName: "gosessionid",
		EnableSetCookie: true,
		CookieLifeTime: 3600,
		Gclifetime: 3600,
		Maxlifetime: 3600,
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
    go globalSessions.GC()
}

// SessionsController operations for Sessions
type SessionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *SessionsController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *SessionsController) Get() {
	c.Layout = "layouts/application.tpl"
	c.TplName = "sessions/login.tpl"
}

// Post ...
// @Title Create
// @Description create Sessions
// @Param	body		body 	models.Sessions	true		"body for Sessions content"
// @Success 201 {object} models.Sessions
// @Failure 403 body is empty
// @router / [post]
func (c *SessionsController) Post() {
	isValidate, account := ValificateAccount(c)
	if !isValidate {
		fmt.Println("Account does not exist")
		c.Abort("400")
	}
	sess, err := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	if err != nil {
		fmt.Println("SessionStart Failed")
		c.Abort("400")
	}
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	sess.Set("acid", account.ID)
	fmt.Println("Success to create session")
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/index.tpl"

}

func ValificateAccount(c *SessionsController)(bool, *models.Account){
	acname := c.GetString("Name")
	account, err := models.CheckAccount(acname)
	if err != nil{
		fmt.Println(err)
	}
	return (account.ID != 0), account
}