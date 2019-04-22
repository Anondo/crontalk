package translator

import (
	"crontalk/helper"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type translator struct {
	cron          string   //the cron sub-expression
	moment        string   // the moment of the sub-expression
	cronRange     []string //then the two ranged values in string
	ranged        bool     //if the sub-expression is ranged
	listed        bool     //if the sub-expression is listed
	base          bool     // if the occurence is base
	stepped       bool     // if the sub-expression contains step values
	cronListedLen int      //the len of the listed sub-expression
	index         int      // the current index of the listed sub-expression
}

func (t *translator) translateWeekMonth() (bool, error) {
	var v, v1, v2 int
	var mtext string

	if t.moment != week && t.moment != month {
		return false, nil
	}
	if t.stepped {
		t.translateStepValues()
		return true, nil
	}

	if t.ranged {
		var err1, err2 error
		v1, err1 = strconv.Atoi(t.cronRange[0])
		v2, err2 = strconv.Atoi(t.cronRange[1])
		if err1 != nil || err2 != nil {
			return false, err1 //any one will do
		}
	} else {
		var err error
		v, err = strconv.Atoi(t.cron)
		if err != nil {
			return false, err
		}
	}

	if t.base {
		mtext = "every"
	}
	mm := weeks
	if !t.base {
		mtext = "on"
	}
	if t.moment == month {
		mm = months
		if !t.base {
			mtext = "on_month_of"
		}
	}

	//the following is for proper & meaningful sentence
	if t.ranged {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+mtext)+mm[v1]+
			viper.GetString(configStr+"to")+mm[v2], t.index == 0) //if this is the first check
		translatedString += helper.GetStrIfTrue(mm[v1]+viper.GetString(configStr+"to")+mm[v2],
			t.index > 0) //just keep adding the value & not the full sentence

	} else {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+mtext)+mm[v], t.index == 0) //if this is the first check
		translatedString += helper.GetStrIfTrue(mm[v], t.index > 0)                                   //just keep adding the value & not the full sentence
	}
	translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), t.listed && t.index < t.cronListedLen-1) //print "and" for listed
	return true, nil
}

func (t *translator) translateDay() {

	if t.stepped {
		t.translateStepValues()
		return
	}
	mtext := "every"
	if !t.base {
		mtext = "onn"
	}
	if t.ranged {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+mtext)+viper.GetString(configStr+t.moment)+
			t.cronRange[0]+viper.GetString(configStr+"to")+t.cronRange[1], t.index == 0)
		translatedString += helper.GetStrIfTrue(t.cronRange[0]+viper.GetString(configStr+"to")+t.cronRange[1],
			t.index > 0)
	} else {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+mtext)+viper.GetString(configStr+t.moment)+
			t.cron, t.index == 0)
		translatedString += helper.GetStrIfTrue(t.cron, t.index > 0)
	}
	translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), t.listed && t.index < t.cronListedLen-1)
}

func (t *translator) translateMinuteAndHour() error {
	m := cronSlice[minuteIndex]
	h := cronSlice[hourIndex]
	mm, listedM := helper.GetList(m, ",")
	hh, listedH := helper.GetList(h, ",")
	for i, min := range mm { // nested loops are required as, if both the minute & hour values are listed ,the
		for j, hr := range hh { //time is be shown as, for each minute the listed hours
			if strings.Contains(min, "/") {
				t.cron = min
				t.moment = minute
				t.translateStepValues()
			}
			if strings.Contains(hr, "/") {
				t.cron = hr
				t.moment = hour
				t.translateStepValues()
			}
			if strings.Contains(min, "/") || strings.Contains(hr, "/") {
				continue // FIXME: "2 */8 * * *" this expression will not be parsed properly, as getting stepped hour will continue and minute wont be parsed
			}
			mrr, rangedM := helper.GetList(min, "-")
			hrr, rangedH := helper.GetList(hr, "-")
			if rangedM { //if the minute is ranged
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+
					viper.GetString(configStr+"minute")+mrr[0]+viper.GetString(configStr+"to")+mrr[1], i == 0 && j == 0)
				translatedString += helper.GetStrIfTrue(mrr[0]+viper.GetString(configStr+"to")+
					mrr[1], i > 0 || j > 0)
			}
			if rangedH { // if the hour is ranged
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"hour")+hrr[0]+viper.GetString(configStr+"to")+
					hrr[1], i == 0 && j == 0)
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"hour")+hrr[0]+viper.GetString(configStr+"to")+
					hrr[1], i > 0 || j > 0)
			}
			if !rangedM && !rangedH { // if none of them are ranged
				pt, err := helper.PrettyTime(hr, min)
				if err != nil {
					return err
				}
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+pt, i == 0 && j == 0)
				translatedString += helper.GetStrIfTrue(pt, i > 0 || j > 0)
			} else if !rangedH && rangedM { //or if only the minute is ranged
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"hour")+hr, true)
			} else if rangedH && !rangedM { //if only the hour is ranged
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"minute")+min, true)
			}

			translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), (listedM || listedH) &&
				(i < len(mm)-1) || (j < len(hh)-1))
		}
	}
	return nil
}

func (t *translator) translateMinuteOrHour() {
	mStr := moments[minuteIndex] // assuming minute is not default
	mVal := cronSlice[minuteIndex]
	if mVal == anyValue { //if so
		hVal := cronSlice[hourIndex] // working with the hour only
		hh, listed := helper.GetList(hVal, ",")
		for i, hr := range hh { // iterating because could be a list
			hrr, ranged := helper.GetList(hr, "-")
			if ranged { // checking if the value is ranged , different output if so
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at_every_minute_of_hour")+hrr[0]+
					viper.GetString(configStr+"to")+hrr[1], i == 0)
				translatedString += helper.GetStrIfTrue(hrr[0]+viper.GetString(configStr+"to")+hrr[1], i > 0)

			} else {
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at_every_minute_of_hour")+hr, i == 0)
				translatedString += helper.GetStrIfTrue(hr, i > 0)
			}

			translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && i < len(hh)-1)
		}
	} else {
		mm, listed := helper.GetList(mVal, ",") // working with minute only
		for i, min := range mm {                //iterating because could be a list
			mr, ranged := helper.GetList(min, "-")
			if ranged { //checking if the value is ranged
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+viper.GetString(configStr+mStr)+
					" "+mr[0]+viper.GetString(configStr+"to")+mr[1], i == 0)
				translatedString += helper.GetStrIfTrue(mr[0]+viper.GetString(configStr+"to")+mr[1], i > 0)
			} else {
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+viper.GetString(configStr+mStr)+
					" "+min, i == 0)
				translatedString += helper.GetStrIfTrue(min, i > 0)
			}

			translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && i < len(mm)-1)
		}
	}

}

func (t *translator) translateStepValues() {
	steppedCron, _ := helper.GetList(t.cron, "/")
	stepValue := steppedCron[1]
	value := steppedCron[0]
	rValue, ranged := helper.GetList(value, "-")

	translatedString += " " //adding a space for optical optimization

	if ranged { //if ranged, instead of a default stopping range like normal(i.e 2/2) ones, the ranged value will be given
		if t.moment == week {
			i1, _ := strconv.Atoi(rValue[0])
			i2, _ := strconv.Atoi(rValue[1])
			translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"day_of_the_week") +
				viper.GetString(configStr+"from") + weeks[i1] + viper.GetString(configStr+"to") + weeks[i2]
		}
		if t.moment == month {
			i1, _ := strconv.Atoi(rValue[0])
			i2, _ := strconv.Atoi(rValue[1])
			translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"month_of_the_year") +
				viper.GetString(configStr+"from") + months[i1] + viper.GetString(configStr+"to") + months[i2]
		}
		if t.moment == day {
			translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"day_of_the_month") +
				viper.GetString(configStr+"from") + rValue[0] + viper.GetString(configStr+"to") + rValue[1]
		}
		if t.moment == hour {
			translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"hour") +
				viper.GetString(configStr+"from") + rValue[0] + viper.GetString(configStr+"to") + rValue[1]
		}
		if t.moment == minute {
			translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"minute") +
				viper.GetString(configStr+"from") + rValue[0] + viper.GetString(configStr+"to") + rValue[1]
		}
	} else { //if not ranged
		if value == anyValue { //if */<step-value>
			if t.moment == week {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"day_of_the_week") +
					" "
			}
			if t.moment == month {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"month_of_the_year") +
					" "
			}
			if t.moment == day {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"day_of_the_month")
			}
			if t.moment == hour {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"hour")
			}
			if t.moment == minute {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"minute")
			}
		} else { //if example: 5/<step-value>
			if t.moment == week {
				i, _ := strconv.Atoi(value)
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"day_of_the_week") +
					viper.GetString(configStr+"from") + weeks[i] + viper.GetString(configStr+"to") +
					viper.GetString(configStr+"sunday")
			}
			if t.moment == month {
				i, _ := strconv.Atoi(value)
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"month_of_the_year") +
					viper.GetString(configStr+"from") + months[i] + viper.GetString(configStr+"to") +
					viper.GetString(configStr+"december")
			}
			if t.moment == day {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"day_of_the_month") +
					viper.GetString(configStr+"from") + value + viper.GetString(configStr+"to") + "31"
			}
			if t.moment == hour {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"hour") +
					viper.GetString(configStr+"from") + value + viper.GetString(configStr+"to") + "23"
			}
			if t.moment == minute {
				translatedString += viper.GetString(configStr+"every") + stepValue + viper.GetString(configStr+"minute") +
					viper.GetString(configStr+"from") + value + viper.GetString(configStr+"to") + "59"
			}

		}
	}

	translatedString += " , "

}
