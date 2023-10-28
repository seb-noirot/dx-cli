package gitlab

import (
	"bufio"
	"dx-cli/config" // replace with your actual config package path
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var registerSSHKeyCmd = &cobra.Command{
	Use:   "register",
	Short: "Register an SSH key with a GitLab account",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentContext, err := config.GetCurrentContext()
		if err != nil {
			return err
		}

		// List available GitLab definitions and prompt for selection
		for i, glContext := range currentContext.GitLabContexts {
			fmt.Printf("%d: %s (%s)\n", i+1, glContext.Name, glContext.Host)
		}

		fmt.Println("Select a GitLab definition:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		selection := scanner.Text()
		selectedIndex, err := strconv.Atoi(selection)
		if err != nil || selectedIndex < 1 || selectedIndex > len(currentContext.GitLabContexts) {
			return fmt.Errorf("Invalid selection")
		}

		selectedGLContext := currentContext.GitLabContexts[selectedIndex-1]

		// Check if key exists
		privateKeyPath := filepath.Join(os.Getenv("HOME"), ".ssh", fmt.Sprintf("id_%s_%s", currentContext.Name, selectedGLContext.Name))
		if _, err := os.Stat(privateKeyPath); os.IsNotExist(err) {
			return fmt.Errorf("SSH key does not exist. Please create one first.")
		}

		// Open GitLab SSH key page
		cleanHost := strings.TrimRight(selectedGLContext.Host, "/")
		gitLabURL := fmt.Sprintf("%s/-/profile/keys", cleanHost)
		err = exec.Command("open", gitLabURL).Run()
		if err != nil {
			return fmt.Errorf("Could not open GitLab URL: %s", err)
		}

		// Display public key content to paste
		publicKeyPath := privateKeyPath + ".pub"
		publicKeyContent, err := os.ReadFile(publicKeyPath)
		if err != nil {
			return fmt.Errorf("Could not read public key: %s", err)
		}
		fmt.Println("Paste the following public key into GitLab:")
		fmt.Println(string(publicKeyContent))

		return nil
	},
}

func TestSSHConnection(host string) error {
	cmd := exec.Command("ssh", "-T", fmt.Sprintf("git@%s", host))
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("SSH connection failed: %s", err)
	}

	fmt.Println("SSH connection successful. Output:", string(output))
	return nil
}

func init() {
	GitlabCmd.AddCommand(registerSSHKeyCmd)
}
