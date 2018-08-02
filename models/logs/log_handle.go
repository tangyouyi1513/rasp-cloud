package logs

import (
	"github.com/astaxie/beego"
	"rasp-cloud/tools"
	"github.com/astaxie/beego/logs"
	"os"
	"path"
)

var (
	AttackAlarmType = "attack-alarm"
	PolicyAlarmType = "policy-alarm"
	AddAlarmFunc    func(string, []byte)
	raspLoggers     = make(map[string]*logs.BeeLogger)
)

func init() {
	if beego.AppConfig.String("RaspLogMode") == "logstash" ||
		beego.AppConfig.String("RaspLogMode") == "" {
		AddAlarmFunc = AddLogWithLogstash
		initRaspLoggers()
	} else if beego.AppConfig.String("RaspLogMode") == "kafka" {
		AddAlarmFunc = AddLogWithKafka
	} else {
		tools.Panic("Unrecognized the value of RaspLogMode config")
	}
}

func initRaspLoggers() {
	raspLoggers[AttackAlarmType] = initRaspLogger("openrasp-logs/attack-alarm", "attack.log")
	raspLoggers[PolicyAlarmType] = initRaspLogger("openrasp-logs/policy-alarm", "policy.log")
}

func initRaspLogger(dirName string, fileName string) *logs.BeeLogger {
	if isExists, _ := tools.PathExists(dirName); !isExists {
		err := os.MkdirAll(dirName, os.ModePerm)
		if err != nil {
			tools.Panic(err.Error())
		}
	}
	logger := logs.NewLogger()
	logPath := path.Join(dirName, fileName)
	err := logger.SetLogger(logs.AdapterFile,
		`{"filename":"`+logPath+`", "daily":true, "maxdays":10, "perm":"0777"}`)
	logger.SetPrefix("")
	if err != nil {
		tools.Panic("failed to init rasp log: " + err.Error())
	}
	return logger
}

func AddLogWithLogstash(alarmType string, content []byte) {
	if logger, ok := raspLoggers[alarmType]; ok && logger != nil {
		logger.Write(content)
	} else {
		logs.Error("failed to write rasp log ,unrecognized log type: " + alarmType)
	}
}

func AddLogWithKafka(alarmType string, content []byte) {

}
