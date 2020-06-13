package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddActionOfForeignKeyOnDeleteOrUpdate_20200613_163338 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddActionOfForeignKeyOnDeleteOrUpdate_20200613_163338{}
	m.Created = "20200613_163338"

	migration.Register("AddActionOfForeignKeyOnDeleteOrUpdate_20200613_163338", m)
}

// Run the migrations
func (m *AddActionOfForeignKeyOnDeleteOrUpdate_20200613_163338) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE thread ADD FOREIGN KEY (host_account_id) REFERENCES account (id) ON DELETE CASCADE ON UPDATE CASCADE")
	m.SQL("ALTER TABLE comment ADD FOREIGN KEY (host_account_id) REFERENCES account (id) ON DELETE CASCADE ON UPDATE CASCADE")
	m.SQL("ALTER TABLE comment ADD FOREIGN KEY (host_thread_id) REFERENCES thread (id) ON DELETE CASCADE ON UPDATE CASCADE")
}

// Reverse the migrations
func (m *AddActionOfForeignKeyOnDeleteOrUpdate_20200613_163338) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE comment DROP FOREIGN KEY (host_thread_id)")
	m.SQL("ALTER TABLE comment ADD FOREIGN KEY (host_thread_id) REFERENCES thread (id)")
	m.SQL("ALTER TABLE comment DROP FOREIGN KEY (host_account_id)")
	m.SQL("ALTER TABLE comment ADD FOREIGN KEY (host_account_id) REFERENCES account (id)")
	m.SQL("ALTER TABLE thread DROP FOREIGN KEY (host_account_id)")
	m.SQL("ALTER TABLE thread ADD FOREIGN KEY (host_account_id) REFERENCES account (id)")
}
