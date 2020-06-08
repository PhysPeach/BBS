package controllers

import (
	"fmt"
	"time"
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
	sessName := c.GetSession("sessName")
	hostAccount, err := models.GetAccountByName(sessName.(string))
	if err != nil{
		fmt.Println("Nil Account")
		c.Abort("400")
	}
	hostThreadid, err := strconv.ParseInt(c.Ctx.Input.Param(":threadid"), 10, 64)
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
	fmt.Println(comment.Content)
	if comment.Content == "" || 1024 < utf8.RuneCountInString(comment.Content){
		c.Abort("400")
	}
	if _, err := models.AddComment(&comment); err != nil {
		c.Abort("500")
	}
	c.Ctx.Redirect(302, "/" + hostThread.HostAccount.Name + "/" + strconv.FormatInt(hostThread.ID, 10))
}