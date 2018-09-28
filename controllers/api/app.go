//Copyright 2017-2018 Baidu Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http: //www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package api

import (
	"net/http"
	"rasp-cloud/models"
	"rasp-cloud/controllers"
	"encoding/json"
	"time"
	"strings"
)

// Operations about app
type AppController struct {
	controllers.BaseController
}

// @router / [get]
func (o *AppController) GetAll() {
	name := o.GetString("name")
	if len(name) >= 512 {
		o.ServeError(http.StatusBadRequest, "the length of app name must be less than 512")
	}
	if name == "" {
		page, err := o.GetInt("page")
		if err != nil {
			o.ServeError(http.StatusBadRequest, "failed to get page param: "+err.Error())
		}
		if page <= 0 {
			o.ServeError(http.StatusBadRequest, "page must be greater than 0")
		}
		perpage, err := o.GetInt("perpage")
		if err != nil {
			o.ServeError(http.StatusBadRequest, "failed to get perpage param: "+err.Error())
		}
		if perpage <= 0 {
			o.ServeError(http.StatusBadRequest, "perpage must be greater than 0")
		}
		var result = make(map[string]interface{})
		total, apps, err := models.GetAllApp(page, perpage)
		if err != nil {
			o.ServeError(http.StatusBadRequest, "failed to get apps: "+err.Error())
		}
		if apps == nil {
			apps = make([]models.App, 0)
		}
		result["total"] = total
		result["count"] = len(apps)
		result["data"] = apps
		o.Serve(result)
	} else {
		app, err := models.GetAppByName(name)
		if err != nil {
			o.ServeError(http.StatusBadRequest, "failed to get app: "+err.Error())
		}
		o.Serve(app)
	}
}

// @router /rasp [get]
func (o *AppController) GetRasps() {
	name := o.GetString("name")
	if len(name) >= 512 {
		o.ServeError(http.StatusBadRequest, "the length of app name must be less than 512")
	}
	page, err := o.GetInt("page")
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get page param: "+err.Error())
	}
	if page <= 0 {
		o.ServeError(http.StatusBadRequest, "page must be greater than 0")
	}
	perpage, err := o.GetInt("perpage")
	if err != nil {
		o.ServeError(http.StatusBadRequest, err.Error())
	}
	if perpage <= 0 {
		o.ServeError(http.StatusBadRequest, "failed to get perpage param: "+"perpage must be greater than 0")
	}

	app, err := models.GetAppByName(name)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get app: "+err.Error())
	}
	if app == nil {
		o.ServeError(http.StatusBadRequest, "the app doesn't exist")
	}
	var result = make(map[string]interface{})
	total, rasps, err := models.GetRaspByAppId(app.Id, page, perpage)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get apps: "+err.Error())
	}
	result["total"] = total
	result["count"] = len(rasps)
	result["data"] = rasps
	o.Serve(result)
}

// @router /config [post]
func (o *AppController) Config() {
	var param map[string]interface{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json format error： "+err.Error())
	}
	appIdParam := param["app_id"]
	if appIdParam == nil {
		o.ServeError(http.StatusBadRequest, "the app_id can not be empty")
	}
	appId, ok := appIdParam.(string)
	if !ok {
		o.ServeError(http.StatusBadRequest, "the app_id must be string")
	}
	app, err := models.GetAppById(appId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get app from mongodb: "+err.Error())
	}

	configParam := param["config"]
	if configParam == nil {
		o.ServeError(http.StatusBadRequest, "the config can not be empty")
	}
	config, ok := configParam.(map[string]interface{})
	if !ok {
		o.ServeError(http.StatusBadRequest, "the type of config must be object")
	}
	validateConfig(config, o)

	configTime := time.Now().UnixNano()
	app.ConfigTime = configTime
	app.Config = config
	err = models.UpdateAppById(app.Id, app)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to update app config: "+err.Error())
	}
	o.Serve(app)
}

// @router / [post]
func (o *AppController) Post() {
	var app = &models.App{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, app)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json format error： "+err.Error())
	}
	if app.Name == "" {
		o.ServeError(http.StatusBadRequest, "app name can not be empty")
	}
	if len(app.Name) >= 512 {
		o.ServeError(http.StatusBadRequest, "the length of app name must be less than 512")
	}
	if app.Description != "" && len(app.Description) >= 1024 {
		o.ServeError(http.StatusBadRequest, "the length of app description must be less than 1024")
	}
	if app.Config != nil {
		validateConfig(app.Config, o)
	} else {
		app.Config = make(map[string]interface{})
	}

	app, err = models.AddApp(app)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "create app failed: "+err.Error())
	}
	o.Serve(app)
}

func validateConfig(config map[string]interface{}, controller *AppController) {
	if config == nil {
		controller.ServeError(http.StatusBadRequest, "the config of app can not be nil")
	}
	for key, value := range config {
		if value == nil {
			controller.ServeError(http.StatusBadRequest, "the value of "+key+" config can not be nil")
		}
		if strings.HasPrefix(key, "hook.white.") {
			whiteUrls, ok := value.([]string)
			if !ok {
				controller.ServeError(http.StatusBadRequest,
					"the type of "+key+" config must be string array")
			}

			if len(whiteUrls) == 0 || len(whiteUrls) > 10 {
				controller.ServeError(http.StatusBadRequest,
					"the count of hook.white's url array must be between (0,10]")
			}
			for _, url := range whiteUrls {
				if len(url) > 200 || len(url) < 1 {
					controller.ServeError(http.StatusBadRequest,
						"the length of hook.white's url must be between [1,200]")
				}
			}
		} else {
			if v, ok := value.(string); ok {
				if len(v) >= 512 {
					controller.ServeError(http.StatusBadRequest,
						"the length of config key "+key+" must less tha 1024")
				}
			}
		}
	}
}
