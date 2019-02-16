package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string `json:"token"`
}

var (
	config Config
)

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Println(err)
	}
}
