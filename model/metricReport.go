package model

type Metric struct {
	Id   int          `json:"id"`
	Data MetricReport `json:"data"`
}

type MetricReport struct {
	ODataType              string                 `json:"@odata.type"`
	ODataContext           string                 `json:"@odata.context"`
	ODataID                string                 `json:"@odata.id"`
	ID                     string                 `json:"Id"`
	Name                   string                 `json:"Name"`
	ReportSequence         string                 `json:"ReportSequence"`
	Timestamp              string                 `json:"Timestamp"`
	MetricReportDefinition MetricReportDefinition `json:"MetricReportDefinition"`
	MetricValues           []MetricValue          `json:"MetricValues"`
	MetricValuesCount      int                    `json:"MetricValues@odata.count"`
	Oem                    OEM                    `json:"Oem"`
}

type MetricReportDefinition struct {
	ODataID string `json:"@odata.id"`
}

type MetricValue struct {
	MetricID    string `json:"MetricId"`
	Timestamp   string `json:"Timestamp"`
	MetricValue string `json:"MetricValue"`
	Oem         OEM    `json:"Oem"`
}

type OEM struct {
	Dell DellOEM `json:"Dell"`
}

type DellOEM struct {
	ODataType                    string `json:"@odata.type"`
	ContextID                    string `json:"ContextID"`
	Label                        string `json:"Label"`
	Source                       string `json:"Source"`
	FQDD                         string `json:"FQDD"`
	ServiceTag                   string `json:"ServiceTag"`
	MetricReportDefinitionDigest string `json:"MetricReportDefinitionDigest"`
	iDRACFirmwareVersion         string `json:"iDRACFirmwareVersion"`
}
