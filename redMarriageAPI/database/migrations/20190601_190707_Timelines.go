package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Timelines_20190601_190707 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Timelines_20190601_190707{}
	m.Created = "20190601_190707"

	migration.Register("Timelines_20190601_190707", m)
}

// Run the migrations
func (m *Timelines_20190601_190707) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Timelines(`id` int(11) NOT NULL AUTO_INCREMENT,`is_complete` tinyint(1) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Timelines_20190601_190707) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Timelines`")
}
