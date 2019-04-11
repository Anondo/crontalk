package translator

import (
	"crontalk/helper"
	"strconv"

	"github.com/spf13/viper"
)

func translateBaseWeekMonth(c string, ranged, listed bool, rr, cc []string, j int, m string) (bool, error) {
	var v, v1, v2 int
	if ranged {
		var err1, err2 error
		v1, err1 = strconv.Atoi(rr[0])
		v2, err2 = strconv.Atoi(rr[1])
		if err1 != nil || err2 != nil {
			return false, err1 //any one will do
		}
	} else {
		var err error
		v, err = strconv.Atoi(c)
		if err != nil {
			return false, err
		}
	}

	if m != week && m != month {
		return false, nil
	}
	moment := weeks
	if m == month {
		moment = months
	}

	//the following is for proper & meaningful sentence
	if ranged {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+moment[v1]+
			viper.GetString(configStr+"to")+moment[v2], j == 0) //if this is the first check
		translatedString += helper.GetStrIfTrue(moment[v1]+viper.GetString(configStr+"to")+moment[v2],
			j > 0) //just keep adding the value & not the full sentence

	} else {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+moment[v], j == 0) //if this is the first check
		translatedString += helper.GetStrIfTrue(moment[v], j > 0)                                     //just keep adding the value & not the full sentence
	}
	translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1) //print "and" for listed
	return true, nil
}

func translateBaseDay(ranged, listed bool, rr, cc []string, j int, c string, moment string) error {
	if ranged {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+moment+" "+
			rr[0]+viper.GetString(configStr+"to")+rr[1], j == 0)
		translatedString += helper.GetStrIfTrue(rr[0]+viper.GetString(configStr+"to")+rr[1],
			j > 0)
	} else {
		translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+moment+" "+c, j == 0)
		translatedString += helper.GetStrIfTrue(c, j > 0)
	}
	translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1)
	return nil
}
