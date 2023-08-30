package models

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (u *User) Match(match map[string]interface{}) bool {
	if id, ok := match["id"]; ok {
		idInt, ok := id.(int)
		if !ok {
			return false
		}
		if idInt != u.ID {
			return false
		}
	}
	if firstName, ok := match["firstName"]; ok {
		firstNameStr, ok := firstName.(string)
		if !ok {
			return false
		}
		if firstNameStr != u.FirstName {
			return false
		}
	}
	if lastName, ok := match["lastName"]; ok {
		lastNameStr, ok := lastName.(string)
		if !ok {
			return false
		}
		if lastNameStr != u.LastName {
			return false
		}
	}
	return true
}
