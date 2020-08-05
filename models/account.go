package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego/orm"
)

type Account struct {
	ID   int64 `orm:"column(id);auto"` 
	Name string `orm:"column(name);size(32);index" form:"Name" valid:"Required;MinSize(1);MaxSize(32)"`
	Password string `orm:"column(password);"`
	CreatedAt time.Time `orm:"column(created_at);type(datetime);auto_now_add"`
	Threads []*Thread `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Account))
}

// AddAccount insert a new Account into database and returns
// last inserted Id on success.
func AddAccount(m *Account) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAccountById retrieves Account by Id. Returns error if
// Id doesn't exist
func GetAccountById(id int64) (v *Account, err error) {
	o := orm.NewOrm()
	v = &Account{ID: id}
	if err = o.QueryTable(new(Account)).Filter("Id", id).RelatedSel().One(v); err != nil {
		return nil, err
	}
	o.LoadRelated(v, "Threads", 1)
	return v, nil
}

func GetAccountByName(name string) (v *Account, err error) {
	o := orm.NewOrm()
	v = &Account{Name: name}
	if err = o.QueryTable(new(Account)).Filter("Name", name).RelatedSel().One(v); err != nil {
		return nil, err
	}
	o.LoadRelated(v, "Threads", 1)
	return v, nil
}

func GetAllAccount() (accounts []Account, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(new(Account)).OrderBy("-created_at").All(&accounts)
	return
}

// UpdateAccount updates Account by Id and returns error if
// the record to be updated doesn't exist
func UpdateAccountById(m *Account) (err error) {
	o := orm.NewOrm()
	v := Account{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAccount deletes Account by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAccount(id int64) (err error) {
	o := orm.NewOrm()
	v := Account{ID: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Account{ID: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func ExistSameAccountName(name string) (exist bool) {
	o := orm.NewOrm()
	exist = o.QueryTable(new(Account)).Filter("Name", name).Exist()
	return
}