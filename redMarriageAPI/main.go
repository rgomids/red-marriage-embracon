package main

import (
	_ "redMarriageAPI/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbURL := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlurls") + ":3306)/" + beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", dbURL)
	orm.DefaultTimeLoc, _ = time.LoadLocation("America/Sao_Paulo")
	orm.Debug = true
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
