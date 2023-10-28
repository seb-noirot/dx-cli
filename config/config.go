package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

const (
	ConfigFilePath = "/Users/snoirot/GolandProjects/dx-cli/contexts.yaml"
)

type Config struct {
	Contexts       []Context `yaml:"contexts"`
	CurrentContext string    `yaml:"currentContext"`
}

type Context struct {
	Name string `yaml:"name"` // Placeholder field
	// Add other fields here as needed
	GitLabContexts []GitLabContext `yaml:"gitlabContexts"`
}

type GitLabContext struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
}

// GetCurrentContext reads the YAML file and returns the current Context object
func GetCurrentContext() (*Context, error) {
	// Read the YAML file
	data, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal into Config struct
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	// Find and return the current context
	for _, ctx := range cfg.Contexts {
		if ctx.Name == cfg.CurrentContext {
			return &ctx, nil
		}
	}
	return nil, nil // or return an error if current context is not found
}

func UpdateCurrentContext(updatedContext *Context) error {
	// Read the existing YAML file
	data, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return err
	}

	// Unmarshal into Config struct
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	// Update the current context in Config struct
	for i, ctx := range cfg.Contexts {
		if ctx.Name == cfg.CurrentContext {
			cfg.Contexts[i] = *updatedContext
			break
		}
	}

	// Marshal back to YAML
	updatedData, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}

	// Write the updated YAML back into the file
	err = os.WriteFile(ConfigFilePath, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
