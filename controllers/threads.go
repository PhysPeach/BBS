package controllers

import (
	"strconv"
	"unicode/utf8"
	"html/template"
	"github.com/physpeach/bbs/models"
	"github.com/astaxie/beego"
)

// ThreadsController operations for Homepage
type ThreadsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ThreadsController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Create", c.Create)
	c.Mapping("Show", c.Show)
	c.Mapping("Destroy", c.Destroy)
}

func (c *ThreadsController) Index() {
	var sessAccountName string
	if sessAccountID := c.GetSession("sessAccountID"); sessAccountID != nil{
		sessAccount, _ := models.GetAccountById(sessAccountID.(int64))
		sessAccountName = sessAccount.Name
	}
	num, threads, err := models.GetAllThread()
	if err != nil{
		c.Abort("500")
	}
	threadExist := true
	if num == 0 {
		threadExist = false
	}
	c.Data["sessAccountName"] = sessAccountName
	c.Data["threadExist"] = threadExist
	c.Data["threads"] = threads
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/index.tpl"
}

func (c *ThreadsController) Create() {
	hostAccount, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil{
		c.Abort("400")
	}
	if sessAccountID := c.GetSession("sessAccountID"); sessAccountID != hostAccount.ID {
		c.Abort("401")
	}
	thread := models.Thread{
		Title: c.GetString("Title"),
		Description: c.GetString("Description"),
		HostAccount: hostAccount,
	}
	if thread.Title == "" || 64 < utf8.RuneCountInString(thread.Title) {
		c.Abort("400")
	}
	if thread.Description == "" || 256 < utf8.RuneCountInString(thread.Description){
		c.Abort("400")
	}
	threadid, err := models.AddThread(&thread)
	if err != nil {
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/" + hostAccount.Name + "/" + strconv.FormatInt(threadid, 10))
}

func (c *ThreadsController) Show() {
	var sessAccountName string
	if sessAccountID := c.GetSession("sessAccountID"); sessAccountID != nil{
		sessAccount, _ := models.GetAccountById(sessAccountID.(int64))
		sessAccountName = sessAccount.Name
	}
	threadidStr := c.Ctx.Input.Param(":threadid")
	if len(threadidStr) >= 2 && threadidStr[0:1] == "0" {
		c.Abort("400")
	}
	threadid, err := strconv.ParseInt(threadidStr, 10, 64)
	if err != nil {
		c.Abort("500")
	}
	thread, err := models.GetThreadById(threadid)
	if err != nil {
		c.Abort("404")
	}
	if thread.HostAccount.Name != c.Ctx.Input.Param(":accountname"){
		c.Abort("404")
	}
	c.SetSession("sessLastVisitedThreadID", threadid)
	c.Data["thread"] = thread
	c.Data["editable"] = (sessAccountName == thread.HostAccount.Name)
	c.Data["sessAccountName"] = sessAccountName
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/show.tpl"
}

func (c *ThreadsController) Destroy() {
	threadid, err := strconv.ParseInt(c.Ctx.Input.Param(":threadid"), 10, 64)
	if err != nil {
		c.Abort("500")
	}
	thread, err := models.GetThreadById(threadid)
	if err != nil {
		c.Abort("500")
	}
	if sessAccountID := c.GetSession("sessAccountID"); sessAccountID != thread.HostAccount.ID {
		c.Abort("401")
	}
	if err = models.DeleteThread(threadid); err != nil {
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/")
}