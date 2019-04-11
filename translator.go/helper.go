package translator

import (
	"crontalk/helper"
	"strconv"

	"github.com/spf13/viper"
)

type translator struct {
	cron          string   //the cron sub-expression
	moment        string   // the moment of the sub-expression
	cronRange     []string //then the two ranged values in string
	ranged        bool     //if the sub-expression is ranged
	listed        bool     //if the sub-expression is listed
	base          bool     // if the occurence is base
	cronListedLen int      //the len of the listed sub-expression
	index         int      // the current index of the listed sub-expression
}

func (t *translator) translateWeekMonth() (bool, error) {
	var v, v1, v2 int
	var mtext string

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

	if t.moment != week && t.moment != month {
		return false, nil
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
	mtext := "every"
	if !t.base {
		mtext = "onn"
	}
	if t.ranged {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+mtext)+t.moment+" "+
			t.cronRange[0]+viper.GetString(configStr+"to")+t.cronRange[1], t.index == 0)
		translatedString += helper.GetStrIfTrue(t.cronRange[0]+viper.GetString(configStr+"to")+t.cronRange[1],
			t.index > 0)
	} else {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+mtext)+t.moment+" "+t.cron, t.index == 0)
		translatedString += helper.GetStrIfTrue(t.cron, t.index > 0)
	}
	translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), t.listed && t.index < t.cronListedLen-1)
}
