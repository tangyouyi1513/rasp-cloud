package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "Config",
			Router: `/config`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "GetRasps",
			Router: `/rasp`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"],
		beego.ControllerComments{
			Method: "GetLatest",
			Router: `/latest`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:RaspController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:RaspController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/find`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"],
		beego.ControllerComments{
			Method: "NewToken",
			Router: `/new`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:UserController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:UserController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
