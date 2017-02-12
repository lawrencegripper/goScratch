package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	appInsightsKey   string = "appInsightsKey"
	appInsightsAppId string = "appInsightsAppId"
)

func StartAlertEngine(alerts []AlertModel) []QueryResult {
	results := []QueryResult{}
	resultsChannel := make(chan QueryResult, len(alerts))
	for _, a := range alerts {
		go ExecuteQuery(a, resultsChannel)
	}

	for index := 0; index < len(alerts); index++ {
		i := <-resultsChannel
		results = append(results, i)
	}

	return results
}

func ExecuteQuery(alert AlertModel, resultsChan chan QueryResult) {

	url := fmt.Sprintf("https://api.applicationinsights.io/beta/apps/%v/query?query=%v", os.Getenv(appInsightsAppId), alert.QueryText)

	client := &http.Client{}
	req, errReq := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", os.Getenv(appInsightsKey))
	res, errRes := client.Do(req)

	if errRes != nil || errReq != nil || res.StatusCode != http.StatusOK {
		resultsChan <- QueryResult{}
	}

	var qResults QueryResult
	decoder := json.NewDecoder(res.Body)
	if errRes := decoder.Decode(&qResults); errRes != nil {
		fmt.Printf("Error serialising json %v", errRes)
	}
	resultsChan <- qResults
}
