package main

import (
	_ "rasp-cloud/routers"
	_ "rasp-cloud/models"
	_ "rasp-cloud/filter"
	_ "rasp-cloud/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"rasp-cloud/controllers"
	"os"
	"rasp-cloud/tools"
)

func main() {
	initLogger()
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

func initLogger() {
	if isExists, _ := tools.PathExists("logs/api"); !isExists {
		err := os.MkdirAll("logs/api", os.ModePerm)
		if err != nil {
			tools.Panic("failed to create logs/api dir")
		}
	}
	logs.SetLogFuncCall(true)
	logs.SetLogger(logs.AdapterFile,
		`{"filename":"logs/api/rasp-cloud.log","daily":true,"maxdays":10,"perm":"0777"}`)
	if beego.BConfig.RunMode == "dev" {
		logs.SetLevel(beego.LevelDebug)
	} else {
		logs.SetLevel(beego.LevelInformational)
	}
}
