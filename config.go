package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configuration struct {
	ConsumerKey    string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	AccessToken    string `json:"accessToken"`
	AccessSecret   string `json:"accessSecret"`
	FortunePath    string `json:"fortunePath"`
	SleepMin       int    `json:"sleepMin"`
	SleepMax       int    `json:"sleepMax"`
}

var config Configuration

func pullConfig(configPath string) {
	raw, err := ioutil.ReadFile(configPath)

	if err != nil {
		logOut("Error reading config file...")
		os.Exit(1)
	}

	json.Unmarshal(raw, &config)
}
