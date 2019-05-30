package generator

import (
	"github.com/Anondo/crontalk/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	everyHr       = "every hour"
	inputHr       = "<input>"
	defaultStepHr = "every <input> hours"
	stepHr        = "every <input> hours from <input> to 59"
	rangedHr      = "<input> to <input> hours"
	rangedstepHr  = "every <input> hours from <input> to <input>"
)

var (
	hourItems = []string{
		done,
		everyHr,
		inputHr,
		defaultStepHr,
		stepHr,
		rangedHr,
		rangedstepHr,
	}
)

func runHour() error {
	var itm string
	var err error
	for itm != done { // iterating until the done option is selected because of list values
		itm, err = helper.RunPrompt("Hour", hourItems)
		if err != nil {
			return err
		}
		if err := parseHourExpression(itm); err != nil {
			fmt.Println(err.Error())
			continue // to avoid adding unwanted ","
		}
		if itm != done { //the check is given because this block will be executed once even after done is selected
			cronSlice[hourIndex] += ","
		}
	}
	cronSlice[hourIndex] = strings.TrimSuffix(cronSlice[hourIndex], ",") // removing the extra comma at the end
	if cronSlice[hourIndex] == "" {                                      // if done is selected at first, then empty value is replaced by the default *
		cronSlice[hourIndex] = "*"
	}
	return nil
}

func parseHourExpression(itm string) error {
	if itm == everyHr {
		cronSlice[hourIndex] += "*"
	}
	if itm == inputHr {
		var hour int
		fmt.Scanf("%d\n", &hour)
		if hour < 0 || hour > 59 {
			return errors.New("hour: the hour cannot be a negative number & cannot be greater than 59")
		}
		cronSlice[hourIndex] += strconv.Itoa(hour)
	}
	if itm == defaultStepHr {
		var stepVal int
		fmt.Scanf("%d\n", &stepVal)
		if stepVal < 1 || stepVal > 59 {
			return errors.New("hour: the step value cannot be less than 1 or greater than 59")
		}
		cronSlice[hourIndex] += fmt.Sprintf("*/%d", stepVal)
	}
	if itm == stepHr {
		var val, stepVal int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &val)
		if stepVal < 1 || stepVal > 59 {
			return errors.New("hour: the step value cannot be less then 1 or greater than 59")
		}
		if val < 0 || val > 59 {
			return errors.New("hour: the hour cannot be a negative number & cannot be greater than 59")
		}
		cronSlice[hourIndex] += fmt.Sprintf("%d/%d", val, stepVal)
	}
	if itm == rangedHr {
		var from, to int
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if from < 0 || to < 0 || from > to {
			return errors.New("hour: the hours cannot be negative numbers & the from hour must be greater than the to hour")
		}
		if from > 59 || to > 59 {
			return errors.New("hour: the values cannot be greater than 59")
		}
		cronSlice[hourIndex] += fmt.Sprintf("%d-%d", from, to)
	}

	if itm == rangedstepHr {
		var stepVal, from, to int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if stepVal < 1 || stepVal > 59 {
			return errors.New("hour: the step value cannot be less then 1 or greater than 59")
		}
		if from < 0 || to < 0 || from > to {
			return errors.New("hour: the hours cannot be negative numbers & the from hour must be greater than the to hour")
		}
		if from > 59 || to > 59 {
			return errors.New("hour: the values cannot be greater than 59")
		}
		cronSlice[hourIndex] += fmt.Sprintf("%d-%d/%d", from, to, stepVal)

	}

	return nil
}
