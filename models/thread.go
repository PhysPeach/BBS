package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Thread struct {
	ID   int64 `orm:"column(id);auto"` 
	Title string `orm:"column(title);size(64);" form:"Title" valid:"Required;MinSize(1);MaxSize(64)"`
	Description string `orm:"column(description);size(256);" form:"Description" valid:"Required;MinSize(1);MaxSize(256)"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	HostAccount *Account `orm:"column(host_account_id);rel(fk)"`
	Comments []*Comment `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Thread))
}

// AddThread insert a new Thread into database and returns
// last inserted Id on success.
func AddThread(m *Thread) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetThreadById retrieves Thread by Id. Returns error if
// Id doesn't exist
func GetThreadById(id int64) (v *Thread, err error) {
	o := orm.NewOrm()
	v = &Thread{ID: id}
	if err = o.QueryTable(new(Thread)).Filter("ID", id).RelatedSel().One(v); err != nil {
		return nil, err
	}
	o.LoadRelated(v, "Comments", 1)
	return v, nil
}

func GetAllThreadByHostAccountId(id int64)(threads []Thread, err error){
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Thread)).Filter("host_account_id", id).RelatedSel().OrderBy("-created_at").All(&threads)
	return
}

func GetAllThread() (threads []Thread, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Thread)).RelatedSel().OrderBy("-created_at").All(&threads)
	return
}

// UpdateThread updates Thread by Id and returns error if
// the record to be updated doesn't exist
func UpdateThreadById(m *Thread) (err error) {
	o := orm.NewOrm()
	v := Thread{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteThread deletes Thread by Id and returns error if
// the record to be deleted doesn't exist
func DeleteThread(id int64) (err error) {
	o := orm.NewOrm()
	v := Thread{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Thread{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}