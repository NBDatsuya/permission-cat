package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	WebSave struct {
		Proxy string `json:"proxy"`
	} `json:"webSave"`

	Sql2Struct struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Charset  string `json:"charset"`
		Type     string `json:"type"`
	} `json:"sql2Struct"`
}

var Conf Config

func LoadConfig() {
	data, err := os.ReadFile("config.json")

	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}

	err = json.Unmarshal(data, &Conf)
}

func init() {
	LoadConfig()
}
