package main

import (
	"log"
	"net/http"
	//"net/url"
)

func main() {
	config := GetConfig()
	checker := Checker{config: config}
	log.Println(config)
	http.Handle("/check", checker)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}
