package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Step_20190601_190919 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Step_20190601_190919{}
	m.Created = "20190601_190919"

	migration.Register("Step_20190601_190919", m)
}

// Run the migrations
func (m *Step_20190601_190919) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Step(`id` int(11) NOT NULL AUTO_INCREMENT,`is_active` tinyint(1) NOT NULL,`description` varchar(128) NOT NULL,`priority_order` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Step_20190601_190919) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Step`")
}
