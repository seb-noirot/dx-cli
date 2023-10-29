package utils

import (
	"dx-cli/config"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadConfig(configFilePath string, verbose bool) (config.Config, error) {
	var conf config.Config

	Println(verbose, "🤖 Loading config file...")

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("🚨 Oops! No config files found.\n")
		return conf, err
	}

	Println(verbose, "✅ Config file loaded.")

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		Printf(verbose, "🚨 Oops! Error parsing the config file: %s\n", err)
		return conf, err
	}

	Println(verbose, "📋 Parsing contexts...")

	return conf, nil
}
