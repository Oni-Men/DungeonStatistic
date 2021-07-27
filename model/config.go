package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Host string `json:"host"`
}

func LoadConfig(cfg *Config) {
	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err = json.Unmarshal(data, cfg); err != nil {
		log.Fatalf(err.Error())
	}
}
