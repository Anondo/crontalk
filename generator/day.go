package generator

import (
	"github.com/Anondo/crontalk/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	everyDay       = "every day"
	inputDay       = "<input>"
	defaultStepDay = "every <input> days"
	stepDay        = "every <input> days from <input> to 31"
	rangedDay      = "<input> to <input> days"
	rangedstepDay  = "every <input> days from <input> to <input>"
)

var (
	dayItems = []string{
		done,
		everyDay,
		inputDay,
		defaultStepDay,
		stepDay,
		rangedDay,
		rangedstepDay,
	}
)

func runDay() error {
	var itm string
	var err error
	for itm != done { // iterating until the done option is selected because of list values
		itm, err = helper.RunPrompt("Day", dayItems)
		if err != nil {
			return err
		}
		if err := parseDayExpression(itm); err != nil {
			fmt.Println(err.Error())
			continue // to avoid adding unwanted ","
		}
		if itm != done { //the check is given because this block will be executed once even after done is selected
			cronSlice[dayIndex] += ","
		}
	}
	cronSlice[dayIndex] = strings.TrimSuffix(cronSlice[dayIndex], ",") // removing the extra comma at the end
	if cronSlice[dayIndex] == "" {                                     // if done is selected at first, then empty value is replaced by the default *
		cronSlice[dayIndex] = "*"
	}
	return nil
}

func parseDayExpression(itm string) error {
	if itm == everyDay {
		cronSlice[dayIndex] += "*"
	}
	if itm == inputDay {
		var day int
		fmt.Scanf("%d\n", &day)
		if day < 1 || day > 31 {
			return errors.New("day: the day cannot be a less than 1 & cannot be greater than 31")
		}
		cronSlice[dayIndex] += strconv.Itoa(day)
	}
	if itm == defaultStepDay {
		var stepVal int
		fmt.Scanf("%d\n", &stepVal)
		if stepVal < 1 || stepVal > 31 {
			return errors.New("day: the step value cannot be less than 1 or greater than 31")
		}
		cronSlice[dayIndex] += fmt.Sprintf("*/%d", stepVal)
	}
	if itm == stepDay {
		var val, stepVal int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &val)
		if stepVal < 1 || stepVal > 31 {
			return errors.New("day: the step value cannot be less than 1 or greater than 31")
		}
		if val < 1 || val > 31 {
			return errors.New("day: the day cannot be a negative number & cannot be greater than 31")
		}
		cronSlice[dayIndex] += fmt.Sprintf("%d/%d", val, stepVal)
	}
	if itm == rangedDay {
		var from, to int
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if from < 1 || to < 1 || from > to {
			return errors.New("day: the days cannot be negative numbers & the from day must be greater than the to day")
		}
		if from > 31 || to > 31 {
			return errors.New("day: the values cannot be greater than 31")
		}
		cronSlice[dayIndex] += fmt.Sprintf("%d-%d", from, to)
	}

	if itm == rangedstepDay {
		var stepVal, from, to int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if stepVal < 1 || stepVal > 31 {
			return errors.New("day: the step value cannot be less than 1 or greater than 31")
		}
		if from < 1 || to < 1 || from > to {
			return errors.New("day: the days cannot be negative numbers & the from day must be greater than the to day")
		}
		if from > 31 || to > 31 {
			return errors.New("day: the values cannot be greater than 31")
		}
		cronSlice[dayIndex] += fmt.Sprintf("%d-%d/%d", from, to, stepVal)

	}

	return nil
}
