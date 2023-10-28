package k8s

import (
	"dx-cli/utils"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
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
		tempFilePath, err := CreateTempCertFile(certificate)
		if err != nil {
			log.Fatalf("Failed to create temp file: %s", err)
		}
		// Set Cluster Configuration
		exec.Command("kubectl", "config", "set-cluster", clusterName, "--server=https://api."+clusterName, "--certificate-authority="+tempFilePath, "--embed-certs=true").Run()

		// Create Cluster User
		exec.Command("kubectl", "config", "set-credentials", adUser).Run()

		// Create Context
		exec.Command("kubectl", "config", "set-context", clusterName, "--cluster="+clusterName, "--namespace=default", "--user="+adUser).Run()

		// Change Context
		exec.Command("kubectx", clusterName).Run()

		fmt.Printf("Kubernetes context for %s has been set up.\n", clusterName)

	},
}

func CreateTempCertFile(certContent string) (string, error) {
	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "cert")
	if err != nil {
		return "", err
	}

	// Ensure the file will be deleted after closing
	defer os.Remove(tempFile.Name())

	// Write the certificate content into the temp file
	_, err = tempFile.WriteString(certContent)
	if err != nil {
		return "", err
	}

	// Close the file
	if err := tempFile.Close(); err != nil {
		return "", err
	}

	// Return the name of the temporary file
	return tempFile.Name(), nil
}

func init() {
	KubeCmd.AddCommand(configK8sContext)
}
