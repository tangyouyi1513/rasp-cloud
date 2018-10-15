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

package logs

import (
	"github.com/astaxie/beego"
	"rasp-cloud/tools"
	"github.com/astaxie/beego/logs"
	"os"
	"path"
	"rasp-cloud/es"
	"time"
)

var (
	AttackAlarmType = "attack-alarm"
	PolicyAlarmType = "policy-alarm"
	AddAlarmFunc    func(string, []byte)
	raspLoggers     = make(map[string]*logs.BeeLogger)
)

func init() {
	es.RegisterTTL(24*365*time.Hour, AliasAttackIndexName+"-*")
	es.RegisterTTL(24*365*time.Hour, AliasPolicyIndexName+"-*")
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
	err := logger.SetLogger(tools.AdapterAlarmFile,
		`{"filename":"`+logPath+`", "daily":true, "maxdays":10, "perm":"0777"}`)
	if err != nil {
		tools.Panic("failed to init rasp log: " + err.Error())
	}
	return logger
}

func AddLogWithLogstash(alarmType string, content []byte) {
	if logger, ok := raspLoggers[alarmType]; ok && logger != nil {
		_, err := logger.Write(content)
		if err != nil {
			logs.Error("failed to write rasp log: " + err.Error())
		}
	} else {
		logs.Error("failed to write rasp log ,unrecognized log type: " + alarmType)
	}
}

func AddLogWithKafka(alarmType string, content []byte) {

}
