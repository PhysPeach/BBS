package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	ID   int64 `orm:"column(id);auto"` 
	Content string `orm:"column(content);size(1024);"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	HostAccount *Account `orm:"column(host_account_id);rel(fk)"`
	HostThread *Thread `orm:"column(host_thread_id);rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

// AddComment insert a new Comment into database and returns
// last inserted Id on success.
func AddComment(m *Comment) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCommentById retrieves Comment by Id. Returns error if
// Id doesn't exist
func GetCommentById(id int64) (v *Comment, err error) {
	o := orm.NewOrm()
	v = &Comment{ID: id}
	if err = o.QueryTable(new(Comment)).Filter("ID", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllCommentByHostThreadId(id int64)(comments []Comment, err error){
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Comment)).Filter("host_thread_id", id).RelatedSel().OrderBy("created_at").All(&comments)
	return
}

func GetAllComment() (comments []Comment, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Comment)).RelatedSel().OrderBy("-created_at").All(&comments)
	return
}

// UpdateComment updates Comment by Id and returns error if
// the record to be updated doesn't exist
func UpdateCommentById(m *Comment) (err error) {
	o := orm.NewOrm()
	v := Comment{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteComment deletes Comment by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComment(id int64) (err error) {
	o := orm.NewOrm()
	v := Comment{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Comment{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}