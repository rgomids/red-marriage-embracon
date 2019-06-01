package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Account_20190601_190959 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Account_20190601_190959{}
	m.Created = "20190601_190959"

	migration.Register("Account_20190601_190959", m)
}

// Run the migrations
func (m *Account_20190601_190959) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Account(`id` int(11) NOT NULL AUTO_INCREMENT,`is_active` tinyint(1) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Account_20190601_190959) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Account`")
}
