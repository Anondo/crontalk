package translator

import (
	"crontalk/helper"
	"net/url"
	"strconv"
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
	day         = "day"
	hour        = "hour"
	minute      = "minute"
	anyValue    = "*"
)

var (
	// CronExprsn is the cron expression
	CronExprsn       string
	cronSlice        []string
	translatedString string
	moments          = map[int]string{
		minuteIndex: minute,
		hourIndex:   hour,
		dayIndex:    day,
		monthIndex:  month,
		weekIndex:   week,
	}
)

// Validate validates the cron expression provided
func Validate() url.Values {
	CronExprsn = helper.TrimExtraSpaces(CronExprsn)
	cronSlice = strings.Split(CronExprsn, " ")

	errs := url.Values{}
	// checking the length of the expression
	if len(cronSlice) != 5 {
		errs.Add("Expression Values", "A cron expression must contain 5 values/sub-expression")
		return errs
	}
	// checking the values provided for the expression
	for i := minuteIndex; i <= weekIndex; i++ {
		if cronSlice[i] != anyValue {
			if !helper.IsDigit(cronSlice[i]) { //the value provided must be a digit
				errs.Add(moments[i]+" value", "The value must a numeric digit or *")
			} else { //checking the validity of the values in the context of each sub-expressions
				v, _ := strconv.Atoi(cronSlice[i])
				if moments[i] == minute {
					if v < 0 || v > 59 {
						errs.Add(minute+" value", "The value must be between 0 to 59")
					}
				} else if moments[i] == hour {
					if v < 0 || v > 23 {
						errs.Add(hour+" value", "The value must be between 0 to 23")
					}
				} else if moments[i] == day {
					if v < 1 || v > 31 {
						errs.Add(day+" value", "The value must be between 1 to 31")
					}
				} else if moments[i] == month {
					if v < 1 || v > 12 {
						errs.Add(month+" value", "The value must be between 1 to 12")
					}
				} else if moments[i] == week {
					if v < 0 || v > 6 {
						errs.Add(week+" value", "The Value must be between 0 to 6")
					}
				}
			}
		}
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
