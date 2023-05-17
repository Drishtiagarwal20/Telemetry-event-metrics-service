package main

import (
	"encoding/json"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	_ "telemetry-metrics-events-service/docs"
	"telemetry-metrics-events-service/redfish"
)

const (
	events   = "https://10.236.78.122/redfish/v1/SSE?$filter=(EventType%20eq%20%27Other%27)"
	metrics  = "https://10.236.78.122/redfish/v1/SSE?$filter=EventFormatType%20eq%20MetricReport"
	username = "root"
	password = "calvin"
)

//	@title			Telemetry Service
//	@version		1.0
//	@description	This is a sample telemetry service to generate the events and metric reports.
//
// @host localhost:8084
// @BasePath /
func main() {
	fmt.Println("API integration.")

	http.HandleFunc("/v1/events", eventsHandler)
	http.HandleFunc("/v1/metrics", metricHandler)
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The generated Swagger JSON file
	))
	http.ListenAndServe(":8084", nil)
}

// Events
// @Summary Get the Event data
// @Description Retrieves the events data for the subscribed client
// @Router /v1/events [get]
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {object} model.Events
// @Failure 400 {string} string "Bad Request Error"
// @Failure 500 {string} string "Internal Server Error"
func eventsHandler(w http.ResponseWriter, r *http.Request) {
	eventsRes := redfish.EventClient()
	res, err := json.Marshal(eventsRes)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// MetricReport
// @Summary Get the metric report
// @Description Retrieves the report for the subscribed client
// @Router /v1/metrics [get]
// @Tags Metrics
// @Accept json
// @Produce json
// @Success 200 {object} model.Metric
// @Failure 400 {string} string "Bad Request Error"
// @Failure 500 {string} string "Internal Server Error"
func metricHandler(w http.ResponseWriter, r *http.Request) {
	metricRes := redfish.MetricReportClient()
	res, err := json.Marshal(metricRes)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
