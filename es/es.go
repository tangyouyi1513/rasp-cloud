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

package es

import (
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"time"
	"context"
	"github.com/astaxie/beego/logs"
	"rasp-cloud/tools"
)

var (
	ElasticClient *elastic.Client
)

func init() {
	client, err := elastic.NewClient(elastic.SetURL(beego.AppConfig.String("EsAddr")))
	if err != nil {
		tools.Panic("init ES failed: " + err.Error())
	}
	ElasticClient = client
}

func CreateEsIndex(name string, aliasName string) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()
	exists, err := ElasticClient.IndexExists(name).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		createResult, err := ElasticClient.CreateIndex(name).Do(ctx)
		if err != nil {
			return err
		}
		logs.Info("create es index: " + createResult.Index)
		exists, err = ElasticClient.IndexExists(aliasName).Do(ctx)
		if err != nil {
			return err
		}
		if !exists {
			aliasResult, err := ElasticClient.Alias().Add(name, aliasName).Do(ctx)
			if err != nil {
				return err
			}
			logs.Info("create es index alias: " + aliasResult.Index)
		}
	}
	return nil
}

func Insert(index string,docType string, doc interface{}) (err error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))
	defer cancel()
	_, err = ElasticClient.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)
	return
}
