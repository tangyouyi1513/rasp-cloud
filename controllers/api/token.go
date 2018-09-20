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
)

type TokenController struct {
	controllers.BaseController
}

// @router / [get]
func (o *TokenController) Get() {
	page, err := o.GetInt("page")
	if err != nil {
		o.ServeError(http.StatusBadRequest, err.Error())
	}
	if page <= 0 {
		o.ServeError(http.StatusBadRequest, "page must be greater than 0")
	}
	perpage, err := o.GetInt("perpage")
	if err != nil {
		o.ServeError(http.StatusBadRequest, err.Error())
	}
	if perpage <= 0 {
		o.ServeError(http.StatusBadRequest, "perpage must be greater than 0")
	}
	total, tokens, err := models.GetAllTokent(page, perpage)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get tokens: "+err.Error())
	}
	if tokens == nil {
		tokens = make([]models.Token, 0)
	}
	var result = make(map[string]interface{})
	result["total"] = total
	result["count"] = len(tokens)
	result["data"] = tokens
	o.Serve(result)
}

// @router /new [get]
func (o *TokenController) NewToken() {
	description := o.GetString("description")
	if len(description) > 1024 {
		o.ServeError(http.StatusBadRequest, "the length of the token description must be less than 1024")
	}
	token, err := models.AddToken(description)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to create new token: "+err.Error())
	}
	o.Serve(token)
}

// @router /delete [delete]
func (o *TokenController) Delete() {
	tokenId := o.GetString("token")
	if len(tokenId) == 0 {
		o.ServeError(http.StatusBadRequest, "the token param can not be empty")
	}
	token, err := models.RemoveToken(tokenId)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to remove token: "+err.Error())
	}
	o.Serve(token)
}
