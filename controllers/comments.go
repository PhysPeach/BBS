package controllers

import (
	"fmt"
	"time"
	"strconv"
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
	c.Mapping("Show", c.Show)
	c.Mapping("Destroy", c.Destroy)
}

func (c *CommentsController) Create() {
	hostAccount, err := models.GetAccountByName(":accountname")
	if err != nil{
		fmt.Println("Nil Account")
		c.Abort("400")
	}
	hostThreadid, err := strconv.ParseInt(c.Ctx.Input.Param(":hostThreadid"), 10, 64)
	if err != nil {
		c.Abort("500")
	}
	hostThread, err := models.GetThreadById(hostThreadid)
	if err != nil{
		fmt.Println("Nil Account")
		c.Abort("400")
	}
	comment := models.Comment{
		Content: c.GetString("Content"),
		CreatedAt: time.Now(),
		HostAccount: hostAccount,
		HostThread: hostThread,
	}
	if comment.Content == "" {
		c.Abort("400")
	}
	if _, err := models.AddComment(&comment); err != nil {
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/")
}

func (c *CommentsController) Show() {
}

func (c *CommentsController) Destroy() {
}