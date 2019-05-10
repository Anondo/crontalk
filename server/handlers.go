package server

import (
	"bytes"
	"crontalk/binded"
	"crontalk/helper"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	translator "crontalk/translator.go"
	"github.com/spf13/viper"
)

func translateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	req := struct {
		Exprsn string `json:"expression"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Couldnt read request body")
		log.Printf("%s %d http://%s%s\n", r.Method, http.StatusBadRequest, r.Host, r.URL.Path)
		return
	}

	translator.CronExprsn = req.Exprsn

	translator.Init()

	if vErr := translator.Validate(); len(vErr) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		for k, v := range vErr {
			fmt.Fprintf(w, "%v: %v\n", k, v)
		}
		log.Printf("%s %d http://%s%s\n", r.Method, http.StatusBadRequest, r.Host, r.URL.Path)
		translator.GetTranslatedStr()
		return
	}

	if err := translator.Translate(); err != nil {
		log.Println(err.Error())
		translator.GetTranslatedStr()
		return
	}

	output := translator.GetTranslatedStr()
	output = helper.TrimExtraSpaces(output)

	if viper.GetBool("bangla") {
		helper.ChangeDigitLanguage(&output, "bangla") //changing the english digits to bangla
	}

	output = helper.AddOrdinals(output)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, output)
	log.Printf("%s %d http://%s%s\n", r.Method, http.StatusOK, r.Host, r.URL.Path)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {

	ast := binded.GetAssets()
	data := struct {
		Assets binded.Assets
	}{
		Assets: ast,
	}
	tm := template.New("main")
	tm.Delims("@{{", "}}@")
	tm.Funcs(template.FuncMap{
		"html": binded.HTML,
		"css":  binded.CSS,
		"js":   binded.JS,
	})
	t, err := tm.Parse(ast.IndexHTML)
	if err != nil {
		log.Fatal(err.Error())
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err.Error())
	}
	w.Header().Add("Content-Type", "text/html")
	w.Write(buf.Bytes())
}
