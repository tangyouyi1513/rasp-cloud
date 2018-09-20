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

package tools

import (
	"github.com/astaxie/beego"
	"github.com/olivere/elastic"
	"time"
	"context"
	"github.com/astaxie/beego/logs"
)

var (
	EsClient *elastic.Client
)

func init() {
	client, err := elastic.NewClient(elastic.SetURL(beego.AppConfig.String("EsAddr")))
	if err != nil {
		Panic("init ES failed: " + err.Error())
	}
	EsClient = client
}

func CreateEsIndex(name string, aliasName string) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))
	defer cancel()
	exist, err := EsClient.IndexExists(name).Do(ctx)
	if err != nil {
		return err
	}
	if !exist {
		createResult, err := EsClient.CreateIndex(name).Do(ctx)
		if err != nil {
			return err
		}
		logs.Info("create es index: " + createResult.Index)
		aliasResult, err := EsClient.Alias().Add(name, aliasName).Do(ctx)
		if err != nil {
			return err
		}
		logs.Info("create es index alias: " + aliasResult.Index)
	}
	return nil
}
