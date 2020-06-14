package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddPasswordOnAccount_20200614_160911 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddPasswordOnAccount_20200614_160911{}
	m.Created = "20200614_160911"

	migration.Register("AddPasswordOnAccount_20200614_160911", m)
}

// Run the migrations
func (m *AddPasswordOnAccount_20200614_160911) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE account ADD COLUMN password varchar(64) default '0123'")
}

// Reverse the migrations
func (m *AddPasswordOnAccount_20200614_160911) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE account DROP COLUMN password")
}
