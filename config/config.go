package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

const (
// ConfigFilePath = "/Users/snoirot/GolandProjects/dx-cli/contexts.yaml"
)

func GetConfigFilePath() (string, error) {
	// Get the path to the current executable
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	// Get the directory of the executable
	dirPath := filepath.Dir(exePath)

	// Join the directory path with "contexts.yaml"
	configFilePath := filepath.Join(dirPath, "contexts.yaml")

	return configFilePath, nil
}

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
	Onboarding         *Onboarding         `yaml:"onboarding"`
}

type Onboarding struct {
	OnboardingCmds []OnboardingCmd `yaml:"command"`
}

type OnboardingCmd struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Execution   string `yaml:"execution"`
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

func UpdateCurrentContext(updatedContext *Context) error {
	// Read the existing YAML file
	path, err := GetConfigFilePath()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	data, err := os.ReadFile(path)
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
	err = os.WriteFile(path, updatedData, 0644)
	if err != nil {
		return err
	}

	return nil
}
