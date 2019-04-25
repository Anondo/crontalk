package server

import (
	"crontalk/helper"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	translator "crontalk/translator.go"
	"github.com/spf13/viper"
)

func translateHandler(w http.ResponseWriter, r *http.Request) {
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
		return
	}

	if err := translator.Translate(); err != nil {
		log.Println(err.Error())
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
