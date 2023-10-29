package k8s

import (
	"dx-cli/config"
	"dx-cli/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

var verbose bool

var configK8sUser = &cobra.Command{
	Use:   "user",
	Short: "👤 Configure a Kubernetes User",
	Long:  `🎉 This command configures a user in the selected Kubernetes context.`,
	Run: func(cmd *cobra.Command, args []string) {
		k8sContext, err := utils.SelectK8sCluster()
		if err != nil {
			utils.Printf(true, "🚨 Error: %s\n", err)
			return
		}
		if k8sContext == nil {
			utils.Printf(true, "🚨 No Kubernetes context selected.\n")
			return
		}

		if err := UpdateKubeConfigWithADUser(k8sContext); err != nil {
			utils.Printf(true, "🚨 Failed to update user: %s\n", err)
			return
		}

		utils.Printf(true, "🎉 Kubernetes user for %s has been set up.\n", k8sContext.ClusterName)
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
				Env:     nil,
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
	utils.Printf(verbose, "New updated data %s", updatedData)
	if err != nil {
		utils.Printf(true, "Error creating yaml %s", err)
		return err
	}
	err = os.WriteFile(kubeConfigFile, updatedData, 0644)
	if err != nil {
		utils.Printf(true, "Error writing config file %s", err)

		return err
	}

	return nil
}

func init() {
	configK8sUser.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	KubeCmd.AddCommand(configK8sUser)
}
