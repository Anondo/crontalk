package helper

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
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

// Get12Hour takes a 24hr format string & returns its 12hr format
func Get12Hour(h string) (string, error) {
	hi, err := strconv.Atoi(h)
	if err != nil {
		return "", err
	}
	hs := strconv.Itoa(hi - 12)
	return hs, nil
}
