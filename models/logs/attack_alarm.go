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
	"encoding/json"
	"crypto/md5"
	"fmt"
)

type AttackAlarm struct {
	content string
}

var (
	AttackIndexName      = "openrasp-attack-alarm"
	AliasAttackIndexName = "real-openrasp-attack-alarm"
	AttackEsMapping      = `
	{
		"mappings": {
			"_default_": {
				"_all": {
					"enabled": false
				},
				"properties": {
					"request_method": {
						"type": "keyword",
						"ignore_above": 50
					},
					"target": {
						"type": "keyword",
						"ignore_above": 256
					},
					"server_ip": {
						"type": "ip"
					},
					"referer": {
						"type": "keyword",
						"ignore_above": 256
					},
					"user_agent": {
						"type": "keyword",
						"ignore_above": 256
					},
					"attack_source": {
						"type": "ip"
					},
					"path": {
						"type": "keyword",
						"ignore_above": 256
					},
					"url": {
						"type": "keyword",
						"ignore_above": 256
					},
					"event_type": {
						"type": "keyword",
						"ignore_above": 256
					},
					"server_hostname": {
						"type": "keyword",
						"ignore_above": 256
					},
					"stack_md5": {
						"type": "keyword",
						"ignore_above": 64
					},
					"server_type": {
						"type": "keyword",
						"ignore_above": 256
					},
					"server_version": {
						"type": "keyword",
						"ignore_above": 256
					},
					"request_id": {
						"type": "keyword",
						"ignore_above": 256
					},
					"body": {
						"type": "keyword
					},
					"app_id": {
						"type": "keyword",
						"ignore_above": 256
					},
					"rasp_id": {
						"type": "keyword",
						"ignore_above": 256
					},
					"local_ip": {
						"type": "ip"
					},
					"event_time": {
						"type": "date"
					},
					"stack_trace": {
						"type": "keyword"
					},
					"intercept_state": {
						"type": "keyword",
						"ignore_above": 64
					},
					"attack_type": {
						"type": "keyword",
						"ignore_above": 256
					},
					"plugin_name": {
						"type": "keyword",
						"ignore_above": 256
					},
					"plugin_confidence": {
						"type": "short"
					},
					"attack_params": {
						"type": "object"
					},
					"plugin_message": {
						"type": "keyword"
					}
				}
			}
		}
	}
	`
)

func AddAttackAlarm(alarm map[string]interface{}) error {
	if stack, ok := alarm["stack_trace"]; ok && stack != nil {
		_, ok = stack.(string)
		if ok {
			alarm["stack_md5"] = fmt.Sprintf("%x", md5.Sum([]byte(stack.(string))))
		}
	}
	content, err := json.Marshal(alarm)
	if err == nil {
		AddAlarmFunc(AttackAlarmType, content)
	}
	return err
}
