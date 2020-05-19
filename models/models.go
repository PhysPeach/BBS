package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	psqlConn := beego.AppConfig.String("psqlConn")
	if err := orm.RegisterDataBase("default", "postgres", psqlConn); err != nil {
		panic(err)
	}
}