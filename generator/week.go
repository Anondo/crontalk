package generator

import (
	"crontalk/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	everyDayOfWeek       = "every day of the week"
	inputDayOfWeek       = "<input>"
	defaultStepDayOfWeek = "every <input> days of the week"
	stepDayOfWeek        = "every <input> days of the week from <input> to 6"
	rangedDayOfWeek      = "<input> to <input> days of the week"
	rangedstepDayOfWeek  = "every <input> days of the week from <input> to <input>"
)

var (
	weekItems = []string{
		done,
		everyDayOfWeek,
		inputDayOfWeek,
		defaultStepDayOfWeek,
		stepDayOfWeek,
		rangedDayOfWeek,
		rangedstepDayOfWeek,
	}
)

func runWeek() error {
	var itm string
	var err error
	for itm != done { // iterating until the done option is selected because of list values
		itm, err = helper.RunPrompt("Week", weekItems)
		if err != nil {
			return err
		}
		if err := parseWeekExpression(itm); err != nil {
			fmt.Println(err.Error())
			continue // to avoid adding unwanted ","
		}
		if itm != done { //the check is given because this block will be executed once even after done is selected
			cronSlice[weekIndex] += ","
		}
	}
	cronSlice[weekIndex] = strings.TrimSuffix(cronSlice[weekIndex], ",") // removing the extra comma at the end
	if cronSlice[weekIndex] == "" {                                      // if done is selected at first, then empty value is replaced by the default *
		cronSlice[weekIndex] = "*"
	}
	return nil
}

func parseWeekExpression(itm string) error {
	if itm == everyDayOfWeek {
		cronSlice[weekIndex] += "*"
	}
	if itm == inputDayOfWeek {
		var week int
		fmt.Scanf("%d\n", &week)
		if week < 0 || week > 6 {
			return errors.New("day of the week: the day of the week cannot be a less than 0 & cannot be greater than 6")
		}
		cronSlice[weekIndex] += strconv.Itoa(week)
	}
	if itm == defaultStepDayOfWeek {
		var stepVal int
		fmt.Scanf("%d\n", &stepVal)
		if stepVal < 1 || stepVal > 6 {
			return errors.New("day of the week: the step value cannot be less than 1 or greater than 6")
		}
		cronSlice[weekIndex] += fmt.Sprintf("*/%d", stepVal)
	}
	if itm == stepDayOfWeek {
		var val, stepVal int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &val)
		if stepVal < 1 || stepVal > 6 {
			return errors.New("day of the week: the step value cannot be less than 1 or greater than 6")
		}
		if val < 0 || val > 6 {
			return errors.New("day of the week: the day of the week cannot be a negative number & cannot be greater than 6")
		}
		cronSlice[weekIndex] += fmt.Sprintf("%d/%d", val, stepVal)
	}
	if itm == rangedDayOfWeek {
		var from, to int
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if from < 0 || to < 0 || from > to {
			return errors.New("day of the week: the days of the week cannot be negative numbers & the from day of the week must be greater than the to day of the week")
		}
		if from > 6 || to > 6 {
			return errors.New("day of the week: the values cannot be greater than 6")
		}
		cronSlice[weekIndex] += fmt.Sprintf("%d-%d", from, to)
	}

	if itm == rangedstepDayOfWeek {
		var stepVal, from, to int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if stepVal < 1 || stepVal > 6 {
			return errors.New("day of the week: the step value cannot be less than 1 or greater than 6")
		}
		if from < 0 || to < 0 || from > to {
			return errors.New("day of the week: the days of the week cannot be negative numbers & the from day of the week must be greater than the to day of the week")
		}
		if from > 6 || to > 6 {
			return errors.New("day of the week: the values cannot be greater than 6")
		}
		cronSlice[weekIndex] += fmt.Sprintf("%d-%d/%d", from, to, stepVal)

	}

	return nil
}
