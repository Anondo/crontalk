package generator

import (
	"crontalk/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	every       = "every"
	input       = "<input>"
	defaultStep = "every <input> minutes"
	step        = "every <input> minutes from <input> to 59"
	ranged      = "<input> to <input> minutes"
	rangedStep  = "every <input> minutes from <input> to <input>"
)

var (
	minItems = []string{
		every,
		input,
		defaultStep,
		step,
		ranged,
		rangedStep,
		done,
	}
)

func runMinute() error {
	cronSlice[minIndex] = ""
	var itm string
	var err error
	for itm != done {
		itm, err = helper.RunPrompt("Minute", minItems)
		if err != nil {
			return err
		}
		if err := parseMinuteExpression(itm); err != nil {
			fmt.Println(err.Error())
		}
		cronSlice[minIndex] += ","
	}
	cronSlice[minIndex] = strings.TrimSuffix(cronSlice[minIndex], ",")
	return nil
}

func parseMinuteExpression(itm string) error {
	if itm == every {
		cronSlice[minIndex] += "*"
	}
	if itm == input {
		var min int
		fmt.Scanf("%d\n", &min)
		if min < 0 {
			return errors.New("minute: the minute cannot be a negative number")
		}
		cronSlice[minIndex] += strconv.Itoa(min)
	}
	if itm == defaultStep {
		var stepVal int
		fmt.Scanf("%d\n", &stepVal)
		if stepVal < 1 {
			return errors.New("minute: the step value cannot be a negative or 0 number")
		}
		cronSlice[minIndex] += fmt.Sprintf("*/%d", stepVal)
	}
	if itm == step {
		var val, stepVal int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &val)
		if stepVal < 1 {
			return errors.New("minute: the step value cannot be a negative or 0 number")
		}
		if val < 0 {
			return errors.New("minute: the minute cannot be a negative number")
		}
		cronSlice[minIndex] += fmt.Sprintf("%d/%d", val, stepVal)
	}
	if itm == ranged {
		var from, to int
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if from < 0 || to < 0 || from > to {
			return errors.New("the minutes cannot be negative numbers & the from minute must be greater than the to minute")
		}
		cronSlice[minIndex] += fmt.Sprintf("%d-%d", from, to)
	}

	if itm == rangedStep {
		var stepVal, from, to int
		fmt.Scanf("%d\n", &stepVal)
		fmt.Scanf("%d\n", &from)
		fmt.Scanf("%d\n", &to)
		if stepVal < 1 {
			return errors.New("minute: the step value cannot be a negative or 0 number")
		}
		if from < 0 || to < 0 || from > to {
			return errors.New("the minutes cannot be negative numbers & the from minute must be greater than the to minute")
		}
		cronSlice[minIndex] += fmt.Sprintf("%d-%d/%d", from, to, stepVal)

	}

	return nil
}
