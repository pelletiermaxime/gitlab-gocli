package main

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
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
		log.Error("Please make sure your have a config.json.")
	}
	json.Unmarshal(configFile, &config)

	return config
}

func NewRequest(rawURL string) []byte {
	url := fmt.Sprintf("%s/%s?private_token=%s", config.Host, rawURL, config.PrivateToken)
	// log.Debug(url)
	response, err := http.Get(url)
	if err != nil {
		log.Error(err)
	}
	contents, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Error(err)
	}
	return contents
}
