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

package agent

import (
	"rasp-cloud/models"
	"net/http"
	"encoding/json"
	"rasp-cloud/controllers"
	"gopkg.in/mgo.v2"
	"time"
)

// Operations about plugin
type HeartbeatController struct {
	controllers.BaseController
}

// @router / [post]
func (o *HeartbeatController) Post() {
	var heartbeat map[string]interface{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &heartbeat)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json format error： "+err.Error())
	}
	raspIdParam := heartbeat["rasp_id"]
	if raspIdParam == nil {
		o.ServeError(http.StatusBadRequest, "rasp_id can not be empty")
	}
	raspId, ok := raspIdParam.(string)
	if !ok {
		o.ServeError(http.StatusBadRequest, "the type of rasp_id must be string")
	}
	rasp, err := models.GetRaspById(raspId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get rasp: "+err.Error())
	}
	rasp.LastHeartbeatTime = time.Now().Unix()
	err = models.UpsertRaspById(raspId, rasp)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to upsert rasp: "+err.Error())
	}

	pluginVersion := heartbeat["plugin_version"]
	if pluginVersion == nil {
		o.ServeError(http.StatusBadRequest, "plugin_version can not be empty")
	}
	pluginVersion, ok = pluginVersion.(string)
	if !ok {
		o.ServeError(http.StatusBadRequest, "the type of plugin_version must be string")
	}
	configTimeParam := heartbeat["config_time"]
	if configTimeParam == nil {
		o.ServeError(http.StatusBadRequest, "config_time can not be empty")
	}
	configTime, ok := configTimeParam.(float64)
	if !ok {
		o.ServeError(http.StatusBadRequest, "the type of config_time must be integer")
	}

	appId := o.Ctx.Input.Header("X-OpenRASP-AppID")
	app, err := models.GetAppById(appId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "can not get the app: "+err.Error())
	}
	if app == nil {
		o.ServeError(http.StatusBadRequest, "can not get the app： "+app.Id)
	}

	var result = make(map[string]interface{})
	// 处理插件
	latestPlugin, err := models.GetLatestPlugin()
	if err != nil && err != mgo.ErrNotFound {
		o.ServeError(http.StatusBadRequest, "failed to get latest plugin： "+err.Error())
	}
	if latestPlugin != nil && pluginVersion.(string) < latestPlugin.Version {
		result["plugin"] = latestPlugin
	}
	// 处理配置
	if app.ConfigTime > 0 && app.ConfigTime > int64(configTime) {
		result["config_time"] = app.ConfigTime
		result["config"] = app.Config
	}

	o.Serve(result)
}
