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

package api

import (
	"rasp-cloud/controllers"
	"net/http"
	"rasp-cloud/models"
	"encoding/json"
)

type RaspController struct {
	controllers.BaseController
}

// @router /find [post]
func (o *RaspController) Post() {
	var data map[string]interface{}
	err := json.Unmarshal(o.Ctx.Input.RequestBody, data)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "json format error： "+err.Error())
	}
	pageParam := data["page"]
	if pageParam == nil {
		o.ServeError(http.StatusBadRequest, "the page param can not be empty")
	}
	page, ok := pageParam.(int)
	if !ok {
		o.ServeError(http.StatusBadRequest, "the page param must be integer")
	}
	if page <= 0 {
		o.ServeError(http.StatusBadRequest, "page must be greater than 0")
	}
	perpageParam := data["perpage"]
	if perpageParam == nil {
		o.ServeError(http.StatusBadRequest, "the perpage param can not be empty")
	}
	perpage, ok := perpageParam.(int)
	if !ok {
		o.ServeError(http.StatusBadRequest, "the perpage param must be integer")
	}
	if perpage <= 0 {
		o.ServeError(http.StatusBadRequest, "perpage must be greater than 0")
	}
	raspDataParam := data["data"]
	if raspDataParam == nil {
		o.ServeError(http.StatusBadRequest, "the data param can not be empty")
	}
	raspData, ok := raspDataParam.(map[string]interface{})
	if !ok {
		o.ServeError(http.StatusBadRequest, "the type of data param must be object")
	}
	total, rasps, err := models.FindRasp(raspData, page, perpage)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get rasp from mongodb: "+err.Error())
	}
	var result = make(map[string]interface{})
	result["total"] = total
	result["count"] = len(rasps)
	result["data"] = rasps
	o.Serve(result)
}
