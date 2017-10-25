package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	config *Configuration
)

type Configuration struct {
	Web struct {
		Port string `json:"port"`
	} `json:"web"`
	Db struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
	} `json:"db"`
	Security struct {
		SaltLength int `json:"saltLength"`
		BcryptCost int `json:"bcryptCost"`
	} `json:"security"`
}

func GetConfig() (*Configuration, error) {
	if config != nil {
		return config, nil
	}

	file, err := os.Open("config.local.json")
	if err != nil {
		fmt.Println("Error importing the config file")
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		fmt.Println("Error decoding the configuration file")
		return nil, err
	}
	return config, nil
}
