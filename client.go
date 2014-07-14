package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Config is for the information you need to config
type Config struct {
	Host         string
	PrivateToken string
}

var config *Config

func parseConfig() *Config {
	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("Please make sure your have a config.json.")
	}
	json.Unmarshal(configFile, &config)

	return config
}

func NewRequest(rawURL string) []byte {
	url := fmt.Sprintf("%s/%s?private_token=%s", config.Host, rawURL, config.PrivateToken)
	response, err := http.Get(url)
	if err != nil {
		assert(err)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		assert(err)
	}
	return contents
}
