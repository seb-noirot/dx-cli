package docker

import (
	"dx-cli/utils" // Import your utils package
	"github.com/spf13/cobra"
	"os/exec"
	"runtime"
	"strings"
)

var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Be the Captain of Your Container Ship 🚢",
	Long: `Want to manage containers like a seasoned sailor? 🌊

With Docker, you'll be at the helm, steering your containerized applications with ease. 🚢
From installing Docker to managing your daemons, this command has you covered.

Full speed ahead, Captain! 🌟`,
	Run: func(cmd *cobra.Command, args []string) {
		checkDocker()
	},
}

func checkDocker() {
	out, err := exec.Command("docker", "--version").Output()

	if err != nil {
		if utils.PromptUser("Docker is not installed. Would you like to install it? 🐳", []string{"Yes", "No"}) == "Yes" {
			utils.Println(true, "To install Docker manually, check out 👉 https://docs.docker.com/get-docker/")
		}
		return
	}

	utils.Println(true, "🐳 Docker installed: ", strings.TrimSpace(string(out)))

	// Check if the Docker daemon is running
	_, err = exec.Command("docker", "info").Output()
	if err != nil {
		utils.Println(true, "⚠️ Docker daemon is not running.")
		handleDaemonStart()
	}
}

func handleDaemonStart() {
	switch runtime.GOOS {
	case "linux":
		utils.Println(true, "🚀 Attempting to start Docker daemon...")
		executeAndLog("sudo", []string{"systemctl", "start", "docker"}, "🚀 Docker daemon started.", "❌ Failed to start Docker daemon.")
	case "darwin":
		utils.Println(true, "🍏 Attempting to start Docker Desktop...")
		executeAndLog("open", []string{"/Applications/Docker.app"}, "🍏 Docker Desktop started.", "❌ Failed to start Docker Desktop.")
	default:
		utils.Println(true, "❌ Please start the Docker daemon manually.")
	}
}

func executeAndLog(command string, args []string, successMsg string, errorMsg string) {
	_, err := exec.Command(command, args...).Output()
	if err != nil {
		utils.Println(true, errorMsg, err)
		return
	}
	utils.Println(true, successMsg)
}
