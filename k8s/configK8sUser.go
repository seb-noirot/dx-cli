package k8s

import (
	"dx-cli/config"
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var configK8sUser = &cobra.Command{
	Use:   "user",
	Short: "config a user",
	Run: func(cmd *cobra.Command, args []string) {
		// Select Kubernetes context
		k8sContext, err := utils.SelectK8sCluster()
		if err != nil {
			return
		}
		if k8sContext == nil {
			return
		}

		UpdateKubeConfigWithADUser(k8sContext)

		fmt.Printf("Kubernetes user for %s has been set up.\n", k8sContext.ClusterName)

	},
}

func UpdateKubeConfigWithADUser(k8sContext *config.KubernetesContext) error {
	// Read the existing config
	kubeConfigFile := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	data, err := os.ReadFile(kubeConfigFile)
	if err != nil {
		return err
	}

	var cfg KubeConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	// Replace or add the user
	newUser := UserEntry{
		Name: k8sContext.ADUser,
		User: UserSpec{
			Exec: UserSpecExec{
				ApiVersion: "client.authentication.k8s.io/v1beta1",
				Args: []string{
					"get-token",
					"--environment",
					"AzurePublicCloud",
					"--server-id",
					k8sContext.ServerId,
					"--client-id",
					k8sContext.ClientId,
					"--tenant-id",
					k8sContext.TenantId,
				},
				Command: "kubelogin",
				Env:     "",
			},
		},
	}

	// Create a map of existing users for easier manipulation
	userMap := make(map[string]UserEntry)
	for _, user := range cfg.Users {
		userMap[user.Name] = user
	}

	userMap[k8sContext.ADUser] = newUser

	// Reconstruct the Users slice from the map
	newUsers := make([]UserEntry, 0, len(userMap))
	for _, user := range userMap {
		newUsers = append(newUsers, user)
	}
	cfg.Users = newUsers

	// Write the updated config back to the file
	updatedData, err := yaml.Marshal(&cfg)
	fmt.Printf("New updated data %s", updatedData)
	if err != nil {
		fmt.Printf("Error creating yaml %s", err)
		return err
	}
	err = os.WriteFile(kubeConfigFile, updatedData, 0644)
	if err != nil {
		fmt.Printf("Error writing config file %s", err)

		return err
	}

	return nil
}

func init() {
	KubeCmd.AddCommand(configK8sUser)
}
