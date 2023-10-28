package docker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var DockerComposeCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "Manage Docker Compose",
	Run: func(cmd *cobra.Command, args []string) {
		checkDockerCompose()
	},
}

func checkDockerCompose() {
	out, err := exec.Command("docker-compose", "--version").Output()

	if err != nil {
		fmt.Println("Docker Compose is not installed. Would you like to install it? (y/n)")
		fmt.Println("To install Docker Compose manually, visit: https://docs.docker.com/compose/install/")
		return
	}

	fmt.Println("Docker Compose installed: ", strings.TrimSpace(string(out)))

	// You could add more functionality here depending on what you'd like to manage with Docker Compose
}
