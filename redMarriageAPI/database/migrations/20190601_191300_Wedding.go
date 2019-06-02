package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Wedding_20190601_191300 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Wedding_20190601_191300{}
	m.Created = "20190601_191300"

	migration.Register("Wedding_20190601_191300", m)
}

// Run the migrations
func (m *Wedding_20190601_191300) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Wedding(`id` int(11) NOT NULL AUTO_INCREMENT,`description` varchar(128) NOT NULL,`style` varchar(128) NOT NULL,`is_active` tinyint(1) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Wedding_20190601_191300) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Wedding`")
}
