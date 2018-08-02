package logs

import (
	"rasp-cloud/controllers"
	"encoding/json"
	"rasp-cloud/models/logs"
	"net/http"
)

// Operations about attack alarm message
type AttackAlarmController struct {
	controllers.BaseController
}

// @router / [post]
func (o *AttackAlarmController) Post() {
	var alarms []interface{}
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &alarms); err != nil {
		o.ServeError(http.StatusBadRequest, "json format error")
		return
	}
	count := 0
	for _, alarm := range alarms {
		content, err := json.Marshal(alarm)
		if err == nil {
			logs.AddAttackAlarm(content)
			count++
		}
	}
	o.Serve(map[string]uint64{"count": uint64(count)})
}
