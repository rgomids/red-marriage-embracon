package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:StepController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:TimelinesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"] = append(beego.GlobalControllerRouter["redMarriageAPI/controllers:WeddingProcessController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
