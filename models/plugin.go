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

package models

import (
	"rasp-cloud/tools"
	"fmt"
	"crypto/md5"
	"rasp-cloud/mongo"
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type Plugin struct {
	Version string `json:"version" bson:"version"`
	Md5     string `json:"md5,omitempty" bson:"_id"`
	Content string `json:"plugin,omitempty" bson:"content"`
}

const (
	pluginCollectionName = "plugin"
)

var (
	mutex      sync.Mutex
	MaxPlugins int
)

func init() {
	createIndex()
	if value, err := beego.AppConfig.Int("MaxPlugins"); err != nil || value <= 0 {
		MaxPlugins = 50
	} else {
		MaxPlugins = value
	}
}

// create mongo index for plugin collection
func createIndex() {
	count, err := mongo.Count(pluginCollectionName)
	if err != nil {
		tools.Panic("failed to get rasp collection count")
	}
	if count <= 0 {
		index := &mgo.Index{
			Key:        []string{"-version"},
			Unique:     true,
			Background: true,
			Name:       "plugin_version",
		}
		mongo.CreateIndex(pluginCollectionName, index)
		if err != nil {
			tools.Panic("failed to create index for plugin collection")
		}
	}
}

func AddPlugin(version string, content []byte) (plugin *Plugin, err error) {
	newMd5 := fmt.Sprintf("%x", md5.Sum(content))
	plugin = &Plugin{Version: version, Md5: newMd5, Content: string(content)}
	mutex.Lock()
	defer mutex.Unlock()
	var count int
	if count, err = mongo.Count(pluginCollectionName); err != nil {
		return
	}
	if count > MaxPlugins-1 {
		var oldPlugins []Plugin
		err = mongo.FindAllBySort(pluginCollectionName, nil, 0,
			count+1-MaxPlugins, &oldPlugins, "version")
		if err != nil {
			return
		}
		for _, oldPlugin := range oldPlugins {
			err = mongo.Remove(pluginCollectionName, bson.M{"_id": oldPlugin.Md5})
			if err != nil {
				return
			}
		}
	}
	err = mongo.Insert(pluginCollectionName, plugin)
	return
}

func GetLatestPlugin() (plugin *Plugin, err error) {
	err = mongo.FindOneBySort(pluginCollectionName, bson.M{}, &plugin, "-version")
	return
}

func GetPluginByVersion(version string) (plugin *Plugin, err error) {
	err = mongo.FindOne(pluginCollectionName, bson.M{"version": version}, plugin)
	return
}

func GetAllPlugin() (plugins []Plugin, err error) {
	newSession := mongo.NewSession()
	defer newSession.Close()
	err = newSession.DB(mongo.DbName).C(pluginCollectionName).Find(nil).All(&plugins)
	return
}

func NewPlugin(version string, content []byte) *Plugin {
	newMd5 := fmt.Sprintf("%x", md5.Sum(content))
	return &Plugin{Version: version, Md5: newMd5, Content: string(content)}
}
