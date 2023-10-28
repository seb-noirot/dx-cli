package k8s

import (
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var configK8sContext = &cobra.Command{
	Use:   "config",
	Short: "config a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		// Select Kubernetes context
		k8sContext, err := utils.SelectK8sCluster()
		if err != nil {
			return
		}
		if k8sContext == nil {
			return
		}

		clusterName := k8sContext.ClusterName
		certificate := k8sContext.Certificate
		adUser := k8sContext.ADUser

		// Create certificate
		tempFilePath, err := CreateTempCertFile(clusterName, certificate)
		if err != nil {
			log.Fatalf("Failed to create temp file: %s", err)
			return
		}

		command := "kubectl config set-cluster " + clusterName + " --server=https://api." + clusterName + " --certificate-authority=" + tempFilePath + " --embed-certs=true"
		// Set Cluster Configuration
		_, err = exec.Command("bash", "-c", command).Output()
		if err != nil {
			fmt.Println("Failed to set cluster: ", command, err)
			return
		}
		commandSetCreds := "kubectl config set-credentials " + adUser
		// Create Cluster User
		_, err = exec.Command("bash", "-c", commandSetCreds).Output()
		if err != nil {
			fmt.Println("Failed to set creds: ", commandSetCreds, err)
			return
		}

		// Create Context
		commandSetContext := "kubectl config set-context " + clusterName + " --cluster=" + clusterName + " --namespace=default --user=" + adUser
		_, err = exec.Command("bash", "-c", commandSetContext).Output()
		if err != nil {
			fmt.Println("Failed to set context: ", commandSetContext, err)
			return
		}
		// Change Context
		commandChangeContext := "kubectx " + clusterName
		_, err = exec.Command("bash", "-c", commandChangeContext).Output()
		if err != nil {
			fmt.Println("Failed to changeContext context: ", commandChangeContext, err)
			return
		}

		fmt.Printf("Kubernetes context for %s has been set up.\n", clusterName)

	},
}

func CreateTempCertFile(cluster string, certContent string) (string, error) {
	// Create a temporary file
	path := "/tmp/" + cluster + ".crt"
	err := os.WriteFile(path, []byte(certContent), 0644)
	if err != nil {
		return "", err
	}

	// Return the name of the temporary file
	return path, nil
}

func init() {
	KubeCmd.AddCommand(configK8sContext)
}
