package config

import (
	"fmt"
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
	GitLabContexts     []GitLabContext     `yaml:"gitlabContexts"`
	KubernetesContexts []KubernetesContext `yaml:"k8s"`
	MavenContext       *MavenContext       `yaml:"mvn"`
}

type MavenContext struct {
	Settings Settings `yaml:"settings"`
}

type Settings struct {
	Content string  `yaml:"content"`
	Tokens  []Token `yaml:"tokens"`
}

type Token struct {
	Name        string  `yaml:"name"`
	Description *string `yaml:"description"`
	Link        *string `yaml:"link"`
}

type KubernetesContext struct {
	ClusterName string `yaml:"name"`
	Certificate string `yaml:"certificate"`
	ADUser      string `yaml:"ad_user"`
	ServerId    string `yaml:"server_id"`
	ClientId    string `yaml:"client_id"`
	TenantId    string `yaml:"tenant_id"`
}

type GitLabContext struct {
	Name         string        `yaml:"name"`
	Host         string        `yaml:"host"`
	GitlabStacks []GitlabStack `yaml:"stacks"`
}

type GitlabStack struct {
	Name     string   `yaml:"name"`
	Path     string   `yaml:"path"`
	Projects []string `yaml:"projects"`
	Javas    []string `yaml:"javas"`
}

// GetCurrentContext reads the YAML file and returns the current Context object
func GetCurrentContext() (*Context, error) {
	// Read the YAML file
	data, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		fmt.Printf("Not config files")
		return nil, err
	}

	// Unmarshal into Config struct
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("Error getting config files. Check structure: %s", err)
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
