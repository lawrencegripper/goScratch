package main

import "fmt"

var alertType = [3]string{"count", "valueAbove", "valueBelow"}

// AlertModel Used to hold alert queries
type AlertModel struct {
	ID        string
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
