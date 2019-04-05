package translator

import (
	"crontalk/helper"
	"strconv"
)

var (
	weeks = map[int]string{
		0: "sunday",
		1: "monday",
		2: "tuesday",
		3: "wednesday",
		4: "thursday",
		5: "friday",
		6: "saturday",
	}
	months = map[int]string{
		1:  "january",
		2:  "february",
		3:  "march",
		4:  "april",
		5:  "may",
		6:  "june",
		7:  "july",
		8:  "august",
		9:  "september",
		10: "october",
		11: "november",
		12: "december",
	}
	baseIndex int
)

func translateBaseOccurence() error {
	var i int
	for i = weekIndex; i > hourIndex; i-- { //start iterating from the last sub-expressions to determine the starting string
		if cronSlice[i] != anyValue {
			if moments[i] == week { // checking for weekly
				wi, err := strconv.Atoi(cronSlice[i])
				if err != nil {
					return err
				}
				translatedString = "Every " + weeks[wi]
				break
			}
			if moments[i] == month { // checking for monthly, in seperate if blocks because , slight change of translated string
				mi, err := strconv.Atoi(cronSlice[i])
				if err != nil {
					return err
				}
				translatedString = "Every month of " + months[mi]
				break
			}
			translatedString = "Every " + moments[i] + " " + cronSlice[i] //checking for the day
			break
		}
	}
	if i == hourIndex { // checking if every sub-expression contains asteriks apart from the time part
		translatedString += " Every Day"
	}
	baseIndex = i //storing the base index so that when checking every other than time , the base is also omitted because its
	//already checked

	return nil

}

func translateAllButBaseTimeOccurence() error {
	for i := dayIndex; i <= weekIndex; i++ { // checking every other sub-expressions apart from the base and time, no need for reverse travel
		if cronSlice[i] != anyValue && i != baseIndex { // no gonna check the base
			if moments[i] == week {
				wi, err := strconv.Atoi(cronSlice[i])
				if err != nil {
					return err
				}
				translatedString += " on " + weeks[wi]
			} else if moments[i] == month {
				mi, err := strconv.Atoi(cronSlice[i])
				if err != nil {
					return err
				}
				translatedString += " on month of " + months[mi]
			} else {
				translatedString += " on " + moments[i] + " " + cronSlice[i] //no breaks like base and a bit different string
			}
		}
	}
	return nil
}

func translateTimeOccurence() error {
	if cronSlice[minuteIndex] == anyValue && cronSlice[hourIndex] == anyValue { // checking if both hour and minute are defaults
		translatedString += " at every minute"
	} else if cronSlice[minuteIndex] != anyValue && cronSlice[hourIndex] != anyValue { //checking if non of them are
		m := cronSlice[minuteIndex]
		h := cronSlice[hourIndex]
		pt, err := helper.PrettyTime(h, m)
		if err != nil {
			return err
		}
		translatedString += " at " + pt
	} else { // checking if  just one of them is default
		mStr := moments[minuteIndex] // assuming minute is not default
		mVal := cronSlice[minuteIndex]
		if mVal == anyValue { //if so
			hVal := cronSlice[hourIndex]
			translatedString += " at every minute of hour " + hVal
		} else {
			translatedString += " at " + mStr + " " + mVal
		}
	}
	return nil
}
