package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Relationships_20190601_213755 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Relationships_20190601_213755{}
	m.Created = "20190601_213755"

	migration.Register("Relationships_20190601_213755", m)
}

// Run the migrations
func (m *Relationships_20190601_213755) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;")
	m.SQL("SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;")
	m.SQL("SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';")
	m.SQL("ALTER TABLE `Account` ADD COLUMN `wedding_process` INT(11) NOT NULL AFTER `is_active`, DROP PRIMARY KEY, ADD PRIMARY KEY (`id`, `wedding_process`), ADD INDEX `fk_Account_WeddingProcess1_idx` (`wedding_process` ASC) VISIBLE;")
	m.SQL("ALTER TABLE `Step` ADD COLUMN `wedding` INT(11) NOT NULL AFTER `priority_order`, DROP PRIMARY KEY, ADD PRIMARY KEY (`id`, `wedding`), ADD INDEX `fk_Step_Wedding1_idx` (`wedding` ASC) VISIBLE;")
	m.SQL("ALTER TABLE `Timelines` ADD COLUMN `wedding_process` INT(11) NOT NULL AFTER `is_complete`, ADD COLUMN `step` INT(11) NOT NULL AFTER `wedding_process`, DROP PRIMARY KEY, ADD PRIMARY KEY (`id`, `wedding_process`, `step`), ADD INDEX `fk_Timelines_WeddingProcess1_idx` (`wedding_process` ASC) VISIBLE, ADD INDEX `fk_Timelines_Step1_idx` (`step` ASC) VISIBLE;")
	m.SQL("ALTER TABLE `User` ADD COLUMN `account` INT(11) NOT NULL AFTER `password`, DROP PRIMARY KEY, ADD PRIMARY KEY (`id`, `account`), ADD INDEX `fk_User_Account1_idx` (`account` ASC) VISIBLE;")
	m.SQL("ALTER TABLE `WeddingProcess` ADD COLUMN `wedding` INT(11) NOT NULL AFTER `is_active`, DROP PRIMARY KEY, ADD PRIMARY KEY (`id`, `wedding`), ADD INDEX `fk_WeddingProcess_Wedding_idx` (`wedding` ASC) VISIBLE;")
	m.SQL("ALTER TABLE `Account` ADD CONSTRAINT `fk_Account_WeddingProcess1` FOREIGN KEY (`wedding_process`) REFERENCES `WeddingProcess` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE `Step` ADD CONSTRAINT `fk_Step_Wedding1` FOREIGN KEY (`wedding`) REFERENCES `Wedding` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE `Timelines` ADD CONSTRAINT `fk_Timelines_WeddingProcess1` FOREIGN KEY (`wedding_process`) REFERENCES `WeddingProcess` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION, ADD CONSTRAINT `fk_Timelines_Step1` FOREIGN KEY (`step`) REFERENCES `Step` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE `User` ADD CONSTRAINT `fk_User_Account1` FOREIGN KEY (`account`) REFERENCES `Account` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE `WeddingProcess` ADD CONSTRAINT `fk_WeddingProcess_Wedding` FOREIGN KEY (`wedding`) REFERENCES `Wedding` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("SET SQL_MODE=@OLD_SQL_MODE;")
	m.SQL("SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;")
	m.SQL("SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;")
}

// Reverse the migrations
func (m *Relationships_20190601_213755) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
