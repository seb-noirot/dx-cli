package k8s

import (
	"dx-cli/config"
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var configK8sContext = &cobra.Command{
	Use:   "config",
	Short: "🔗 Link Up Your Cluster!",
	Long:  `🚀 Easily configure your Kubernetes cluster for seamless operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		configureCluster()
	},
}

func configureCluster() {
	// Select Kubernetes context
	k8sContext, err := utils.SelectK8sCluster()
	if err != nil {
		utils.Printf(true, "🚨 Error selecting Kubernetes cluster: %s\n", err)
		return
	}
	if k8sContext == nil {
		utils.Printf(true, "🛑 Operation aborted.\n")
		return
	}

	if err := executeClusterCommands(k8sContext); err != nil {
		utils.Printf(true, "🚨 %s\n", err.Error())
	}
}

func executeClusterCommands(k8sContext *config.KubernetesContext) error { // Replace YourK8sContextType with the actual type
	clusterName := k8sContext.ClusterName
	certificate := k8sContext.Certificate
	adUser := k8sContext.ADUser

	// Create certificate
	tempFilePath, err := createTempCertFile(clusterName, certificate)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %s", err)
	}

	if err := runCommand("kubectl config set-cluster " + clusterName + " --server=https://api." + clusterName + " --certificate-authority=" + tempFilePath + " --embed-certs=true"); err != nil {
		return fmt.Errorf("failed to set cluster: %s", err)
	}

	if err := runCommand("kubectl config set-credentials " + adUser); err != nil {
		return fmt.Errorf("failed to set credentials: %s", err)
	}

	if err := runCommand("kubectl config set-context " + clusterName + " --cluster=" + clusterName + " --namespace=default --user=" + adUser); err != nil {
		return fmt.Errorf("failed to set context: %s", err)
	}

	if err := runCommand("kubectx " + clusterName); err != nil {
		return fmt.Errorf("failed to change context: %s", err)
	}

	utils.Printf(true, "🎉 Kubernetes context for %s has been set up.\n", clusterName)
	return nil
}

func runCommand(command string) error {
	_, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		utils.Printf(true, "🚨 Failed to execute: %s\n", command)
		utils.Printf(true, "🚀 Please run it manually.\n")
	}
	return err
}

func createTempCertFile(cluster string, certContent string) (string, error) {
	// Create a temporary file
	path := "/tmp/" + cluster + ".crt"
	err := os.WriteFile(path, []byte(certContent), 0644)
	if err != nil {
		return "", err
	}
	return path, nil
}

func init() {
	KubeCmd.AddCommand(configK8sContext)
}
