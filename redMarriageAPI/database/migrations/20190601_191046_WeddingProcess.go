package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type WeddingProcess_20190601_191046 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &WeddingProcess_20190601_191046{}
	m.Created = "20190601_191046"

	migration.Register("WeddingProcess_20190601_191046", m)
}

// Run the migrations
func (m *WeddingProcess_20190601_191046) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE WeddingProcess(`id` int(11) NOT NULL AUTO_INCREMENT,`is_active` tinyint(1) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *WeddingProcess_20190601_191046) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `WeddingProcess`")
}
