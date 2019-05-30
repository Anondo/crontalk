package generator

import (
	"errors"
	"fmt"
	"github.com/Anondo/crontalk/helper"
	"strconv"
	"strings"
)

const (
	everyMonth       = "every month"
	inputMonth       = "<input>"
	defaultStepMonth = "every <input> months"
	stepMonth        = "every <input> months from <input> to 12"
	rangedMonth      = "<input> to <input> months"
	rangedstepMonth  = "every <input> months from <input> to <input>"
)

var (
	monthItems = []string{
		done,
		everyMonth,
		inputMonth,
		defaultStepMonth,
		stepMonth,
		rangedMonth,
		rangedstepMonth,
	}
)

func runMonth() error {
	var itm string
	var err error
	for itm != done { // iterating until the done option is selected because of list values
		itm, err = helper.RunPrompt("Month", monthItems)
		if err != nil {
			return err
		}
		if err := parseMonthExpression(itm); err != nil {
			fmt.Println(err.Error())
			continue // to avoid adding unwanted ","
		}
		if itm != done { //the check is given because this block will be executed once even after done is selected
			cronSlice[monthIndex] += ","
		}
	}
	cronSlice[monthIndex] = strings.TrimSuffix(cronSlice[monthIndex], ",") // removing the extra comma at the end
	if cronSlice[monthIndex] == "" {                                       // if done is selected at first, then empty value is replaced by the default *
		cronSlice[monthIndex] = "*"
	}
	return nil
}

func parseMonthExpression(itm string) error {
	if itm == everyMonth {
		cronSlice[monthIndex] += "*"
	}
	if itm == inputMonth {
		var month int
		fmt.Scanf("%d\n", &month)
		if month < 1 || month > 12 {
			return errors.New("month: the month cannot be a less than 1 & cannot be greater than 12")
		}
		cronSlice[monthIndex] += strconv.Itoa(month)
	}
	if itm == defaultStepMonth {
		var stepVal int
		fmt.Scanf("%d\n", &stepVal)
		if stepVal < 1 || stepVal > 12 {
			return errors.New("month: the step value cannot be less than 1 or greater than 12")
		}
		cronSlice[monthIndex] += fmt.Sprintf("*/%d", stepVal)
	}
	if itm == stepMonth {
		var val, stepVal int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &val)
		if stepVal < 1 || stepVal > 12 {
			return errors.New("month: the step value cannot be less than 1 or greater than 12")
		}
		if val < 1 || val > 12 {
			return errors.New("month: the month cannot be a negative number & cannot be greater than 12")
		}
		cronSlice[monthIndex] += fmt.Sprintf("%d/%d", val, stepVal)
	}
	if itm == rangedMonth {
		var from, to int
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if from < 1 || to < 1 || from > to {
			return errors.New("month: the months cannot be negative numbers & the from month must be greater than the to month")
		}
		if from > 12 || to > 12 {
			return errors.New("month: the values cannot be greater than 12")
		}
		cronSlice[monthIndex] += fmt.Sprintf("%d-%d", from, to)
	}

	if itm == rangedstepMonth {
		var stepVal, from, to int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if stepVal < 1 || stepVal > 12 {
			return errors.New("month: the step value cannot be less than 1 or greater than 12")
		}
		if from < 1 || to < 1 || from > to {
			return errors.New("month: the months cannot be negative numbers & the from month must be greater than the to month")
		}
		if from > 12 || to > 12 {
			return errors.New("month: the values cannot be greater than 12")
		}
		cronSlice[monthIndex] += fmt.Sprintf("%d-%d/%d", from, to, stepVal)

	}

	return nil
}
