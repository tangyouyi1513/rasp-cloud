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
