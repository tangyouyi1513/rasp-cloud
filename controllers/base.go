package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
)

// base controller
type BaseController struct {
	beego.Controller
}

func (o *BaseController) Prepare() {
	// auth
	authId := o.Ctx.Input.Header("X-OpenRASP-Authentication-ID")
	if authId != beego.AppConfig.String("AuthID") {
		o.ServeError(401)
	}
}

func (o *BaseController) Serve(data interface{}) {
	o.Data["json"] = map[string]interface{}{"status": 0, "description": "ok", "data": data}
	o.ServeJSON()
}

func (o *BaseController) ServeError(code int, description ...string) {
	var des string
	if len(description) == 0 {
		des = http.StatusText(code)
	} else {
		des = description[0]
	}
	o.Data["json"] = map[string]interface{}{"status": code, "description": des}
	o.ServeJSON()
}
