package filter

import (
	"github.com/astaxie/beego"
	"rasp-cloud/models"
	"net/http"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/v1/agent/*", beego.BeforeRouter, authAgent)
	beego.InsertFilter("/v1/api/*", beego.BeforeRouter, authApi)
}

func authAgent(ctx *context.Context) {
	appId := ctx.Input.Header("X-OpenRASP-AppID")
	app, err := models.GetAppById(appId)
	if appId == "" || err != nil || app == nil {
		ctx.Output.JSON(map[string]interface{}{
			"status": http.StatusUnauthorized, "description": http.StatusText(http.StatusUnauthorized)},
			false, false)
	}
}

func authApi(ctx *context.Context) {
	cookie := ctx.GetCookie(models.AuthCookieName)
	if has, err := models.HasCookie(cookie); !has || err != nil {
		token := ctx.Input.Header("RASP-AUTH-ST-TOKEN")
		if has, err = models.HasTokent(token); !has || err != nil {
			ctx.Output.JSON(map[string]interface{}{
				"status": http.StatusUnauthorized, "description": http.StatusText(http.StatusUnauthorized)},
				false, false)
			panic("")
		}
	}
}
