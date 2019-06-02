package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20190601_185532 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20190601_185532{}
	m.Created = "20190601_185532"

	migration.Register("User_20190601_185532", m)
}

// Run the migrations
func (m *User_20190601_185532) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE User(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`surname` varchar(128) NOT NULL,`cpf` varchar(128) NOT NULL,`rg` varchar(128) NOT NULL,`cellphone` varchar(128) NOT NULL,`phone` varchar(128) NOT NULL,`address` varchar(128) NOT NULL,`is_active` tinyint(1) NOT NULL,`email` varchar(128) NOT NULL,`password` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *User_20190601_185532) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `User`")
}
