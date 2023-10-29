package utils

import (
	"dx-cli/config"
	"fmt"
)

func SelectK8sCluster() (*config.KubernetesContext, error) {
	// Fetch the current context
	currentContext, err := GetCurrentContext(config.ConfigFilePath, false)
	if err != nil {
		fmt.Printf("Error fetching current context: %s\n", err)
		return nil, nil
	}

	if currentContext == nil {
		fmt.Println("No current context defined.")
		return nil, nil
	}

	if len(currentContext.KubernetesContexts) == 0 {
		fmt.Println("No Kubernetes context available.")
		return nil, nil
	}

	fmt.Println("Select a Kubernetes context:")
	for i, k8sContext := range currentContext.KubernetesContexts {
		fmt.Printf("[%d] %s\n", i+1, k8sContext.ClusterName)
	}

	var choice int
	fmt.Scanln(&choice)

	if choice < 1 || choice > len(currentContext.KubernetesContexts) {
		fmt.Println("Invalid choice.")
		return nil, nil
	}
	selectedK8sContext := currentContext.KubernetesContexts[choice-1]
	return &selectedK8sContext, nil
}
