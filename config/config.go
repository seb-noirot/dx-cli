package config

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
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
	fmt.Println("🤖 Loading config file...")

	// Read the YAML file
	data, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		fmt.Printf("🚨 Oops! No config files found.\n")
		return nil, err
	}

	fmt.Println("✅ Config file loaded.")

	// Unmarshal into Config struct
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Printf("🚨 Oops! Error parsing the config file: %s\n", err)
		return nil, err
	}

	fmt.Println("📋 Parsing contexts...")

	// Create a list of context names
	var contextNames []string
	for _, ctx := range cfg.Contexts {
		contextNames = append(contextNames, ctx.Name)
	}

	fmt.Println("✅ Contexts parsed.")

	// Use survey to get user's choice
	var selectedContextName string
	prompt := &survey.Select{
		Message: "🌍 Choose a context:",
		Options: contextNames,
	}
	fmt.Println("🤖 Awaiting your selection...")
	err = survey.AskOne(prompt, &selectedContextName)
	if err != nil {
		fmt.Println("🚨 Oops! Something went wrong.")
		return nil, err
	}

	fmt.Printf("🎉 You've chosen: %s\n", selectedContextName)

	// Find and return the selected context
	for _, ctx := range cfg.Contexts {
		if ctx.Name == selectedContextName {
			fmt.Println("✅ Context successfully set.")
			return &ctx, nil
		}
	}

	fmt.Println("🚨 Oops! Selected context not found.")
	return nil, nil // or return an error if context is not found
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
