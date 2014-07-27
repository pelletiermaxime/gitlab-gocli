package main

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
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

func ReadRequest(response *http.Response) []byte {
	contents, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Error(err)
	}
	return contents
}

func NewRequest(rawURL string) []byte {
	return NewRequestGET(rawURL)
}

func NewRequestGET(rawURL string) []byte {
	url := fmt.Sprintf("%s/%s?private_token=%s", config.Host, rawURL, config.PrivateToken)
	response, err := http.Get(url)
	if err != nil {
		log.Error(err)
	}
	content := ReadRequest(response)
	return content
}

func NewRequestPOST(rawURL string, values url.Values) []byte {
	url := fmt.Sprintf("%s/%s?private_token=%s", config.Host, rawURL, config.PrivateToken)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Error(err)
	}
	content := ReadRequest(response)
	return content
}

func NewRequestDELETE(rawURL string) []byte {
	url := fmt.Sprintf("%s/%s?private_token=%s", config.Host, rawURL, config.PrivateToken)
	// response, err := http.PostForm(url, values)
	var pomme io.Reader
	request, err := http.NewRequest("DELETE", url, pomme)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Error(err)
	}
	content := ReadRequest(response)
	return content
}
