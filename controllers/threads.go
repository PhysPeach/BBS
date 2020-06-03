package controllers

import (
	_ "github.com/physpeach/bbs/models"
	"github.com/astaxie/beego"
)

// ThreadsController operations for Homepage
type ThreadsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ThreadsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *ThreadsController) Get() {
	c.Data["sessName"] = c.GetSession("sessName")
	c.Layout = "layouts/application.tpl"
	c.TplName = "threads/index.tpl"
}

// Post ...
// @Title Create
// @Description create Homepage
// @Param	body		body 	models.Homepage	true		"body for Homepage content"
// @Success 201 {object} models.Homepage
// @Failure 403 body is empty
// @router / [post]
func (c *ThreadsController) Post() {
}

// GetOne ...
// @Title GetOne
// @Description get Homepage by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Homepage
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ThreadsController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Homepage
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Homepage
// @Failure 403
// @router / [get]
func (c *ThreadsController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Homepage
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Homepage	true		"body for Homepage content"
// @Success 200 {object} models.Homepage
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ThreadsController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Homepage
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ThreadsController) Delete() {

}
