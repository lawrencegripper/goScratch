package main

//ActiveAlerts used to track alerts
var activeAlerts map[string]AlertModel

//ActiveAlerts get the currently configured alerts
func ActiveAlerts() map[string]AlertModel {
	return activeAlerts
}

//SaveAlert perists the alert definition
func SaveAlert(a AlertModel) {
	if activeAlerts == nil {
		activeAlerts = make(map[string]AlertModel)
	}

	activeAlerts[a.ID] = a
}
