package generator

import (
	"github.com/Anondo/crontalk/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	everyMin       = "every minute"
	inputMin       = "<input>"
	defaultStepMin = "every <input> minutes"
	stepMin        = "every <input> minutes from <input> to 59"
	rangedMin      = "<input> to <input> minutes"
	rangedstepMin  = "every <input> minutes from <input> to <input>"
)

var (
	minItems = []string{
		done,
		everyMin,
		inputMin,
		defaultStepMin,
		stepMin,
		rangedMin,
		rangedstepMin,
	}
)

func runMinute() error {
	var itm string
	var err error
	for itm != done { // iterating until the done option is selected because of list values
		itm, err = helper.RunPrompt("Minute", minItems)
		if err != nil {
			return err
		}
		if err := parseMinuteExpression(itm); err != nil {
			fmt.Println(err.Error())
			continue // to avoid adding unwanted ","
		}
		if itm != done { //the check is given because this block will be executed once even after done is selected
			cronSlice[minIndex] += ","
		}
	}
	cronSlice[minIndex] = strings.TrimSuffix(cronSlice[minIndex], ",") // removing the extra comma at the end
	if cronSlice[minIndex] == "" {                                     // if done is selected at first, then empty value is replaced by the default *
		cronSlice[minIndex] = "*"
	}
	return nil
}

func parseMinuteExpression(itm string) error {
	if itm == everyMin {
		cronSlice[minIndex] += "*"
	}
	if itm == inputMin {
		var min int
		fmt.Scanf("%d\n", &min)
		if min < 0 || min > 59 {
			return errors.New("minute: the minute cannot be a negative number & cannot be greater than 59")
		}
		cronSlice[minIndex] += strconv.Itoa(min)
	}
	if itm == defaultStepMin {
		var stepVal int
		fmt.Scanf("%d\n", &stepVal)
		if stepVal < 1 || stepVal > 59 {
			return errors.New("minute: the step value cannot be a negative or 0 number or greater than 59")
		}
		cronSlice[minIndex] += fmt.Sprintf("*/%d", stepVal)
	}
	if itm == stepMin {
		var val, stepVal int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &val)
		if stepVal < 1 || stepVal > 59 {
			return errors.New("minute: the step value cannot be a negative or 0 number or greater than 59")
		}
		if val < 0 || val > 59 {
			return errors.New("minute: the minute cannot be a negative number & cannot be greater than 59")
		}
		cronSlice[minIndex] += fmt.Sprintf("%d/%d", val, stepVal)
	}
	if itm == rangedMin {
		var from, to int
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if from < 0 || to < 0 || from > to {
			return errors.New("minute: the minutes cannot be negative numbers & the from minute must be greater than the to minute")
		}
		if from > 59 || to > 59 {
			return errors.New("minute: the values cannot be greater than 59")
		}
		cronSlice[minIndex] += fmt.Sprintf("%d-%d", from, to)
	}

	if itm == rangedstepMin {
		var stepVal, from, to int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if stepVal < 1 || stepVal > 59 {
			return errors.New("minute: the step value cannot be a negative or 0 number or greater than 59")
		}
		if from < 0 || to < 0 || from > to {
			return errors.New("the minutes cannot be negative numbers & the from minute must be greater than the to minute")
		}
		if from > 59 || to > 59 {
			return errors.New("minute: the values cannot be greater than 59")
		}
		cronSlice[minIndex] += fmt.Sprintf("%d-%d/%d", from, to, stepVal)

	}

	return nil
}
