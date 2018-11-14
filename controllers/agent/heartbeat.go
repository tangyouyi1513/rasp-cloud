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

type heartbeatParam struct {
	RaspId        string `json:"rasp_id"`
	PluginVersion string `json:"plugin_version"`
	PluginMd5     string `json:"plugin_md5"`
	ConfigTime    int64  `json:"config_time"`
}

// @router / [post]
func (o *HeartbeatController) Post() {
	var heartbeat heartbeatParam
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &heartbeat)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json format error： "+err.Error())
	}
	rasp, err := models.GetRaspById(heartbeat.RaspId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get rasp: "+err.Error())
	}
	rasp.LastHeartbeatTime = time.Now().Unix()
	rasp.PluginVersion = heartbeat.PluginVersion
	err = models.UpsertRaspById(heartbeat.RaspId, rasp)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to update rasp: "+err.Error())
	}
	pluginMd5 := heartbeat.PluginMd5
	configTime := heartbeat.ConfigTime
	appId := o.Ctx.Input.Header("X-OpenRASP-AppID")
	app, err := models.GetAppById(appId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "cannot get the app: "+err.Error())
	}
	if app == nil {
		o.ServeError(http.StatusBadRequest, "cannot get the app： "+app.Id)
	}

	result := make(map[string]interface{})
	isUpdate := false
	// handle plugin
	selectedPlugin, err := models.GetSelectedPlugin(appId)
	if err != nil && err != mgo.ErrNotFound {
		o.ServeError(http.StatusBadRequest, "failed to get selected plugin： "+err.Error())
	}
	if selectedPlugin != nil {
		if pluginMd5 != selectedPlugin.Md5 {
			isUpdate = true
		}
		if app.ConfigTime > 0 && app.ConfigTime > int64(configTime) {
			isUpdate = true
		}
	}
	if isUpdate {
		for k, v := range app.WhiteListConfig {
			app.GeneralConfig[k] = v
		}
		app.GeneralConfig["algorithm_config"] = selectedPlugin.AlgorithmConfig
		result["plugin"] = selectedPlugin
		result["config_time"] = app.ConfigTime
		result["config"] = app.GeneralConfig
	}
	o.Serve(result)
}
