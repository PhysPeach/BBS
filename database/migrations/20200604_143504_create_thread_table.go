package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateThreadTable_20200604_143504 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateThreadTable_20200604_143504{}
	m.Created = "20200604_143504"

	migration.Register("CreateThreadTable_20200604_143504", m)
}

// Run the migrations
func (m *CreateThreadTable_20200604_143504) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE thread (id serial PRIMARY KEY, title varchar(64) NOT NULL, description varchar(256) NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, host_account_id int NOT NULL REFERENCES account (id))")
}

// Reverse the migrations
func (m *CreateThreadTable_20200604_143504) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE thread")
}
