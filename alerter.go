package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/createAlert", func(response http.ResponseWriter, req *http.Request) {

		switch req.Method {
		case http.MethodPut:
			createAlert(response, req)
		default:
			http.Error(response, "Unsupported method.", http.StatusBadRequest)
		}

	})

	alerts := []AlertModel{AlertModel{AlertType: "test"}}
	go StartAlertEngine(alerts)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
