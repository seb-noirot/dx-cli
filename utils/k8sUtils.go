package utils

import (
	"dx-cli/config"
	"fmt"
)

func SelectK8sCluster() (*config.KubernetesContext, error) {
	// Fetch the current context
	path, err := config.GetConfigFilePath()
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	currentContext, err := GetCurrentContext(path, false)
	if err != nil {
		Printf(true, "🚨 Error fetching current context: %s\n", err)
		return nil, err
	}

	if currentContext == nil {
		Println(true, "🛑 No current context defined.")
		return nil, nil
	}

	if len(currentContext.KubernetesContexts) == 0 {
		Println(true, "🛑 No Kubernetes context available.")
		return nil, nil
	}

	Println(true, "🎯 Select a Kubernetes context:")

	options := make([]string, len(currentContext.KubernetesContexts))
	for i, k8sContext := range currentContext.KubernetesContexts {
		options[i] = k8sContext.ClusterName
	}

	choice := PromptUser("⚓ Choose a Kubernetes context:", options)

	if choice == "" {
		Println(true, "🛑 Invalid choice.")
		return nil, nil
	}

	// Find the selected context
	var selectedK8sContext *config.KubernetesContext
	for _, k8sContext := range currentContext.KubernetesContexts {
		if k8sContext.ClusterName == choice {
			selectedK8sContext = &k8sContext
			break
		}
	}

	return selectedK8sContext, nil
}
