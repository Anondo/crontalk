package config

import (
	"encoding/json"

	"github.com/spf13/viper"
)

// LangConfig represents the language configuration
type LangConfig struct {
	Day                 string `json:"day"`
	Sunday              string `json:"sunday"`
	EveryDay            string `json:"every_day"`
	Onn                 string `json:"onn"`
	Tuesday             string `json:"tuesday"`
	AtEveryMinuteOfHour string `json:"at_every_minute_of_hour"`
	To                  string `json:"to"`
	January             string `json:"january"`
	April               string `json:"april"`
	Every               string `json:"every"`
	AtEveryMinute       string `json:"at_every_minute"`
	July                string `json:"july"`
	May                 string `json:"may"`
	June                string `json:"june"`
	From                string `json:"from"`
	Hour                string `json:"hour"`
	Friday              string `json:"friday"`
	February            string `json:"february"`
	October             string `json:"october"`
	December            string `json:"december"`
	OnMonthOf           string `json:"on_month_of"`
	DayOfTheMonth       string `json:"day_of_the_month"`
	Wednesday           string `json:"wednesday"`
	EveryMonthOf        string `json:"every_month_of"`
	And                 string `json:"and"`
	Thursday            string `json:"thursday"`
	Saturday            string `json:"saturday"`
	March               string `json:"march"`
	August              string `json:"august"`
	September           string `json:"september"`
	November            string `json:"november"`
	DayOfTheWeek        string `json:"day_of_the_week"`
	Monday              string `json:"monday"`
	Minute              string `json:"minute"`
	MonthOfTheYear      string `json:"month_of_the_year"`
	At                  string `json:"at"`
	Number0             string `json:"number0"`
	Number1             string `json:"number1"`
	Number2             string `json:"number2"`
	Number3             string `json:"number3"`
	Number4             string `json:"number4"`
	Number5             string `json:"number5"`
	Number6             string `json:"number6"`
	Number7             string `json:"number7"`
	Number8             string `json:"number8"`
	Number9             string `json:"number9"`
}

// ToMap transforms LanguageConfig instance to map
func (l LangConfig) ToMap() map[string]string {
	data := map[string]string{}
	jsonByte, _ := json.Marshal(l)
	json.Unmarshal(jsonByte, &data)
	return data
}

var langCfg = map[string]LangConfig{}

// LoadLanguage populates the language config map
func LoadLanguage() {
	langMap := viper.GetStringMap("language")
	for key, val := range langMap {
		cfg := val.(map[string]interface{})

		number0 := ""
		number1 := ""
		number2 := ""
		number3 := ""
		number4 := ""
		number5 := ""
		number6 := ""
		number7 := ""
		number8 := ""
		number9 := ""

		if cfg["0"] != nil {
			number0 = cfg["0"].(string)
		}
		if cfg["1"] != nil {
			number1 = cfg["1"].(string)
		}
		if cfg["2"] != nil {
			number2 = cfg["2"].(string)
		}
		if cfg["3"] != nil {
			number3 = cfg["3"].(string)
		}
		if cfg["4"] != nil {
			number4 = cfg["4"].(string)
		}
		if cfg["5"] != nil {
			number5 = cfg["5"].(string)
		}
		if cfg["6"] != nil {
			number6 = cfg["6"].(string)
		}
		if cfg["7"] != nil {
			number7 = cfg["7"].(string)
		}
		if cfg["8"] != nil {
			number8 = cfg["8"].(string)
		}
		if cfg["9"] != nil {
			number9 = cfg["9"].(string)
		}

		langCfg[key] = LangConfig{
			Day:                 cfg["day"].(string),
			Sunday:              cfg["sunday"].(string),
			EveryDay:            cfg["every_day"].(string),
			Onn:                 cfg["onn"].(string),
			Tuesday:             cfg["tuesday"].(string),
			AtEveryMinuteOfHour: cfg["at_every_minute_of_hour"].(string),
			To:                  cfg["to"].(string),
			January:             cfg["january"].(string),
			April:               cfg["april"].(string),
			Every:               cfg["every"].(string),
			AtEveryMinute:       cfg["at_every_minute"].(string),
			July:                cfg["july"].(string),
			May:                 cfg["may"].(string),
			June:                cfg["june"].(string),
			From:                cfg["from"].(string),
			Hour:                cfg["hour"].(string),
			Friday:              cfg["friday"].(string),
			February:            cfg["february"].(string),
			October:             cfg["october"].(string),
			OnMonthOf:           cfg["on_month_of"].(string),
			DayOfTheMonth:       cfg["day_of_the_month"].(string),
			Wednesday:           cfg["wednesday"].(string),
			EveryMonthOf:        cfg["every_month_of"].(string),
			And:                 cfg["and"].(string),
			Thursday:            cfg["thursday"].(string),
			Saturday:            cfg["saturday"].(string),
			March:               cfg["march"].(string),
			August:              cfg["august"].(string),
			September:           cfg["september"].(string),
			November:            cfg["november"].(string),
			December:            cfg["december"].(string),
			DayOfTheWeek:        cfg["day_of_the_week"].(string),
			Monday:              cfg["monday"].(string),
			Minute:              cfg["minute"].(string),
			MonthOfTheYear:      cfg["month_of_the_year"].(string),
			At:                  cfg["at"].(string),
			Number0:             number0,
			Number1:             number1,
			Number2:             number2,
			Number3:             number3,
			Number4:             number4,
			Number5:             number5,
			Number6:             number6,
			Number7:             number7,
			Number8:             number8,
			Number9:             number9,
		}
	}
}

// Language returns a specific alraedy populated LangConfig instance
func Language(lang string) LangConfig {
	return langCfg[lang]
}

// LanguageMap returns the LangConfig map instance
func LanguageMap() map[string]LangConfig {
	return langCfg
}
