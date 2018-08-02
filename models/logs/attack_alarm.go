package logs

import (
	"rasp-cloud/tools"
)

type AttackAlarm struct {
	content string
}

var (
	attackIndexName      = "openrasp-attack-alarm"
	aliasAttackIndexName = "real-openrasp-attack-alarm"
)

func init() {
	err := tools.CreateEsIndex(attackIndexName, aliasAttackIndexName)
	if err != nil {
		tools.Panic("failed to create index " + aliasAttackIndexName + ": " + err.Error())
	}
}

func AddAttackAlarm(content []byte) {
	AddAlarmFunc(AttackAlarmType, content)
}
