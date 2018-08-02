package controllers

import (
	"rasp-cloud/tools"
	"rasp-cloud/models"
	"github.com/astaxie/beego"
	"path"
	"os"
	"io/ioutil"
	"net/http"
)

// Operations about plugin
type PluginController struct {
	BaseController
}

var (
	maxPlugins   int
	latestPlugin *models.Plugin
)

func init() {
	if value, err := beego.AppConfig.Int("RaspLogMode"); err != nil || value <= 0 {
		maxPlugins = 50
	} else {
		maxPlugins = value
	}
	if latestPlugin == nil {
		newPlugin, err := models.GetLatestPluginFromDir()
		if err != nil {
			tools.Panic("can not get latest plugin: " + err.Error())
		}
		latestPlugin = newPlugin
	}
}

// @router / [get]
func (o *PluginController) Get() {
	var plugin *models.Plugin
	oldVersion := o.GetString("version")

	if latestPlugin != nil {
		plugin = latestPlugin
	} else {
		plugin = &models.Plugin{Version: oldVersion}
	}

	o.Serve(plugin)
}

// @router / [post]
func (o *PluginController) Post() {
	file, info, err := o.GetFile("plugin")
	if file == nil {
		o.ServeError(http.StatusBadRequest, "must have the plugin parameter")
		return
	}
	defer file.Close()
	if err != nil {
		o.ServeError(http.StatusBadRequest, "parse file error: "+err.Error())
		return
	}
	if info.Size == 0 {
		o.ServeError(http.StatusBadRequest, "upload file can not be empty")
		return
	}
	fileName := info.Filename
	if len(fileName) <= 0 || len(fileName) > 50 {
		o.ServeError(http.StatusBadRequest, "the length of upload file name must be (0.50]")
		return
	}
	if path.Ext(fileName) != ".js" {
		o.ServeError(http.StatusBadRequest, "the file name suffix must be .js")
		return
	}

	jsFiles, err := tools.ListFiles("plugin", "js", models.PluginPrefix)
	if err != nil {
		o.ServeError(http.StatusInternalServerError, "failed to list plugin directory")
		return
	}
	// 超过插件数量上限,删除多于插件
	if len(jsFiles) > maxPlugins-1 && maxPlugins > 0 {
		for _, value := range jsFiles[maxPlugins-1:] {
			if err := os.Remove(value); err != nil {
				o.ServeError(http.StatusInternalServerError, "failed to remove plugin "+value+": "+err.Error())
				return
			}
		}
	}

	newVersion := fileName[0 : len(fileName)-3]
	if newVersion <= latestPlugin.Version {
		o.ServeError(http.StatusBadRequest, "the file version must be larger than the current version")
		return
	}

	// 文件不存在创建文件,存在则覆盖文件
	if err := o.SaveToFile("plugin", "plugin/"+fileName); err != nil {
		o.ServeError(http.StatusInternalServerError, "failed to save upload file: "+err.Error())
		return
	}
	content, err := ioutil.ReadFile("plugin/" + fileName)
	if err != nil {
		o.ServeError(http.StatusInternalServerError, "failed to read new file: "+err.Error())
		return
	}
	latestPlugin = models.NewPlugin(newVersion, content)
	o.Serve(latestPlugin)
}
