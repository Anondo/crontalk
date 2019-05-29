package generator

import "strings"

const (
	minIndex   = 0
	hourIndex  = 1
	dayIndex   = 2
	monthIndex = 3
	weekIndex  = 4
	done       = "done"
)

var (
	cronSlice = []string{"", "", "", "", ""} //empty strings to append the cron values because values can be listed
)

// GenerateCron generates a cron expression by prompting english word options
func GenerateCron() (string, error) {

	if err := runMinute(); err != nil {
		return "", err
	}
	if err := runHour(); err != nil {
		return "", err
	}
	if err := runDay(); err != nil {
		return "", err
	}
	if err := runMonth(); err != nil {
		return "", err
	}
	if err := runWeek(); err != nil {
		return "", err
	}
	cronExpression := strings.Join(cronSlice, " ")
	return cronExpression, nil
}
