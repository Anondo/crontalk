package translator

import (
	"crontalk/helper"
	"strconv"

	"github.com/spf13/viper"
)

// TODO: Fix normal hour ranged minute bug
// TODO: fix single 24hour hour only bug
// TODO: convert hour in bangla

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
	//---------------------start iterating from the last sub-expressions to determine the starting string ------------------
	for i = weekIndex; i > hourIndex; i-- {
		if cronSlice[i] != anyValue {
			cc, listed := helper.GetList(cronSlice[i], ",")
			//-----------------------------iterating because values can be listed------------------------
			for j, c := range cc {
				rr, ranged := helper.GetList(c, "-")
				//------------------------------- checking for weekly----------------------------------------
				if moments[i] == week {
					var wi, wi1, wi2 int
					if ranged {
						var err1, err2 error
						wi1, err1 = strconv.Atoi(rr[0])
						wi2, err2 = strconv.Atoi(rr[1])
						if err1 != nil || err2 != nil {
							return err1 //any one will do
						}
					} else {
						var err error
						wi, err = strconv.Atoi(c)
						if err != nil {
							return err
						}
					}

					//the following is for proper & meaningful sentence
					if ranged {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+weeks[wi1]+
							viper.GetString(configStr+"to")+weeks[wi2], j == 0) //if this is the first check
						translatedString += helper.GetStrIfTrue(weeks[wi1]+viper.GetString(configStr+"to")+weeks[wi2],
							j > 0) //just keep adding the value & not the full sentence

					} else {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+weeks[wi], j == 0) //if this is the first check
						translatedString += helper.GetStrIfTrue(weeks[wi], j > 0)                                     //just keep adding the value & not the full sentence
					}
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1) //print "and" for listed
					continue

				}
				//----------------------------------------------------------------------------------------
				//------------- checking for monthly, in seperate if blocks because , slight change of translated string-------------
				if moments[i] == month {
					var mi, mi1, mi2 int
					if ranged {
						var err1, err2 error
						mi1, err1 = strconv.Atoi(rr[0])
						mi2, err2 = strconv.Atoi(rr[1])
						if err1 != nil || err2 != nil {
							return err1 //again, any one will do
						}
					} else {
						var err error
						mi, err = strconv.Atoi(c)
						if err != nil {
							return err
						}
					}

					if ranged {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+months[mi1]+
							viper.GetString(configStr+"to")+months[mi2], j == 0)
						translatedString += helper.GetStrIfTrue(months[mi1]+viper.GetString(configStr+"to")+
							months[mi2], j > 0)
					} else {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+months[mi], j == 0)
						translatedString += helper.GetStrIfTrue(months[mi], j > 0)
					}
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1)
					continue
				}
				//---------------------------------------------------------------------------------------
				////--------------------------------checking for the day--------------------------------------
				if ranged {
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+moments[i]+" "+
						rr[0]+viper.GetString(configStr+"to")+rr[1], j == 0)
					translatedString += helper.GetStrIfTrue(rr[0]+viper.GetString(configStr+"to")+rr[1],
						j > 0)
				} else {
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"every")+moments[i]+" "+c, j == 0)
					translatedString += helper.GetStrIfTrue(c, j > 0)
				}
				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1)
				continue
				//---------------------------------------------------------------------------------------------
			}
			//---------------------------------------------------------------------------------------------
			break
		}
	}
	//---------------------------------------------------------------------------------------
	if i == hourIndex { // checking if every sub-expression contains asteriks apart from the time part
		translatedString += viper.GetString(configStr + "every_day")
	}
	baseIndex = i //storing the base index so that when checking every other than time , the base is also omitted because its
	//already checked

	return nil

}

func translateAllButBaseTimeOccurence() error {
	//------- checking every other sub-expressions apart from the base and time, no need for reverse travel---------
	for i := dayIndex; i <= weekIndex; i++ {
		// ----------------not gonna check the base ---------------------------
		if cronSlice[i] != anyValue && i != baseIndex {
			cc, listed := helper.GetList(cronSlice[i], ",")
			//-------------------------iterating the single sub-expressions-------------------------
			for j, c := range cc {
				rr, ranged := helper.GetList(c, "-")
				if moments[i] == week {
					var wi, wi1, wi2 int
					if ranged {
						var err1, err2 error
						wi1, err1 = strconv.Atoi(rr[0])
						wi2, err2 = strconv.Atoi(rr[1])
						if err1 != nil || err2 != nil {
							return err1
						}
					} else {
						var err error
						wi, err = strconv.Atoi(c)
						if err != nil {
							return err
						}
					}
					if ranged {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"on")+weeks[wi1]+
							viper.GetString(configStr+"to")+weeks[wi2], j == 0)
						translatedString += helper.GetStrIfTrue(weeks[wi1]+viper.GetString(configStr+"to")+
							weeks[wi2], j > 0)
					} else {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"on")+weeks[wi], j == 0)
						translatedString += helper.GetStrIfTrue(weeks[wi], j > 0)
					}
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1)
				} else if moments[i] == month {
					var mi, mi1, mi2 int
					if ranged {
						var err1, err2 error
						mi1, err1 = strconv.Atoi(rr[0])
						mi2, err2 = strconv.Atoi(rr[1])
						if err1 != nil || err2 != nil {
							return err1
						}
					} else {
						var err error
						mi, err = strconv.Atoi(c)
						if err != nil {
							return err
						}
					}
					if ranged {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"on_month_of")+months[mi1]+
							viper.GetString(configStr+"to")+months[mi2], j == 0)
						translatedString += helper.GetStrIfTrue(months[mi1]+viper.GetString(configStr+"to")+
							months[mi2], j > 0)
					} else {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"on_month_of")+months[mi], j == 0)
						translatedString += helper.GetStrIfTrue(months[mi], j > 0)
					}
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1)

				} else {
					if ranged {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"onn")+moments[i]+" "+
							rr[0]+viper.GetString(configStr+"to")+rr[1], j == 0)
						translatedString += helper.GetStrIfTrue(rr[0]+viper.GetString(configStr+"to")+rr[1],
							j > 0)
					} else {
						translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"onn")+moments[i]+" "+c, j == 0) //no breaks like base and a bit different string
						translatedString += helper.GetStrIfTrue(c, j > 0)
					}
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && j < len(cc)-1)

				}
			}
			//------------------------------------------------------------------------------

		}
		//-----------------------------------------------------------------------------
	}
	//----------------------------------------------------------------------------------------
	return nil
}

func translateTimeOccurence() error {
	if cronSlice[minuteIndex] == anyValue && cronSlice[hourIndex] == anyValue { // checking if both hour and minute are defaults
		translatedString += viper.GetString(configStr + "at_every_minute")
	} else if cronSlice[minuteIndex] != anyValue && cronSlice[hourIndex] != anyValue { //checking if non of them are
		m := cronSlice[minuteIndex]
		h := cronSlice[hourIndex]
		mm, listedM := helper.GetList(m, ",")
		hh, listedH := helper.GetList(h, ",")
		for i, min := range mm {
			for j, hr := range hh {
				mrr, rangedM := helper.GetList(min, "-")
				hrr, rangedH := helper.GetList(hr, "-")
				if rangedM {
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+
						moments[minuteIndex]+" "+mrr[0]+viper.GetString(configStr+"to")+mrr[1], i == 0 && j == 0)
					translatedString += helper.GetStrIfTrue(mrr[0]+viper.GetString(configStr+"to")+
						mrr[1], i > 0 || j > 0)
				}
				if rangedH {
					hr1, err1 := helper.Get12Hour(hrr[0])
					hr2, err2 := helper.Get12Hour(hrr[1])
					if err1 != nil || err2 != nil {
						return err1
					}
					translatedString += helper.GetStrIfTrue(" "+moments[hourIndex]+" "+hr1+viper.GetString(configStr+"to")+
						hr2, i == 0 && j == 0)
					translatedString += helper.GetStrIfTrue(hr1+viper.GetString(configStr+"to")+
						hr2, i > 0 || j > 0)
				}
				if !rangedM && !rangedH {
					pt, err := helper.PrettyTime(hr, min)
					if err != nil {
						return err
					}
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+pt, i == 0 && j == 0)
					translatedString += helper.GetStrIfTrue(pt, i > 0 || j > 0)
				}

				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), (listedM || listedH) &&
					(i < len(mm)-1) || (j < len(hh)-1))
			}
		}

	} else { // checking if  just one of them is default
		mStr := moments[minuteIndex] // assuming minute is not default
		mVal := cronSlice[minuteIndex]
		if mVal == anyValue { //if so
			hVal := cronSlice[hourIndex]
			hh, listed := helper.GetList(hVal, ",")
			for i, hr := range hh {
				hrr, ranged := helper.GetList(hr, "-")
				if ranged {
					hr1, err1 := helper.Get12Hour(hrr[0])
					hr2, err2 := helper.Get12Hour(hrr[1])
					if err1 != nil || err2 != nil {
						return err1
					}
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at_every_minute_of_hour")+hr1+
						viper.GetString(configStr+"to")+hr2, i == 0)
					translatedString += helper.GetStrIfTrue(hr1+viper.GetString(configStr+"to")+hr2, i > 0)

				} else {
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at_every_minute_of_hour")+hr, i == 0)
					translatedString += helper.GetStrIfTrue(hr, i > 0)
				}

				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && i < len(hh)-1)
			}
		} else {
			mm, listed := helper.GetList(mVal, ",")
			for i, min := range mm {
				mr, ranged := helper.GetList(min, "-")
				if ranged {
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+mStr+" "+mr[0]+
						viper.GetString(configStr+"to")+mr[1], i == 0)
					translatedString += helper.GetStrIfTrue(mr[0]+viper.GetString(configStr+"to")+mr[1], i > 0)
				} else {
					translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"at")+mStr+" "+mVal, i == 0)
					translatedString += helper.GetStrIfTrue(min, i > 0)
				}

				translatedString += helper.GetStrIfTrue(viper.GetString(configStr+"and"), listed && i < len(mm)-1)
			}
		}
	}
	return nil
}
