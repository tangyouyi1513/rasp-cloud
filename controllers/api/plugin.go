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
	"bufio"
	"regexp"
	"rasp-cloud/controllers"
	"io/ioutil"
	"io"
	"gopkg.in/mgo.v2"
)

// Operations about plugin
type PluginController struct {
	controllers.BaseController
}

// @router /upload [post]
func (o *PluginController) Upload() {
	file, info, err := o.GetFile("plugin")
	if file == nil {
		o.ServeError(http.StatusBadRequest, "must have the plugin parameter")
	}
	defer file.Close()
	if err != nil {
		o.ServeError(http.StatusBadRequest, "parse file error: "+err.Error())
	}
	if info.Size == 0 {
		o.ServeError(http.StatusBadRequest, "upload file can not be empty")
	}
	fileName := info.Filename
	if len(fileName) <= 0 || len(fileName) > 10 {
		o.ServeError(http.StatusBadRequest, "the length of upload file name must be (0.50]")
	}
	if path.Ext(fileName) != ".js" {
		o.ServeError(http.StatusBadRequest, "the file name suffix must be .js")
	}

	firstLine, err := bufio.NewReader(file).ReadString('\n')
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to read the plugin: "+err.Error())
	}

	var newVersion string
	if versionArr := regexp.MustCompile(`'.+'|".+"`).FindAllString(firstLine, -1); len(versionArr) > 0 {
		newVersion = versionArr[0][1 : len(versionArr[0])-1]
	} else {
		o.ServeError(http.StatusBadRequest, "failed to find the plugin version: "+err.Error())
	}

	plugin, err := models.GetLatestPlugin()
	if err != nil && err != mgo.ErrNotFound {
		o.ServeError(http.StatusBadRequest, "failed to get latest plugin: "+err.Error())
	}
	if plugin != nil && plugin.Version >= newVersion {
		o.ServeError(http.StatusBadRequest, "the file version must be larger than the current version")
	}

	file.Seek(0, io.SeekStart)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to read upload file: "+err.Error())
	}
	latestPlugin, err := models.AddPlugin(newVersion, content)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to add plugin to mongodb: "+err.Error())
	}
	o.Serve(latestPlugin)

}

// 如果不加参数返回，最新插件
// 如果加 version 插件返回响应版本的插件
// @router / [get]
func (o *PluginController) Get() {
	version := o.GetString("version")
	if len(version) == 0 {
		plugins, err := models.GetAllPlugin()
		if err != nil {
			o.ServeError(http.StatusBadRequest, "failed to get latest plugin from mongodb: "+err.Error())
		}
		if plugins == nil {
			o.Serve([]models.Plugin{})
		} else {
			o.Serve(plugins)
		}
	} else {
		plugin, err := models.GetPluginByVersion(version)
		if err != nil {
			o.ServeError(http.StatusBadRequest, "failed to get plugin from mongodb: "+err.Error())
		}
		if plugin == nil {
			o.Serve(make(map[string]interface{}))
		} else {
			o.Serve(plugin)
		}
	}
}

// 获取所有历史版本插件
// @router /latest [get]
func (o *PluginController) GetLatest() {
	plugin, err := models.GetLatestPlugin()
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get latest plugin from mongodb: "+err.Error())
	}
	if plugin == nil {
		o.Serve(make(map[string]interface{}))
	} else {
		o.Serve(plugin)
	}
}
