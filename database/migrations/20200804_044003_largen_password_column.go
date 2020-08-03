package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type LargenPasswordColumn_20200804_044003 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &LargenPasswordColumn_20200804_044003{}
	m.Created = "20200804_044003"

	migration.Register("LargenPasswordColumn_20200804_044003", m)
}

// Run the migrations
func (m *LargenPasswordColumn_20200804_044003) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE account ALTER COLUMN password TYPE varchar(128)")
}

// Reverse the migrations
func (m *LargenPasswordColumn_20200804_044003) Down() {
	m.SQL("ALTER TABLE account ALTER COLUMN password TYPE varchar(64)")
}
