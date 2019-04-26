package binded

import (
	"html/template"
	"log"
)

// Assets contains the string representation of the binded assets
type Assets struct {
	IndexHTML   string
	StylesCSS   string
	TranslateJS string
}

func (a *Assets) populate() {
	a.IndexHTML = getBindedData("assets/index.html")
	a.StylesCSS = getBindedData("assets/css/styles.css")
	a.TranslateJS = getBindedData("assets/js/translate.js")
}

var assets Assets

func init() {
	assets.populate()
}

func getBindedData(a string) string {
	data, err := Asset(a)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(data)
}

// GetAssets returns the populated assets
func GetAssets() Assets {
	return assets
}

// HTML ...
func HTML(v string) template.HTML {
	return template.HTML(v)
}

// CSS ...
func CSS(v string) template.CSS {
	return template.CSS(v)
}

// JS ...
func JS(v string) template.JS {
	return template.JS(v)
}
