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
	"rasp-cloud/models"
	"path"
	"net/http"
	"rasp-cloud/controllers"
	"encoding/json"
	"rasp-cloud/mongo"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"bytes"
	"bufio"
	"regexp"
	"github.com/robertkrimen/otto"
)

// Operations about plugin
type PluginController struct {
	controllers.BaseController
}

// @router / [post]
func (o *PluginController) Upload() {
	appId := o.GetString("app_id")
	if appId == "" {
		o.ServeError(http.StatusBadRequest, "app_id can not be empty")
	}
	_, err := models.GetAppById(appId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get app: "+err.Error())
	}
	uploadFile, info, err := o.GetFile("plugin")
	defer uploadFile.Close()
	if uploadFile == nil {
		o.ServeError(http.StatusBadRequest, "must have the plugin parameter")
	}
	if err != nil {
		o.ServeError(http.StatusBadRequest, "parse uploadFile error: "+err.Error())
	}
	if info.Size == 0 {
		o.ServeError(http.StatusBadRequest, "the upload file cannot be empty")
	}
	fileName := info.Filename
	if len(fileName) <= 0 || len(fileName) > 50 {
		o.ServeError(http.StatusBadRequest, "the length of upload uploadFile name must be (0,50]")
	}
	if path.Ext(fileName) != ".js" {
		o.ServeError(http.StatusBadRequest, "the upload file name suffix must be .js")
	}
	pluginContent, err := ioutil.ReadAll(uploadFile)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to read upload plugin: "+err.Error())
	}
	pluginReader := bufio.NewReader(bytes.NewReader(pluginContent))
	firstLine, err := pluginReader.ReadString('\n')
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to read the plugin.js in the zip file: "+err.Error())
	}
	var newVersion string
	if versionArr := regexp.MustCompile(`'.+'|".+"`).FindAllString(firstLine, -1); len(versionArr) > 0 {
		newVersion = versionArr[0][1 : len(versionArr[0])-1]
	} else {
		o.ServeError(http.StatusBadRequest, "failed to find the plugin version: "+err.Error())
	}
	var algorithmStartMsg = "// BEGIN ALGORITHM CONFIG //"
	var algorithmEndMsg = "// END ALGORITHM CONFIG //"
	algorithmStart := bytes.Index(pluginContent, []byte(algorithmStartMsg))
	if algorithmStart < 0 {
		o.ServeError(http.StatusBadRequest, "failed to find the start of algorithmConfig variable: "+algorithmStartMsg)
	}
	algorithmStart = algorithmStart + len([]byte(algorithmStartMsg))
	algorithmEnd := bytes.Index(pluginContent, []byte(algorithmEndMsg))
	if algorithmEnd < 0 {
		o.ServeError(http.StatusBadRequest, "failed to find the end of algorithmConfig variable: "+algorithmEndMsg)
	}
	jsVm := otto.New()
	_, err = jsVm.Run(string(pluginContent[algorithmStart:algorithmEnd]) + "\n algorithmContent=JSON.stringify(algorithmConfig)")
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get algorithm config from plugin: "+err.Error())
	}
	algorithmContent, err := jsVm.Get("algorithmContent")
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get algorithm config from plugin: "+err.Error())
	}
	var algorithmData map[string]interface{}
	err = json.Unmarshal([]byte(algorithmContent.String()), &algorithmData)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to unmarshal algorithm json data: "+err.Error())
	}
	latestPlugin, err := models.AddPlugin(newVersion, pluginContent, appId, algorithmData)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to add plugin to mongodb: "+err.Error())
	}
	o.Serve(latestPlugin)
}

// @router /get [post]
func (o *PluginController) Get() {
	var param map[string]string
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json format error： "+err.Error())
	}
	pluginId := param["id"]
	if pluginId == "" {
		o.ServeError(http.StatusBadRequest, "plugin_id cannot be empty")
	}
	plugin, err := models.GetPluginById(pluginId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get plugin: "+err.Error())
	}
	o.Serve(plugin)
}

// @router /delete [post]
func (o *PluginController) Delete() {
	var param map[string]string
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json format error： "+err.Error())
	}
	pluginId := param["id"]
	if pluginId == "" {
		o.ServeError(http.StatusBadRequest, "plugin_id cannot be empty")
	}
	var app *models.App
	err = mongo.FindOne("app", bson.M{"selected_plugin_id": pluginId}, &app)
	if err != nil && err != mgo.ErrNotFound {
		o.ServeError(http.StatusBadRequest, "failed to get app: "+err.Error())
	}
	if app != nil {
		o.ServeError(http.StatusBadRequest, "failed to delete the plugin,it is used by app: "+app.Id)
	}
	err = models.DeletePlugin(pluginId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to delete plugin: "+err.Error())
	}
	o.ServeWithEmptyData()
}
