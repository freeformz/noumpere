package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Checker struct {
	config Config
}

func (chk Checker) ValidateParams(r *http.Request) (errors []string) {
	if _, present := r.Form["metric"]; !present {
		errors = append(errors, "metric is required")
	}

	rangeValues, rangePresent := r.Form["range"]

	if !rangePresent {
		errors = append(errors, "range is required")
	} else {
		if len(rangeValues) > 0 {
			errors = append(errors, "range valud only once")
		} else {
			_, err := strconv.Atoi(rangeValues[0])
			if err != nil {
				errors = append(errors, err.Error())
			}
		}
	}

	if _, present := r.Form["range"]; present != true {
		errors = append(errors, "range is required")
	}

	_, minPresent := r.Form["min"]
	_, maxPresent := r.Form["max"]
	if !minPresent && !maxPresent {
		errors = append(errors, "one of min or max are required")
	}

	emptyokValue, emptyokPresent := r.Form["empty_ok"]
	if emptyokPresent {
		if len(emptyokValue) > 1 {
			errors = append(errors, "empty_ok valid only once")
		} else {
			var validEmptyOkFound bool
			for i := 0; i < len(chk.config.ValidEmptyOkValues); i++ {
				if chk.config.ValidEmptyOkValues[i] == emptyokValue[0] {
					validEmptyOkFound = true
					break
				}
			}
			if !validEmptyOkFound {
				errors = append(errors, "empty_ok must be one of yes/y/1/true")
			}
		}
	}
	return
}

func writeErrors(errors []string, w http.ResponseWriter) {
	errWrapper := make(map[string][]string)
	errWrapper["errors"] = errors
	errorMessages, err := json.Marshal(errWrapper)
	if err != nil {
		log.Fatal(err)
	}
	h := w.Header()
	h["Content-Type"] = []string{"application/json"}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(errorMessages)
}

func (chk Checker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	r.ParseForm()
	if errors := chk.ValidateParams(r); len(errors) > 0 {
		writeErrors(errors, w)
		return
	}

	log.Println(r.Form)
}
