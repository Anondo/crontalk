package translator

import (
	"strings"

	"github.com/Anondo/crontalk/config"
	"github.com/Anondo/crontalk/helper"

	"github.com/spf13/viper"
)

var (
	weeks     = map[int]string{}
	months    = map[int]string{}
	baseIndex int
	cfg       = config.LangConfig{}
)

// Init initializes the translator
func Init() {
	language := viper.GetString("lang")
	cfg = config.Language(language)

	weeks = map[int]string{
		0: cfg.Sunday,
		1: cfg.Monday,
		2: cfg.Tuesday,
		3: cfg.Wednesday,
		4: cfg.Thursday,
		5: cfg.Friday,
		6: cfg.Saturday,
	}
	months = map[int]string{
		1:  cfg.January,
		2:  cfg.February,
		3:  cfg.March,
		4:  cfg.April,
		5:  cfg.May,
		6:  cfg.June,
		7:  cfg.July,
		8:  cfg.August,
		9:  cfg.September,
		10: cfg.October,
		11: cfg.November,
		12: cfg.December,
	}
}

func translateBaseOccurence() error {
	var i int
	for i = weekIndex; i > hourIndex; i-- { //start iterating from the last sub-expressions to determine the starting string
		if cronSlice[i] != every {
			cc, listed := helper.GetList(cronSlice[i], list)
			for j, c := range cc { //iterating because values can be listed
				validWordParse(&c, moments[i])
				rr, ranged := helper.GetList(c, rangee)
				if ranged && (moments[i] == week || moments[i] == month) {
					validWordParse(&rr[0], moments[i])
					validWordParse(&rr[1], moments[i])
				}
				t := translator{
					cron:          c,
					moment:        moments[i],
					cronRange:     rr,
					ranged:        ranged,
					listed:        listed,
					stepped:       strings.Contains(c, step),
					base:          true,
					cronListedLen: len(cc),
					index:         j,
				}
				if found, err := t.translateWeekMonth(); err != nil {
					return err
				} else if found {
					continue
				}
				t.translateDay()
			}
			break //once the base value is found no need for further iterations
		}
	}
	if i == hourIndex { // checking if every sub-expression contains asteriks apart from the time part
		translatedString += cfg.EveryDay
	}
	baseIndex = i //storing the base index so that when checking every other than time , the base is also omitted because its
	//already checked
	return nil

}

func translateAllButBaseTimeOccurence() error {

	for i := dayIndex; i <= weekIndex; i++ { //checking every other sub-expressions apart from the base and time, no need for reverse travel
		if cronSlice[i] != every && i != baseIndex { //not gonna check the base
			cc, listed := helper.GetList(cronSlice[i], list)
			for j, c := range cc { //iterating the single sub-expressions
				validWordParse(&c, moments[i])
				rr, ranged := helper.GetList(c, rangee)
				if ranged && (moments[i] == week || moments[i] == month) {
					validWordParse(&rr[0], moments[i])
					validWordParse(&rr[1], moments[i])
				}
				t := translator{
					cron:          c,
					moment:        moments[i],
					cronRange:     rr,
					ranged:        ranged,
					listed:        listed,
					stepped:       strings.Contains(c, step),
					base:          false,
					cronListedLen: len(cc),
					index:         j,
				}
				if found, err := t.translateWeekMonth(); err != nil {
					return err
				} else if !found {
					t.translateDay()
				}
			}
		}
	}
	return nil
}

func translateTimeOccurence() error {
	var t translator
	if cronSlice[minuteIndex] == every && cronSlice[hourIndex] == every { // checking if both hour and minute are defaults
		translatedString += cfg.AtEveryMinute
	} else if cronSlice[minuteIndex] != every && cronSlice[hourIndex] != every { //checking if non of them are
		if err := t.translateMinuteAndHour(); err != nil {
			return err
		}
	} else { // checking if  just one of them is default
		t.translateMinuteOrHour()
	}
	return nil
}

func translateDigits() {

	if viper.GetString("lang") == helper.LanguageEnglish {
		return
	}

	for _, c := range translatedString {
		cs := string(c)
		if helper.IsDigit(cs) {
			char := ""

			switch cs {
			case "0":
				char = cfg.Number0
			case "1":
				char = cfg.Number1
			case "2":
				char = cfg.Number2
			case "3":
				char = cfg.Number3
			case "4":
				char = cfg.Number4
			case "5":
				char = cfg.Number5
			case "6":
				char = cfg.Number6
			case "7":
				char = cfg.Number7
			case "8":
				char = cfg.Number8
			case "9":
				char = cfg.Number9
			}

			if char != "" {
				translatedString = strings.Replace(translatedString, cs, char, -1)
			}
		}
	}
}
