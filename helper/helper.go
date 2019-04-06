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

// GetList takes a string & returns a slice separated by the separator provided & also
//converts the string to int
func GetList(str, seperator string) ([]int, bool) {
	ss := strings.Split(str, seperator)
	ii := []int{}
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		if s != "*" {
			ii = append(ii, i)
		}
	}
	if len(ii) > 1 {
		return ii, true
	}
	return ii, false
}
