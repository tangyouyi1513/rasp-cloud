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

package models

import (
	"rasp-cloud/es"
	"time"
	"github.com/olivere/elastic"
	"context"
)

type ReportData struct {
	RaspId     string `json:"rasp_id"`
	Time       int64  `json:"time"`
	RequestSum int64  `json:"request_sum"`
	InsertTime  int64  `json:"@timestamp"`
}

var (
	ReportIndexName      = "openrasp-report-data"
	AliasReportIndexName = "real-openrasp-report-data"
	reportType           = "doc"
	ReportEsMapping      = `
		{
			"mappings": {
				"_default_": {
					"_all": {
						"enabled": false
					},
					"properties": {
						"@timestamp":{
							"type":"date"
         				},
						"time": {
							"type": "date"
						},
						"request_sum": {
							"type": "long"
						},
						"rasp_id": {
							"type": "keyword",
							"ignore_above" : 256
						}
					}
				}
			}
		}
	`
)

func init() {
	es.RegisterTTL(24*100*time.Hour, AliasReportIndexName+"-*")
}

func AddReportData(reportData *ReportData, appId string) error {
	reportData.InsertTime = time.Now().UnixNano()/1000
	return es.Insert(AliasReportIndexName+"-"+appId, reportType, reportData)
}

func GetHistoryRequestSum(startTime int64, endTime int64, interval string, appId string, raspId string) (err error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))
	defer cancel()
	sumQuery := elastic.NewBoolQuery()
	if raspId != "" {
		sumQuery.Must(elastic.NewTermQuery("rasp_id", raspId))
	}
	sumQuery.Filter(elastic.NewRangeQuery("time").From(startTime).To(endTime))
	timeSource := elastic.NewCompositeAggregationDateHistogramValuesSource("group_time", interval).
		Field("time")
	sumAggr := elastic.NewSumAggregation().Field("request_sum")
	requestSumAggr := elastic.NewCompositeAggregation().SubAggregation("sum_request", sumAggr).
		Size(10000).Sources(timeSource)
	index := AliasReportIndexName
	if appId == "" {
		index += "-*"
	} else {
		index += "-" + appId
	}
	result, err := es.ElasticClient.Search(index).Type(reportType).
		Query(sumQuery).
		Aggregation("sum_request", requestSumAggr).
		Size(0).
		Do(ctx)
	if err != nil {
		return
	}
	if result != nil {

		//fmt.Printf("%+v", result.Hits.Hits[0])
	}
	return
}
