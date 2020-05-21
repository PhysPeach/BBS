package controllers

import (
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
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *AccountsController) Get(){
	c.Layout = "layouts/application.tpl"
	c.TplName = "accounts/signup.tpl"
}

// Post ...
// @Title Create
// @Description create Accounts
// @Param	Name {string} string true
// @router / [post]
func (c *AccountsController) Post() {
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
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/index.tpl"
}


// GetOne ...
// @Title GetOne
// @Description get Accounts by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Accounts
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AccountsController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Accounts
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Accounts
// @Failure 403
// @router / [get]
func (c *AccountsController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Accounts
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Accounts	true		"body for Accounts content"
// @Success 200 {object} models.Accounts
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AccountsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Accounts
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AccountsController) Delete() {

}
