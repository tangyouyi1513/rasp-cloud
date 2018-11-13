package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "ConfigAlarm",
			Router: `/alarm/config`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "UpdateAppAlgorithmConfig",
			Router: `/algorithm/config`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("id"),
				param.New("algorithmConfig"),
			),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "TestDing",
			Router: `/ding/test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("config"),
			),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "TestEmail",
			Router: `/email/test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "UpdateAppGeneralConfig",
			Router: `/general/config`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "GetApp",
			Router: `/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "TestHttp",
			Router: `/http/test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("config"),
			),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "GetPlugins",
			Router: `/plugin/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "SetSelectedPlugin",
			Router: `/plugin/select`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "GetSelectedPlugin",
			Router: `/plugin/select/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "GetRasps",
			Router: `/rasp/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:AppController"],
		beego.ControllerComments{
			Method: "UpdateAppWhiteListConfig",
			Router: `/whitelist/config`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(
				param.New("id"),
				param.New("whiteListConfig"),
			),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:PluginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:RaspController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:RaspController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:RaspController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:RaspController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/search`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:ReportController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:ReportController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/dashboard`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers/api:TokenController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"post"},
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
