package models

import (
	"fmt"
	"strings"
)

type Severity uint8

const (
	SeverityLow Severity = iota
	SeverityMedium
	SeverityHigh
)

func (s Severity) String() string {
	switch s {
	case SeverityLow:
		return "low"
	case SeverityMedium:
		return "medium"
	default:
		return "high"
	}
}

func ParseSeverity(str string) (Severity, error) {
	switch strings.ToLower(str) {
	case "low":
		return SeverityLow, nil
	case "medium":
		return SeverityMedium, nil
	case "high":
		return SeverityHigh, nil
	}
	return SeverityLow, fmt.Errorf("invalid severity %s", str)
}
