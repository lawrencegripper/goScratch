package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
