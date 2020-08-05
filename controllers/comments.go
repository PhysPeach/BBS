package controllers

import (
	"strconv"
	"unicode/utf8"
	"github.com/physpeach/bbs/models"
	"github.com/astaxie/beego"
)

// CommentsController operations for Homepage
type CommentsController struct {
	beego.Controller
}

// URLMapping ...
func (c *CommentsController) URLMapping() {
	c.Mapping("Create", c.Create)
}

func (c *CommentsController) Create() {
	sessAccountID := c.GetSession("sessAccountID")
	if sessAccountID == nil{
		c.Abort("401")
	}
	hostAccount, err := models.GetAccountById(sessAccountID.(int64))
	if err != nil{
		c.Abort("400")
	}
	hostThreadid, err := strconv.ParseInt(c.Ctx.Input.Param(":threadid"), 10, 64)
	if err != nil {
		c.Abort("500")
	}
	if sessLastVisitedThread := c.GetSession("sessLastVisitedThread"); sessLastVisitedThread != hostThreadid {
		c.Abort("400")
	}
	hostThread, err := models.GetThreadById(hostThreadid)
	if err != nil{
		c.Abort("400")
	}
	comment := models.Comment{
		Content: c.GetString("Content"),
		HostAccount: hostAccount,
		HostThread: hostThread,
	}
	if comment.Content == "" || 1024 < utf8.RuneCountInString(comment.Content){
		c.Abort("400")
	}
	if _, err := models.AddComment(&comment); err != nil {
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/" + hostThread.HostAccount.Name + "/" + strconv.FormatInt(hostThread.ID, 10))
}