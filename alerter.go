package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var alertType = [3]string{"count", "valueAbove", "valueBelow"}

// AlertModel Used to hold alert queries
type AlertModel struct {
	QueryText string
	AlertType string
	Threshold int
}

// IsValid checks the model is valid
func (model *AlertModel) IsValid() bool {
	return model.AlertType != "" && model.QueryText != "" && model.Threshold != -1
}

func (model *AlertModel) String() string {
	return fmt.Sprintf("Query %v Type %v Threshold %v", model.QueryText, model.AlertType, model.Threshold)
}

func createAlert(response http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var alertModel AlertModel

	if err := decoder.Decode(&alertModel); err != nil {
		http.Error(response, "Invalid Json supplied, unparsable.", http.StatusBadRequest)
		return
	}

	fmt.Println(alertModel.String())

	if !alertModel.IsValid() {
		http.Error(response, "Invalid Json supplied. Expect AlertType, QueryText and Threshold", http.StatusBadRequest)
		return
	}

	if json, err := json.Marshal(alertModel); err != nil {
		fmt.Fprintf(response, "Error serialising json %v", err)

	} else {
		fmt.Fprintf(response, "Json: %v", string(json))

	}
}

func main() {
	http.HandleFunc("/createAlert", func(response http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodPut:
			createAlert(response, req)
		default:
			http.Error(response, "Unsupported method.", http.StatusBadRequest)
		}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
