package docker

import (
	"dx-cli/utils" // Import your utils package
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

var DockerComposeCmd = &cobra.Command{
	Use:   "docker-compose",
	Short: "Orchestrate Your Containers Like a Maestro 🎼",
	Long: `Ever felt like you're juggling too many containers? 🤹

Don't sweat it! Docker Compose will have you orchestrating containers as if you're a maestro. 🎼
Install, configure, and manage your projects with just a few commands.

Let the concert begin! 🎶`,
	Run: func(cmd *cobra.Command, args []string) {
		checkDockerCompose()
	},
}

func checkDockerCompose() {
	out, err := exec.Command("docker-compose", "--version").Output()

	if err != nil {
		if utils.PromptUser("Docker Compose is not installed. Would you like to install it? 🎵", []string{"Yes", "No"}) == "Yes" {
			utils.Println(true, "🛠 To install Docker Compose manually, check out 👉 https://docs.docker.com/compose/install/")
		}
		return
	}

	utils.Println(true, "🎵 Docker Compose installed: ", strings.TrimSpace(string(out)))

	// Add more functionality here as needed
}
