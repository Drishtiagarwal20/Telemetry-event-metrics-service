package redfish

import (
	"bufio"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"telemetry-metrics-events-service/model"
	"time"
)

const (
	metricReportPath = "/SSE?$filter=EventFormatType%20eq%20MetricReport"
	eventsPath       = "/SSE?$filter=(EventType%20eq%20%27Other%27)"

	redfishApi = "https://10.236.78.122/redfish/v1"

	basicUsername = "root"
	basicPassword = "calvin"
)

func MetricReportClient() *model.Metric {

	//clientId := 545689
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client := &http.Client{Transport: tr}
	request, _ := http.NewRequestWithContext(ctx, "GET", redfishApi+metricReportPath, nil)
	request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(basicUsername+":"+basicPassword)))
	response, err := client.Do(request)

	//defer response.Body.Close()

	if err != nil {
		log.Println("Couldn't connect to the redfish api due to : ", err.Error())
		return nil
	}
	var res []string

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line of response data as needed
		res = append(res, line)
		// Check if the timeout has expired
		if ctx.Err() != nil {
			log.Println("Timeout exceeded. Closing the connection.")
			break
		}
	}
	joinedString := strings.Join(res, " ")
	joinedString = strings.Replace(joinedString, "data:", ",\"data\":", 1)
	joinedString = strings.Replace(joinedString, "id:", "{\"id\":", 1)
	joinedString = joinedString + "}"
	log.Println(joinedString)
	byteSlice := []byte(joinedString)
	var metric model.Metric
	err = json.Unmarshal(byteSlice, &metric)
	if err != nil {
		log.Println(err)
	}
	return &metric
}

func EventClient() *model.Events {

	//clientId := 545689
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := &http.Client{Transport: tr}
	request, _ := http.NewRequestWithContext(ctx, "GET", redfishApi+eventsPath, nil)
	request.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(basicUsername+":"+basicPassword)))
	response, err := client.Do(request)

	//defer response.Body.Close()
	if err != nil {
		log.Println("Couldn't connect to the redfish api due to : ", err.Error())
		return nil
	}
	var res []string

	scanner := bufio.NewScanner(response.Body)
	for scanner.Scan() {
		line := scanner.Text()
		// Process the line of response data as needed
		res = append(res, line)
		// Check if the timeout has expired
		if ctx.Err() != nil {
			log.Println("Timeout exceeded. Closing the connection.")
			break
		}
	}
	joinedString := strings.Join(res, " ")
	joinedString = strings.Replace(joinedString, "data:", ",\"data\":", -1)
	joinedString = strings.Replace(joinedString, "id:", "{\"id\":", 1)
	joinedString = strings.Replace(joinedString, "id:", "@@@@@Null@@@@@", -1)

	log.Println(joinedString)

	parts := strings.Split(joinedString, "@@@@@Null@@@@@")
	joinedString = parts[0]
	log.Println(joinedString)

	joinedString = joinedString + "}"
	byteSlice := []byte(joinedString)

	var events model.Events
	err = json.Unmarshal(byteSlice, &events)
	if err != nil {
		log.Println(err)
	}
	return &events
}
