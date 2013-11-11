package main

import (
	"os"
	"sort"
)

const (
	VERSION = "0.0.1"
)

type Config struct {
	GraphiteUrl        string
	LibratoEmail       string
	LibratoKey         string
	ApiKey             string
	Port               string
	ValidEmptyOkValues []string
}

func GetConfig() (config Config) {
	config.GraphiteUrl = os.Getenv("GRAPHITE_URL")
	config.LibratoEmail = os.Getenv("LIBRATO_EMAIL")
	config.LibratoKey = os.Getenv("LIBRATO_KEY")
	config.ApiKey = os.Getenv("API_KEY")
	config.Port = os.Getenv("PORT")
	config.ValidEmptyOkValues = []string{"1", "true", "y", "yes"}
	sort.Strings(config.ValidEmptyOkValues)
	return
}
