package controllers

import (
	"fmt"
	"time"
	"strconv"
	"unicode/utf8"
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
	c.Data["sessName"] = c.GetSession("sessName")
	threads, err := models.GetAllThread()
	if err != nil{
		c.Abort("500")
	}
	c.Data["threads"] = threads
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/index.tpl"
}

func (c *ThreadsController) Create() {
	hostAccount, err := models.GetAccountByName(c.Ctx.Input.Param(":accountname"))
	if err != nil{
		fmt.Println("Nil Account")
		c.Abort("400")
	}
	thread := models.Thread{
		Title: c.GetString("Title"),
		Description: c.GetString("Description"),
		CreatedAt: time.Now(),
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
	sessName := c.GetSession("sessName")
	threadid, err := strconv.ParseInt(c.Ctx.Input.Param(":threadid"), 10, 64)
	if err != nil {
		c.Abort("500")
	}
	thread, err := models.GetThreadById(threadid)
	if err != nil {
		c.Abort("500")
	}
	comments, err := models.GetAllCommentByHostThreadId(threadid)
	if err != nil {
		c.Abort("500")
	}
	c.Data["thread"] = thread
	c.Data["comments"] = comments
	c.Data["editable"] = (sessName == thread.HostAccount.Name)
	c.Data["sessName"] = sessName
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/show.tpl"
}

func (c *ThreadsController) Destroy() {
	threadid, err := strconv.ParseInt(c.Ctx.Input.Param(":threadid"), 10, 64)
	if err != nil {
		c.Abort("500")
	}
	if err = models.DeleteThread(threadid); err != nil {
		fmt.Println(err)
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/")
}