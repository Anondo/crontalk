package translator

import (
	"crontalk/helper"
	"net/url"
	"strconv"
)

func validateSubExpressions(errs *url.Values, moment, c string) {
	vv, ranged := helper.GetList(c, "-")
	if !helper.IsDigit(c) && !ranged { //the value provided must be a digit
		errs.Add(moment+" value", "The value must a positive numeric digit or *")
	} else if ranged && (!helper.IsDigit(vv[0]) || !helper.IsDigit(vv[1])) {
		errs.Add(moment+" value", "The value must a positive numeric digit or *")
	} else { //checking the validity of the values in the context of each sub-expressions
		var v, vr1, vr2 int
		if ranged {
			vr1, _ = strconv.Atoi(vv[0])
			vr2, _ = strconv.Atoi(vv[1])
		} else {
			v, _ = strconv.Atoi(c)
		}

		if moment == minute {
			if (v < 0 || v > 59) && !ranged {
				errs.Add(minute+" value", "The value must be between 0 to 59")
			}
			if ranged {
				if (vr1 < 0 || vr1 > 59) && (vr2 < 0 || vr2 > 59) {
					errs.Add(minute+" value", "The value must be between 0 to 59")
				}
				if vr1 >= vr2 {
					errs.Add(minute+" value", "The starting range must be lower than the trailing range")
				}
			}
		} else if moment == hour {
			if (v < 0 || v > 23) && !ranged {
				errs.Add(hour+" value", "The value must be between 0 to 23")
			}
			if ranged {
				if (vr1 < 0 || vr1 > 23) && (vr2 < 0 || vr2 > 23) {
					errs.Add(hour+" value", "The value must be between 0 to 23")
				}
				if vr1 >= vr2 {
					errs.Add(hour+" value", "The starting range must be lower than the trailing range")
				}
			}
		} else if moment == day {
			if (v < 1 || v > 31) && !ranged {
				errs.Add(day+" value", "The value must be between 1 to 31")
			}
			if ranged {
				if (vr1 < 1 || vr1 > 31) && (vr2 < 1 || vr2 > 31) {
					errs.Add(day+" value", "The value must be between 1 to 31")
				}
				if vr1 >= vr2 {
					errs.Add(day+" value", "The starting range must be lower than the trailing range")
				}
			}
		} else if moment == month {
			if (v < 1 || v > 12) && !ranged {
				errs.Add(month+" value", "The value must be between 1 to 12")
			}
			if ranged {
				if (vr1 < 1 || vr1 > 12) && (vr2 < 1 || vr2 > 12) {
					errs.Add(month+" value", "The value must be between 1 to 12")
				}
				if vr1 >= vr2 {
					errs.Add(month+" value", "The starting range must be lower than the trailing range")
				}
			}
		} else if moment == week {

			if (v < 0 || v > 6) && !ranged {
				errs.Add(week+" value", "The Value must be between 0 to 6")
			}
			if ranged {
				if (vr1 < 0 || vr1 > 6) || (vr2 < 0 || vr2 > 6) {
					errs.Add(week+" value", "The value must be between 0 to 6")
				}
				if vr1 >= vr2 {
					errs.Add(week+" value", "The starting range must be lower than the trailing range")
				}
			}
		}
	}
}

func validateSteppedSubExpression(errs *url.Values, se, moment string) {
	steppedCron, _ := helper.GetList(se, "/")
	stepValue := steppedCron[1]
	val, err := strconv.Atoi(stepValue)

	if err != nil {
		errs.Add("Invalid Value", err.Error())
	}

	if moment == day {
		if val < 1 || val > 31 {
			errs.Add(day+" value", "The step value for day must be a numberic between 1 & 31")
		}
	}
	if moment == month {
		if val < 1 || val > 12 {
			errs.Add(month+" value", "The step value for month must be a numberic between 1 & 12")
		}
	}
	if moment == week {
		if val < 0 || val > 6 {
			errs.Add(week+" value", "The step value for week must be a numberic between 0 & 6")
		}
	}
	if moment == hour {
		if val < 0 || val > 23 {
			errs.Add(hour+" value", "The step value for hour must be between 0 to 23")
		}
	}
	if moment == minute {
		if val < 0 || val > 59 {
			errs.Add(minute+" value", "The step value for minute must be between 0 to 59")
		}
	}

}
