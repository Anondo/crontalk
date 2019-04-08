package translator

import (
	"crontalk/helper"
	"strconv"

	"github.com/spf13/viper"
)

const (
	english = "english"
	bangla  = "bangla"
)

var (
	weeks     = map[int]string{}
	months    = map[int]string{}
	baseIndex int
	configStr = ""
	language  = english
)

// Init initializes the translator
func Init() {
	if viper.GetBool(bangla) {
		language = bangla
		moments[dayIndex] = "দিন:" //not taking from config for the sake of simplicity
		moments[minuteIndex] = "মিনিট:"
	}
	configStr = "language." + language + "." //the config index to parse the config yaml file from viper

	weeks = map[int]string{
		0: viper.GetString(configStr + "sunday"),
		1: viper.GetString(configStr + "monday"),
		2: viper.GetString(configStr + "tuesday"),
		3: viper.GetString(configStr + "wednesday"),
		4: viper.GetString(configStr + "thursday"),
		5: viper.GetString(configStr + "friday"),
		6: viper.GetString(configStr + "saturday"),
	}
	months = map[int]string{
		1:  viper.GetString(configStr + "january"),
		2:  viper.GetString(configStr + "february"),
		3:  viper.GetString(configStr + "march"),
		4:  viper.GetString(configStr + "april"),
		5:  viper.GetString(configStr + "may"),
		6:  viper.GetString(configStr + "june"),
		7:  viper.GetString(configStr + "july"),
		8:  viper.GetString(configStr + "august"),
		9:  viper.GetString(configStr + "september"),
		10: viper.GetString(configStr + "october"),
		11: viper.GetString(configStr + "november"),
		12: viper.GetString(configStr + "december"),
	}
}

func translateBaseOccurence() error {
	var i int
	for i = weekIndex; i > hourIndex; i-- { //start iterating from the last sub-expressions to determine the starting string
		if cronSlice[i] != anyValue {
			if moments[i] == week { // checking for weekly
				wi, err := strconv.Atoi(cronSlice[i])
				if err != nil {
					return err
				}
				translatedString = viper.GetString(configStr+"every") + weeks[wi]
				break
			}
			if moments[i] == month { // checking for monthly, in seperate if blocks because , slight change of translated string
				mi, err := strconv.Atoi(cronSlice[i])
				if err != nil {
					return err
				}
				translatedString = viper.GetString(configStr+"every_month_of") + months[mi]
				break
			}
			translatedString = viper.GetString(configStr+"every") + moments[i] + " " + cronSlice[i] //checking for the day
			break
		}
	}
	if i == hourIndex { // checking if every sub-expression contains asteriks apart from the time part
		translatedString += viper.GetString(configStr + "every_day")
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
				translatedString += viper.GetString(configStr+"on") + weeks[wi]
			} else if moments[i] == month {
				mi, err := strconv.Atoi(cronSlice[i])
				if err != nil {
					return err
				}
				translatedString += viper.GetString(configStr+"on_month_of") + months[mi]
			} else {
				translatedString += viper.GetString(configStr+"onn") + moments[i] + " " + cronSlice[i] //no breaks like base and a bit different string
			}
		}
	}
	return nil
}

func translateTimeOccurence() error {
	if cronSlice[minuteIndex] == anyValue && cronSlice[hourIndex] == anyValue { // checking if both hour and minute are defaults
		translatedString += viper.GetString(configStr + "at_every_minute")
	} else if cronSlice[minuteIndex] != anyValue && cronSlice[hourIndex] != anyValue { //checking if non of them are
		m := cronSlice[minuteIndex]
		h := cronSlice[hourIndex]
		pt, err := helper.PrettyTime(h, m)
		if err != nil {
			return err
		}
		translatedString += viper.GetString(configStr+"at") + pt
	} else { // checking if  just one of them is default
		mStr := moments[minuteIndex] // assuming minute is not default
		mVal := cronSlice[minuteIndex]
		if mVal == anyValue { //if so
			hVal := cronSlice[hourIndex]
			translatedString += viper.GetString(configStr+"at_every_minute_of_hour") + hVal
		} else {
			translatedString += viper.GetString(configStr+"at") + mStr + " " + mVal
		}
	}
	return nil
}
