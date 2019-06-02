package routers

import (
	"redMarriageAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
		beego.NSNamespace("/step",
			beego.NSInclude(
				&controllers.StepController{},
			),
		),
		beego.NSNamespace("/timelines",
			beego.NSInclude(
				&controllers.TimelinesController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/wedding",
			beego.NSInclude(
				&controllers.WeddingController{},
			),
		),
		beego.NSNamespace("/weddingprocess",
			beego.NSInclude(
				&controllers.WeddingProcessController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
