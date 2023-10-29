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

func GetCurrentContext(configFilePath string, verbose bool) (*config.Context, error) {

	cfg, err := LoadConfig(config.ConfigFilePath, verbose)

	if err != nil {
		return nil, err
	}

	if cfg.CurrentContext == "" {
		Println(true, "🚫 No current context defined.")
		return nil, fmt.Errorf("no current context defined")
	}

	Println(verbose, "📋 Locating current context...")

	// Find and return the current context
	for _, ctx := range cfg.Contexts {
		if ctx.Name == cfg.CurrentContext {
			Println(verbose, "✅ Current context fetched successfully.")
			return &ctx, nil
		}
	}

	Println(true, "🚨 Oops! Current context not found.")
	return nil, fmt.Errorf("current context not found")
}
