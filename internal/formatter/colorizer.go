package formatter

import "github.com/fatih/color"

type thresholds struct {
	coldMax, coolMax, warmMax, hotMax float64
}

func getThresholds(unitSymbol string) thresholds {
	if unitSymbol == "°F" {
		return thresholds{coldMax: 32, coolMax: 59, warmMax: 77, hotMax: 95}
	}
	return thresholds{coldMax: 0, coolMax: 15, warmMax: 25, hotMax: 35}
}

func colorizeTemp(temp float64, unitSymbol string) string {
	thresholds := getThresholds(unitSymbol)
	template := "%.1f %s"
	switch {
	case temp <= thresholds.coldMax:
		return color.New(color.FgCyan).Sprintf(template, temp, unitSymbol)
	case temp <= thresholds.coolMax:
		return color.New(color.FgBlue).Sprintf(template, temp, unitSymbol)
	case temp <= thresholds.warmMax:
		return color.New(color.FgYellow).Sprintf(template, temp, unitSymbol)
	case temp <= thresholds.hotMax:
		return color.New(color.FgRed).Sprintf(template, temp, unitSymbol)
	default:
		return color.New(color.FgHiRed, color.Bold).Sprintf(template, temp, unitSymbol)
	}
}
