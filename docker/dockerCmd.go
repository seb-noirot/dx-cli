package docker

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Manage Docker",
	Run: func(cmd *cobra.Command, args []string) {
		checkDocker()
	},
}

func checkDocker() {
	out, err := exec.Command("docker", "--version").Output()

	if err != nil {
		fmt.Println("Docker is not installed. Would you like to install it? (y/n)")
		fmt.Println("To install Docker manually, visit: https://docs.docker.com/get-docker/")
		return
	}

	fmt.Println("Docker installed: ", strings.TrimSpace(string(out)))

	// Check if the Docker daemon is running
	_, err = exec.Command("docker", "info").Output()
	if err != nil {
		fmt.Println("Docker daemon is not running.")

		switch runtime.GOOS {
		case "linux":
			fmt.Println("Starting Docker daemon...")
			_, err := exec.Command("sudo", "systemctl", "start", "docker").Output()
			if err != nil {
				fmt.Println("Failed to start Docker daemon: ", err)
				return
			}
			fmt.Println("Docker daemon started.")
		case "darwin":
			fmt.Println("Starting Docker Desktop...")
			_, err := exec.Command("open", "/Applications/Docker.app").Output()
			if err != nil {
				fmt.Println("Failed to start Docker Desktop: ", err)
				return
			}
			fmt.Println("Docker Desktop started.")
		default:
			fmt.Println("Please start the Docker daemon manually.")
		}
	}
}
