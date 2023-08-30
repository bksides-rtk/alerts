package models

type Alert struct {
	ID          int      `json:"id"`
	Severity    Severity `json:"severity"`
	Loud        bool     `json:"loud"`
	Description string   `json:"description"`
	ClaimedBy   *User    `json:"claimedBy"`
	Source      string   `json:"source"`
	Dismissed   bool     `json:"dismissed"`
}

func (a *Alert) Match(match map[string]interface{}) bool {
	if id, ok := match["id"]; ok {
		idInt, ok := id.(int)
		if !ok {
			return false
		}
		if idInt != a.ID {
			return false
		}
	}
	if severity, ok := match["severity"]; ok {
		severityStr, ok := severity.(string)
		if !ok {
			return false
		}
		severityEnum, err := ParseSeverity(severityStr)
		if err != nil {
			return false
		}
		if severityEnum != a.Severity {
			return false
		}
	}
	if loud, ok := match["loud"]; ok {
		loudBool, ok := loud.(bool)
		if !ok {
			return false
		}
		if loudBool != a.Loud {
			return false
		}
	}
	if description, ok := match["description"]; ok {
		descriptionStr, ok := description.(string)
		if !ok {
			return false
		}
		if descriptionStr != a.Description {
			return false
		}
	}
	if claimedBy, ok := match["claimedBy"]; ok {
		claimedByUser, ok := claimedBy.(map[string]interface{})
		if !ok {
			return false
		}
		if !a.ClaimedBy.Match(claimedByUser) {
			return false
		}
	}
	if source, ok := match["source"]; ok {
		sourceStr, ok := source.(string)
		if !ok {
			return false
		}
		if sourceStr != a.Source {
			return false
		}
	}
	if dismissed, ok := match["dismissed"]; ok {
		dismissedBool, ok := dismissed.(bool)
		if !ok {
			return false
		}
		if dismissedBool != a.Dismissed {
			return false
		}
	}
	return true
}

type LinkAlert struct {
	Alert
	Href string
}

type OptionAlert struct {
	Alert
	Options []AlertOption
}

type AlertOption struct {
	Text string
	Func func()
}

func (o *AlertOption) Do() {
	o.Func()
}
