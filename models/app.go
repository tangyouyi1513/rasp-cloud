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
	"rasp-cloud/mongo"
	"fmt"
	"strconv"
	"time"
	"math/rand"
	"rasp-cloud/tools"
	"gopkg.in/mgo.v2"
	"crypto/sha1"
	"gopkg.in/mgo.v2/bson"
	"rasp-cloud/es"
	"rasp-cloud/models/logs"
)

type App struct {
	Id          string                 `json:"id" bson:"_id"`
	Name        string                 `json:"name"  bson:"name"`
	Description string                 `json:"description"  bson:"description"`
	ConfigTime  int64                  `json:"config_time"  bson:"config_time"`
	Config      map[string]interface{} `json:"config"  bson:"config"`
}

const (
	appCollectionName = "app"
)

func init() {
	count, err := mongo.Count(appCollectionName)
	if err != nil {
		tools.Panic("failed to get app collection count")
	}
	if count <= 0 {
		index := &mgo.Index{
			Key:        []string{"name"},
			Unique:     true,
			Background: true,
			Name:       "app_name",
		}
		err = mongo.CreateIndex(appCollectionName, index)
		if err != nil {
			tools.Panic("failed to create index for app collection")
		}
	}
}

func AddApp(app *App) (result *App, err error) {
	if app.Id == "" {
		app.Id = generateAppId(app)
	}
	err = es.CreateEsIndex(logs.PolicyIndexName+"-"+app.Id, logs.AliasPolicyIndexName+"-"+app.Id, logs.PolicyEsMapping)
	if err != nil {
		return
	}
	err = es.CreateEsIndex(logs.AttackIndexName+"-"+app.Id, logs.AliasAttackIndexName+"-"+app.Id, logs.AttackEsMapping)
	if err != nil {
		return
	}
	err = es.CreateEsIndex(ReportIndexName+"-"+app.Id, AliasReportIndexName+"-"+app.Id, ReportEsMapping)
	if err != nil {
		return
	}
	// ES must be created before mongo
	err = mongo.UpsertId(appCollectionName, app.Id, app)
	if err != nil {
		return
	}
	result = app
	return
}

func generateAppId(app *App) string {
	random := "openrasp_app" + app.Name + strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(rand.Intn(10000))
	return fmt.Sprintf("%x", sha1.Sum([]byte(random)))
}

func GetAllApp(page int, perpage int) (count int, result []App, err error) {
	count, err = mongo.FindAll(appCollectionName, nil, &result, perpage*(page-1), perpage)
	return
}

func GetAppByName(name string) (app *App, err error) {
	err = mongo.FindOne(appCollectionName, bson.M{"name": name}, &app)
	return
}

func GetAppById(id string) (app *App, err error) {
	err = mongo.FindOne(appCollectionName, bson.M{"_id": id}, &app)
	return
}

func UpdateAppById(id string, app *App) (err error) {
	return mongo.UpdateId(appCollectionName, id, app)
}

func RemoveAppById(id string) (err error) {
	return mongo.Remove(appCollectionName, bson.M{"_id": id})
}
