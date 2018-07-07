package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config represents the content of config.json file
type Config struct {
	Version string
	RPC     RPC `json:"rpc"`
}

// RPC options
type RPC struct {
	Port string `json:"port"`
}

// LoadConfig loads the config.json file and unmarshals it
func LoadConfig() Config {
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var config Config
	json.Unmarshal(raw, &config)

	return config
}
