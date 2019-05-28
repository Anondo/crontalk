package helper

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

// THe app logo
const (
	Logo = `
 _______   _______  _______  _       _________ _______  _        _
 (  ____ \(  ____ )(  ___  )( (    /|\__   __/(  ___  )( \      | \    /\
 | (    \/| (    )|| (   ) ||  \  ( |   ) (   | (   ) || (      |  \  / /
 | |      | (____)|| |   | ||   \ | |   | |   | (___) || |      |  (_/ /
 | |      |     __)| |   | || (\ \) |   | |   |  ___  || |      |   _ (
 | |      | (\ (   | |   | || | \   |   | |   | (   ) || |      |  ( \ \
 | (____/\| ) \ \__| (___) || )  \  |   | |   | )   ( || (____/\|  /  \ \
 (_______/|/   \__/(_______)|/    )_)   )_(   |/     \|(_______/|_/    \/
                                                                         `
)

//TrimExtraSpaces Trims any extra space (should be one space gaps)
func TrimExtraSpaces(s string) string {
	space := regexp.MustCompile(`\s+`)
	s = space.ReplaceAllString(s, " ")
	return s
}

// PrettyTime takes hour and minute and returns then with a proper time layout
func PrettyTime(h string, m string) (string, error) {
	layout12 := "03:04PM"
	layout24 := "15:04"
	hi, err := strconv.Atoi(h)
	if err != nil {
		return "", err
	}
	mi, err := strconv.Atoi(m)
	if err != nil {
		return "", err
	}
	hstr := ""
	mstr := ""

	if hi < 10 {
		hstr = "0" + strconv.Itoa(hi)
	} else {
		hstr = strconv.Itoa(hi)
	}
	if mi < 10 {
		mstr = "0" + strconv.Itoa(mi)
	} else {
		mstr = strconv.Itoa(mi)
	}

	t, err := time.Parse(layout24, hstr+":"+mstr)

	if err != nil {
		log.Fatal(err)
	}

	return t.Format(layout12), nil

}

// GetList takes a string & returns a slice separated by the separator provided & true/false based on the length of the slice
func GetList(str, seperator string) ([]string, bool) {
	ss := strings.Split(str, seperator)
	if len(ss) > 1 {
		return ss, true
	}
	return ss, false
}

// IsDigit determines whether the given string is a digit or not
func IsDigit(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	}

	return false
}

//GetStrIfTrue returns the given string if the provided bool is true
func GetStrIfTrue(s string, l bool) string {
	if l {
		return s
	}
	return ""
}

// ChangeDigitLanguage changes the language of any numeric chars in the string
func ChangeDigitLanguage(str *string, lang string) {
	configStr := "language." + lang + "."
	for _, c := range *str {
		cs := string(c)
		if IsDigit(cs) {
			char := viper.GetString(configStr + cs)
			*str = strings.Replace(*str, cs, char, -1)
		}
	}
}

// AddOrdinals add ordinal indicators  like 1 -> 1st 2 -> 2nd and so on
func AddOrdinals(s string) string {
	theNumber := ""
	ss := strings.Split(s, "")
	for i := 0; i < len(ss)-1; i++ { // cant go to the last element because s[i+1] will produce runtime error
		c := ss[i]
		nextC := ss[i+1]
		if IsDigit(c) {
			theNumber += c
			if IsDigit(nextC) {
				continue
			} else if nextC == " " {
				newC := addOrdinalIndicator(c)
				ss[i] = newC
				theNumber = ""
			} else {
				theNumber = ""
			}
		}

	}
	return strings.Join(ss, "")
}

func addOrdinalIndicator(s string) string {
	switch s {
	case "1":
		s += "st"
	case "2":
		s += "nd"
	case "3":
		s += "rd"
	default:
		s += "th"
	}
	return s
}

// IndexOf returns the index of the first occurence of the element matched in the string slice
func IndexOf(ss []string, e string) int {
	for i, s := range ss {
		if s == e {
			return i
		}
	}
	return -1
}

// RunPrompt executes a prompt to chose an item from the cli
func RunPrompt(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, res, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return res, err
}
