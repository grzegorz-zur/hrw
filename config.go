package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

// Configuration file name.
const ConfigurationFileName = ".config/hrw.json"

// Configuration for router and ping.
type Configuration struct {
	Router   Router `json:"router"`
	Interval int    `json:"interval"`
	Ping     Ping   `json:"ping"`
}

// Router configuration.
type Router struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

// Ping configuration.
type Ping struct {
	Address string `json:"address"`
	Timeout int    `json:"timeout"`
}

// ReadConfiguration retrieves configuration from file in user HOME directory.
func ReadConfiguration() (*Configuration, error) {
	home := os.Getenv("HOME")
	file := path.Join(home, ConfigurationFileName)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	config := &Configuration{}
	err = json.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
