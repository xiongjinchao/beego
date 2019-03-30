package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["hello/controllers:ArticleController"] = append(beego.GlobalControllerRouter["hello/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/article/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:ArticleController"] = append(beego.GlobalControllerRouter["hello/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/article/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:ArticleController"] = append(beego.GlobalControllerRouter["hello/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/article/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:ArticleController"] = append(beego.GlobalControllerRouter["hello/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/article/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["hello/controllers:ArticleController"] = append(beego.GlobalControllerRouter["hello/controllers:ArticleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/article/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
