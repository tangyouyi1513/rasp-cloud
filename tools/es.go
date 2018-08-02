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
		Panic("init ES failed")
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
