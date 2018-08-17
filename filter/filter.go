package filter

import (
	"github.com/astaxie/beego/logs"
	"os"
	"rasp-cloud/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"time"
)

var (
	accessLogger *logs.BeeLogger
)

func init() {
	initAccessLogger()
	beego.InsertFilter("/*", beego.BeforeRouter, logAccess)
}

func logAccess(ctx *context.Context) {
	var cont string
	cont += "[T]" + formatTime(time.Now().Unix(), "15:04:05") + " " + ctx.Input.Method() + " " +
		ctx.Input.Site() + ctx.Input.URL() + " - [I]" + ctx.Input.IP() + " | [U]" + ctx.Input.UserAgent()
	if ctx.Input.Referer() != "" {
		cont += "[F]" + ctx.Input.Referer()
	}

	accessLogger.Info(cont)
}

func formatTime(timestamp int64, format string) (times string) {
	tm := time.Unix(timestamp, 0)
	times = tm.Format(format)
	return
}

func initAccessLogger() {
	if isExists, _ := tools.PathExists("logs/access"); !isExists {
		err := os.MkdirAll("logs/access", os.ModePerm)
		if err != nil {
			tools.Panic("failed to create logs/access dir: " + err.Error())
		}
	}

	accessLogger = logs.NewLogger()
	accessLogger.EnableFuncCallDepth(true)
	accessLogger.SetLogFuncCallDepth(4)
	err := accessLogger.SetLogger(logs.AdapterFile,
		`{"filename":"logs/access/access.log","daily":true,"maxdays":10,"perm":"0777"}`)
	if err != nil {
		tools.Panic("failed to init access log: " + err.Error())
	}
}
