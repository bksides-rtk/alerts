package services

import "github.com/RTK-Tickets/alerts/models"

var alerts []models.Alert = []models.Alert{}

func CreateAlert(alert models.Alert) (models.Alert, error) {
	id := len(alerts)
	alert.ID = id
	alerts = append(alerts, alert)
	return alert, nil
}

func GetAlert(id int) (models.Alert, error) {
	return alerts[id], nil
}

func GetAlerts(match map[string]interface{}) ([]models.Alert, error) {
	var alertList []models.Alert
	for _, alert := range alerts {
		if alert.Match(match) {
			alertList = append(alertList, alert)
		}
	}
	return alertList, nil
}
