package logs

import (
	"rasp-cloud/tools"
)

type RaspLog struct {
	content string
}

var (
	policyIndexName      = "openrasp-policy-alarm"
	aliasPolicyIndexName = "real-openrasp-policy-alarm"
)

func init() {
	err := tools.CreateEsIndex(policyIndexName, aliasPolicyIndexName)
	if err != nil {
		tools.Panic("failed to create index " + aliasPolicyIndexName + ": " + err.Error())
	}
}

func AddPolicyAlarm(content []byte) {
	AddAlarmFunc(PolicyAlarmType, content)
}
