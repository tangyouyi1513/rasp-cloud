package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["rasp-cloud/controllers:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers:PluginController"],
		beego.ControllerComments{
			Method: "Upgrade",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["rasp-cloud/controllers:PluginController"] = append(beego.GlobalControllerRouter["rasp-cloud/controllers:PluginController"],
		beego.ControllerComments{
			Method: "Upload",
			Router: `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
