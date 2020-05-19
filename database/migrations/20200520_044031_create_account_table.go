package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateAccountTable_20200520_044031 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAccountTable_20200520_044031{}
	m.Created = "20200520_044031"

	migration.Register("CreateAccountTable_20200520_044031", m)
}

// Run the migrations
func (m *CreateAccountTable_20200520_044031) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE account (id serial PRIMARY KEY, name varchar(32) NOT NULL, created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP)")
}

// Reverse the migrations
func (m *CreateAccountTable_20200520_044031) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE account")

}
