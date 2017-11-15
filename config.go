package gobot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var DefaultConfig = &Config{
	Token:   "",
	Prefix:  "~",
	ApiKey:  "",
	OwnerID: "",
}

// Config maps how a general config should be.
type Config struct {
	Token   string `json:"token"`
	ApiKey  string `json:"api_key"`
	Prefix  string `json:"prefix"`
	OwnerID string `json:"owner_id"`
}

// ReadConfigFromFile reads from disk the config file, unmarshal's it into the Go struct and sets it as the DefaultConfig.
func ReadConfigFromFile(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("no config found, exiting.... (%s)", err)
	}
	config := &Config{}
	if err := json.Unmarshal(file, config); err != nil {
		return fmt.Errorf("failed parsing json, exiting.... (%s)", err)
	}
	DefaultConfig = config
	return nil
}

func ReadJsonFile(filename string) map[string]int {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Errorf("no config found, exiting.... (%s)", err)
		return nil
	}

	var test map[string]int

	if err := json.Unmarshal(file, &test); err != nil {
		fmt.Errorf("failed parsing json, exiting.... (%s)", err)
		return nil
	}

	return test
}

func SaveJsonFile(v interface{}, path string) {
	fo, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	e := json.NewEncoder(fo)
	if err := e.Encode(v); err != nil {
		panic(err)
	}
}
