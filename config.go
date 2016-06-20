package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config is a global config
type Config struct {
	BindAddress string `json:"address"`
	Port        int    `json:"port"`
	Dbfile      string `json:"dbfile"`
}

// LoadConfigFromFile loads configuration from file into Config structure
func LoadConfigFromFile(filename string) *Config {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &config
}
