package models

import "rasp-cloud/es"

type ReportData struct {
	RaspId     string `json:"rasp_id"`
	Time       int64  `json:"time"`
	RequestSum int64  `json:"request_sum"`
}

var (
	ReportIndexName      = "openrasp-report-data"
	AliasReportIndexName = "real-openrasp-report-data"
)

func AddReportData(reportData *ReportData, appId string) error {
	return es.Insert(AliasReportIndexName+"-"+appId, "doc", reportData)
}
