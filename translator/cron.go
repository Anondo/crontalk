package translator

import (
	"net/url"
	"strings"

	"github.com/Anondo/crontalk/helper"
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
	every       = "*"
	rangee      = "-"
	step        = "/"
	list        = ","
)

var (
	cronSlice []string
	moments   = map[int]string{
		minuteIndex: minute,
		hourIndex:   hour,
		dayIndex:    day,
		monthIndex:  month,
		weekIndex:   week,
	}
)

// Validate validates the cron expression provided
func (t *Translator) Validate() url.Values {
	t.CronExpression = helper.TrimExtraSpaces(t.CronExpression)
	cronSlice = strings.Split(t.CronExpression, " ")

	for i, cron := range cronSlice { // removing duplicate cron expression values
		cc, _ := helper.GetList(cron, list)
		cronSlice[i] = strings.Join(helper.Unique(cc), ",")
	}

	errs := url.Values{}
	// checking the length of the expression
	if len(cronSlice) != 5 {
		errs.Add("Expression Values", "A cron expression must contain 5 values/sub-expressions")
		return errs
	}
	// checking the values provided for the expression
	for i := minuteIndex; i <= weekIndex; i++ {
		if cronSlice[i] != every {
			cc, _ := helper.GetList(cronSlice[i], list)
			for _, c := range cc { //iterating because values can be listed
				if strings.Contains(c, step) { // if the expression is a stepped value
					validateSteppedSubExpression(&errs, c, moments[i]) // validate just the step value
				}
				slashIndex := helper.IndexOf(strings.Split(c, ""), step)
				if slashIndex != -1 { // just validate everything apart from the step values
					c = c[:slashIndex]
				}
				if c != every {
					validateSubExpressions(&errs, moments[i], c)
				}

			}

		}
	}
	return errs
}

// Translate does everything to translate a cron expression to english sentence
func (t *Translator) Translate() error {
	// translate the base occurence
	if err := t.translateBaseOccurence(); err != nil {
		return err
	}
	//translate every other occurence
	if err := t.translateAllButBaseTimeOccurence(); err != nil {
		return err
	}
	// translate the time at the very end
	if err := t.translateTimeOccurence(); err != nil {
		return err
	}
	return nil

}

// GetTranslatedStr returns the translated string
func (t *Translator) GetTranslatedStr() string {
	return t.translatedStr
}
