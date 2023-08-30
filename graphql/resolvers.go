package graphql

import (
	"github.com/RTK-Tickets/alerts/models"
	"github.com/RTK-Tickets/alerts/services"
	"github.com/graphql-go/graphql"
)

func createAlert(p graphql.ResolveParams) (interface{}, error) {
	alertMap := p.Args["alert"].(map[string]interface{})

	alert := models.Alert{}
	if id, ok := alertMap["id"].(int); ok {
		alert.ID = id
	}
	if severity, ok := alertMap["severity"].(string); ok {
		var err error
		alert.Severity, err = models.ParseSeverity(severity)
		if err != nil {
			return nil, err
		}
	}
	if loud, ok := alertMap["loud"].(bool); ok {
		alert.Loud = loud
	}
	if description, ok := alertMap["description"].(string); ok {
		alert.Description = description
	}
	if claimedBy, ok := alertMap["claimedBy"].(models.User); ok {
		alert.ClaimedBy = &claimedBy
	}
	if source, ok := alertMap["source"].(string); ok {
		alert.Source = source
	}
	if dismissed, ok := alertMap["dismissed"].(bool); ok {
		alert.Dismissed = dismissed
	}

	return services.CreateAlert(alert)
}

func getAlert(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(int)
	return services.GetAlert(id)
}

func getAlerts(p graphql.ResolveParams) (interface{}, error) {
	match := p.Args["match"].(map[string]interface{})
	return services.GetAlerts(match)
}
