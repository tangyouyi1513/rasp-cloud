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

package fore_logs

import (
	"rasp-cloud/controllers"
	"encoding/json"
	"net/http"
	"rasp-cloud/models"
	"rasp-cloud/models/logs"
)

// Operations about attack alarm message
type AttackAlarmController struct {
	controllers.BaseController
}

// @router /aggr/time [post]
func (o *AttackAlarmController) AggregationWithTime() {
	var param = &logs.AggrTimeParam{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json decode error： "+err.Error())
	}
	if param.AppId != "" {
		_, err = models.GetAppById(param.AppId)
		if err != nil {
			o.ServeError(http.StatusBadRequest, "failed to get the app: "+param.AppId)
		}
	} else {
		param.AppId = "*"
	}
	if param.StartTime <= 0 {
		o.ServeError(http.StatusBadRequest, "start_time must be greater than 0")
	}
	if param.EndTime <= 0 {
		o.ServeError(http.StatusBadRequest, "end_time must be greater than 0")
	}
	if param.StartTime > param.EndTime {
		o.ServeError(http.StatusBadRequest, "start_time cannot be greater than end_time")
	}
	if param.Interval == "" {
		o.ServeError(http.StatusBadRequest, "interval cannot be empty")
	}
	if param.TimeZone == "" {
		o.ServeError(http.StatusBadRequest, "time_zone cannot be empty")
	}
	if len(param.Interval) > 32 {
		o.ServeError(http.StatusBadRequest, "the length of interval cannot be greater than 32")
	}
	if len(param.TimeZone) > 32 {
		o.ServeError(http.StatusBadRequest, "the length of time_zone cannot be greater than 32")
	}
	result, err :=
		logs.AggregationAttackWithTime(param.StartTime, param.EndTime, param.Interval, param.TimeZone, param.AppId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get aggregation from es: "+err.Error())
	}
	o.Serve(result)
}

// @router /aggr/type [post]
func (o *AttackAlarmController) AggregationWithType() {
	var param = &logs.AggrFieldParam{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json decode error： "+err.Error())
	}
	o.validFieldAggrParam(param)
	result, err :=
		logs.AggregationAttackWithType(param.StartTime, param.EndTime, param.Size, param.AppId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get aggregation from es: "+err.Error())
	}
	o.Serve(result)
}

// @router /aggr/ua [post]
func (o *AttackAlarmController) AggregationWithUserAgent() {
	var param = &logs.AggrFieldParam{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json decode error： "+err.Error())
	}
	o.validFieldAggrParam(param)
	result, err :=
		logs.AggregationAttackWithUserAgent(param.StartTime, param.EndTime, param.Size, param.AppId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get aggregation from es: "+err.Error())
	}
	o.Serve(result)
}

// @router /search [post]
func (o *AttackAlarmController) Search() {
	var param = &logs.SearchLogParam{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &param)
	if param.AppId != "" {
		_, err := models.GetAppById(param.AppId)
		if err != nil {
			o.ServeError(http.StatusBadRequest, "cannot get the app: "+param.AppId)
		}
	} else {
		param.AppId = "*"
	}
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json decode error： "+err.Error())
	}
	if param.StartTime < 0 {
		o.ServeError(http.StatusBadRequest, "start_time can not be less than 0")
	}
	if param.EndTime < 0 {
		o.ServeError(http.StatusBadRequest, "end_time can not be less than 0")
	}
	if param.StartTime > param.EndTime {
		o.ServeError(http.StatusBadRequest, "start_time cannot be greater than end_time")
	}
	if param.Page <= 0 {
		o.ServeError(http.StatusBadRequest, "page must be greater than 0")
	}
	if param.Perpage <= 0 {
		o.ServeError(http.StatusBadRequest, "perpage must be greater than 0")
	}
	total, result, err := logs.SearchLogs(param.StartTime, param.EndTime, param.Data, "event_time",
		param.Page, param.Perpage, false, logs.AliasAttackIndexName+"-"+param.AppId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to search data from es: "+err.Error())
	}
	o.Serve(map[string]interface{}{
		"total": total,
		"data":  result,
	})
}

func (o *AttackAlarmController) validFieldAggrParam(param *logs.AggrFieldParam) {
	if param.AppId != "" {
		_, err := models.GetAppById(param.AppId)
		if err != nil {
			o.ServeError(http.StatusBadRequest, "cannot get the app: "+param.AppId)
		}
	} else {
		param.AppId = "*"
	}
	if param.StartTime <= 0 {
		o.ServeError(http.StatusBadRequest, "start_time must be greater than 0")
	}
	if param.EndTime <= 0 {
		o.ServeError(http.StatusBadRequest, "end_time must be greater than 0")
	}
	if param.StartTime > param.EndTime {
		o.ServeError(http.StatusBadRequest, "start_time cannot be greater than end_time")
	}
	if param.Size <= 0 {
		o.ServeError(http.StatusBadRequest, "size must be greater than 0")
	}
}
