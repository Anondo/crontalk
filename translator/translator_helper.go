package translator

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Anondo/crontalk/helper"
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
		return true, t.translateStepValues()

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
		mtext = cfg.Every
	}
	mm := weeks
	if !t.base {
		mtext = cfg.Onn
	}
	if t.moment == month {
		mm = months
		if !t.base {
			mtext = cfg.OnMonthOf
		}
	}

	//the following is for proper & meaningful sentence
	if t.ranged {
		translatedString += helper.GetStrIfTrue(mtext+mm[v1]+
			cfg.To+mm[v2], t.index == 0) //if this is the first check
		translatedString += helper.GetStrIfTrue(mm[v1]+cfg.To+mm[v2],
			t.index > 0) //just keep adding the value & not the full sentence

	} else {
		translatedString += helper.GetStrIfTrue(mtext+mm[v], t.index == 0) //if this is the first check
		translatedString += helper.GetStrIfTrue(mm[v], t.index > 0)        //just keep adding the value & not the full sentence
	}
	translatedString += helper.GetStrIfTrue(cfg.And, t.listed && t.index < t.cronListedLen-1) //print "and" for listed
	return true, nil
}

func (t *translator) translateDay() {

	if t.stepped {
		t.translateStepValues()
		return
	}
	mtext := cfg.Every //setting the starting text for translation
	if !t.base {
		mtext = cfg.Onn
	}
	if t.ranged {
		translatedString += helper.GetStrIfTrue(mtext+cfg.ToMap()[t.moment]+
			t.cronRange[0]+cfg.To+t.cronRange[1], t.index == 0)
		translatedString += helper.GetStrIfTrue(t.cronRange[0]+cfg.To+t.cronRange[1],
			t.index > 0)
	} else {
		translatedString += helper.GetStrIfTrue(mtext+cfg.ToMap()[t.moment]+
			t.cron, t.index == 0)
		translatedString += helper.GetStrIfTrue(t.cron, t.index > 0)
	}
	translatedString += helper.GetStrIfTrue(cfg.And, t.listed && t.index < t.cronListedLen-1)
}

func (t *translator) translateMinuteAndHour() error {
	m := cronSlice[minuteIndex]
	h := cronSlice[hourIndex]
	mm, listedM := helper.GetList(m, list)
	hh, listedH := helper.GetList(h, list)
	for i, min := range mm { // nested loops are required as, if both the minute & hour values are listed ,the
		for j, hr := range hh { //time is be shown as, for each minute the listed hours
			steppedHour := false
			steppedMinute := false
			if strings.Contains(min, step) { // if the minute is stepped
				t.cron = min
				t.moment = minute
				t.translateStepValues()
				steppedMinute = true
			}
			if strings.Contains(hr, step) { // if the hour is stepped
				t.cron = hr
				t.moment = hour
				t.translateStepValues()
				steppedHour = true
			}
			_, rangedM := helper.GetList(min, rangee)
			_, rangedH := helper.GetList(hr, rangee)

			if !rangedM && !rangedH && !steppedHour && !steppedMinute { // if none of them are ranged or stepped then print like a normal time
				pt, err := helper.PrettyTime(hr, min)
				if err != nil {
					return err
				}
				translatedString += helper.GetStrIfTrue(cfg.At+pt, i == 0 && j == 0)
				translatedString += helper.GetStrIfTrue(pt, i > 0 || j > 0)
			} else if (!rangedH && !steppedHour) && (rangedM || steppedMinute) { //or if only the minute is ranged or stepped
				translatedString += helper.GetStrIfTrue(cfg.Hour+hr, true)
			} else if (rangedH || steppedHour) && (!rangedM && !steppedMinute) { //if only the hour is ranged or stepped
				translatedString += helper.GetStrIfTrue(cfg.Minute+min, true)
			}

			translatedString += helper.GetStrIfTrue(cfg.And, (listedM || listedH) &&
				(i < len(mm)-1) || (j < len(hh)-1))
		}
	}
	return nil
}

func (t *translator) translateMinuteOrHour() {
	mStr := moments[minuteIndex] // assuming minute is not default
	mVal := cronSlice[minuteIndex]
	if mVal == every { //if default
		hVal := cronSlice[hourIndex] // working with the hour only
		hh, listed := helper.GetList(hVal, list)
		for i, hr := range hh { // iterating because could be a list
			if strings.Contains(hr, step) { // if the hour is stepped
				t.cron = hr
				t.moment = hour
				translatedString += cfg.AtEveryMinute // as the minute is * and minute will not be parsed later in this function
				t.translateStepValues()
				continue
			}
			hrr, ranged := helper.GetList(hr, rangee)
			if ranged { // checking if the value is ranged , different output if so
				translatedString += helper.GetStrIfTrue(cfg.AtEveryMinuteOfHour+hrr[0]+
					cfg.To+hrr[1], i == 0)
				translatedString += helper.GetStrIfTrue(hrr[0]+cfg.To+hrr[1], i > 0)

			} else {
				translatedString += helper.GetStrIfTrue(cfg.AtEveryMinuteOfHour+hr, i == 0)
				translatedString += helper.GetStrIfTrue(hr, i > 0)
			}

			translatedString += helper.GetStrIfTrue(cfg.And, listed && i < len(hh)-1)
		}
	} else {
		mm, listed := helper.GetList(mVal, list) // working with minute only
		for i, min := range mm {                 //iterating because could be a list
			if strings.Contains(min, step) { // if the minute is stepped
				t.cron = min
				t.moment = minute
				t.translateStepValues()
				continue
			}
			mr, ranged := helper.GetList(min, rangee)
			if ranged { //checking if the value is ranged
				translatedString += helper.GetStrIfTrue(cfg.At+cfg.ToMap()[mStr]+
					" "+mr[0]+cfg.To+mr[1], i == 0)
				translatedString += helper.GetStrIfTrue(mr[0]+cfg.To+mr[1], i > 0)
			} else {
				translatedString += helper.GetStrIfTrue(cfg.At+cfg.ToMap()[mStr]+
					" "+min, i == 0)
				translatedString += helper.GetStrIfTrue(min, i > 0)
			}

			translatedString += helper.GetStrIfTrue(cfg.And, listed && i < len(mm)-1)
		}
	}

}

func (t *translator) translateStepValues() error {
	steppedCron, _ := helper.GetList(t.cron, step)
	stepValue := steppedCron[1]
	value := steppedCron[0]
	rValue, ranged := helper.GetList(value, rangee)

	if !ranged && value != every && !validWordParse(&value, t.moment) {
		return errors.New("invalid value word")
	}

	translatedString += " " //adding a space for optical optimization

	if ranged { //if ranged, instead of a default stopping range like normal(i.e 2/2) ones, the ranged value will be given
		if t.moment == week {
			if !validWordParse(&rValue[0], week) || !validWordParse(&rValue[1], week) {
				return errors.New("Invalid ranged step value word")
			}
			i1, _ := strconv.Atoi(rValue[0])
			i2, _ := strconv.Atoi(rValue[1])
			translatedString += cfg.Every + stepValue + cfg.DayOfTheWeek +
				cfg.From + weeks[i1] + cfg.To + weeks[i2]
		}
		if t.moment == month {
			if !validWordParse(&rValue[0], month) || !validWordParse(&rValue[1], month) {
				return errors.New("Invalid ranged step value word")
			}
			i1, _ := strconv.Atoi(rValue[0])
			i2, _ := strconv.Atoi(rValue[1])
			translatedString += cfg.Every + stepValue + cfg.MonthOfTheYear +
				cfg.From + months[i1] + cfg.To + months[i2]
		}
		if t.moment == day {
			translatedString += cfg.Every + stepValue + cfg.DayOfTheMonth +
				cfg.From + rValue[0] + cfg.To + rValue[1]
		}
		if t.moment == hour {
			translatedString += cfg.Every + stepValue + cfg.Hour +
				cfg.From + rValue[0] + cfg.To + rValue[1]
		}
		if t.moment == minute {
			translatedString += cfg.Every + stepValue + cfg.Minute +
				cfg.From + rValue[0] + cfg.To + rValue[1]
		}
	} else { //if not ranged
		if value == every { //if */<step-value>
			if t.moment == week {
				translatedString += cfg.Every + stepValue + cfg.DayOfTheWeek +
					" "
			}
			if t.moment == month {
				translatedString += cfg.Every + stepValue + cfg.MonthOfTheYear +
					" "
			}
			if t.moment == day {
				translatedString += cfg.Every + stepValue + cfg.DayOfTheMonth
			}
			if t.moment == hour {
				translatedString += cfg.Every + stepValue + cfg.Hour
			}
			if t.moment == minute {
				translatedString += cfg.Every + stepValue + cfg.Minute
			}
		} else { //if example: 5/<step-value>
			if t.moment == week {
				i, _ := strconv.Atoi(value)
				translatedString += cfg.Every + stepValue + cfg.DayOfTheWeek +
					cfg.From + weeks[i] + cfg.To +
					cfg.Sunday
			}
			if t.moment == month {
				i, _ := strconv.Atoi(value)
				translatedString += cfg.Every + stepValue + cfg.MonthOfTheYear +
					cfg.From + months[i] + cfg.To +
					cfg.December
			}
			if t.moment == day {
				translatedString += cfg.Every + stepValue + cfg.DayOfTheMonth +
					cfg.From + value + cfg.To + helper.LastDay
			}
			if t.moment == hour {
				translatedString += cfg.Every + stepValue + cfg.Hour +
					cfg.From + value + cfg.To + helper.LastHour
			}
			if t.moment == minute {
				translatedString += cfg.Every + stepValue + cfg.Minute +
					cfg.From + value + cfg.To + helper.LastMinute
			}

		}
	}

	if t.moment != minute && t.moment != hour {
		translatedString += " , "
	}

	return nil

}
