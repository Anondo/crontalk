package translator

import (
	"crontalk/helper"
	"net/url"
	"strings"
)

const (
	minuteIndex = 0
	hourIndex   = 1
	dayIndex    = 2
	monthIndex  = 3
	weekIndex   = 4
	week        = "week"
	month       = "month"
	anyValue    = "*"
)

var (
	// CronExprsn is the cron expression
	CronExprsn       string
	cronSlice        []string
	translatedString string
	moments          = map[int]string{
		minuteIndex: "minute",
		hourIndex:   "hour",
		dayIndex:    "day",
		monthIndex:  "month",
		weekIndex:   "week",
	}
)

// Validate validates the cron expression provided
func Validate() url.Values {
	CronExprsn = helper.TrimExtraSpaces(CronExprsn)
	cronSlice = strings.Split(CronExprsn, " ")

	errs := url.Values{}
	if len(cronSlice) != 5 {
		errs.Add("Expression Values", "A cron expression must contain 5 values/sub-expression")
	}
	return errs
}

// Translate does everything to translate a cron expression to english sentence
func Translate() error {

	// translate the base occurence
	if err := translateBaseOccurence(); err != nil {
		return err
	}
	//translate every other occurence
	if err := translateAllButBaseTimeOccurence(); err != nil {
		return err
	}
	// translate the time at the very end
	if err := translateTimeOccurence(); err != nil {
		return err
	}
	return nil

}

// GetTranslatedStr returns the translated string
func GetTranslatedStr() string {
	return translatedString
}
