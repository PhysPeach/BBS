package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCommentTable_20200607_062752 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCommentTable_20200607_062752{}
	m.Created = "20200607_062752"

	migration.Register("CreateCommentTable_20200607_062752", m)
}

// Run the migrations
func (m *CreateCommentTable_20200607_062752) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE comment (id serial PRIMARY KEY, content varchar(1024) NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, host_account_id int NOT NULL REFERENCES account (id), host_thread_id int NOT NULL REFERENCES thread (id))")
}

// Reverse the migrations
func (m *CreateCommentTable_20200607_062752) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE comment")
}
