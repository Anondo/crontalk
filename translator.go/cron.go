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
	week        = "Week"
	month       = "Month"
	day         = "Day"
	hour        = "Hour"
	minute      = "Minute"
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
			cc, _ := helper.GetList(cronSlice[i], ",")
			for _, c := range cc { //iterating because values can be listed
				vv, ranged := helper.GetList(c, "-")
				if !helper.IsDigit(c) && !ranged { //the value provided must be a digit
					errs.Add(moments[i]+" value", "The value must a numeric digit or *")
				} else { //checking the validity of the values in the context of each sub-expressions
					var v, vr1, vr2 int
					if ranged {
						vr1, _ = strconv.Atoi(vv[0])
						vr2, _ = strconv.Atoi(vv[1])
					} else {
						v, _ = strconv.Atoi(c)
					}

					if moments[i] == minute {
						if (v < 0 || v > 59) && !ranged {
							errs.Add(minute+" value", "The value must be between 0 to 59")
						}
						if ranged {
							if (vr1 < 0 || vr1 > 59) && (vr2 < 0 || vr2 > 59) {
								errs.Add(minute+" value", "The value must be between 0 to 59")
							}
						}
					} else if moments[i] == hour {
						if (v < 0 || v > 23) && !ranged {
							errs.Add(hour+" value", "The value must be between 0 to 23")
						}
						if ranged {
							if (vr1 < 0 || vr1 > 23) && (vr2 < 0 || vr2 > 23) {
								errs.Add(hour+" value", "The value must be between 0 to 23")
							}
						}
					} else if moments[i] == day {
						if (v < 1 || v > 31) && !ranged {
							errs.Add(day+" value", "The value must be between 1 to 31")
						}
						if ranged {
							if (vr1 < 1 || vr1 > 31) && (vr2 < 1 || vr2 > 31) {
								errs.Add(day+" value", "The value must be between 1 to 31")
							}
						}
					} else if moments[i] == month {
						if (v < 1 || v > 12) && !ranged {
							errs.Add(month+" value", "The value must be between 1 to 12")
						}
						if ranged {
							if (vr1 < 1 || vr1 > 12) && (vr2 < 1 || vr2 > 12) {
								errs.Add(month+" value", "The value must be between 1 to 12")
							}
						}
					} else if moments[i] == week {
						if (v < 0 || v > 6) && !ranged {
							errs.Add(week+" value", "The Value must be between 0 to 6")
						}
						if ranged {
							if (vr1 < 0 || vr1 > 6) && (vr2 < 0 || vr2 > 6) {
								errs.Add(week+" value", "The value must be between 0 to 6")
							}
						}
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
