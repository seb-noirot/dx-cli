/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"dx-cli/context"
	"dx-cli/docker"
	"dx-cli/git"
	"dx-cli/gitlab"
	"dx-cli/ide"
	"dx-cli/k8s"
	"dx-cli/maven"
	"dx-cli/sdkman"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dx-cli",
	Short: "Welcome to DevBot 🤖: Your Developer Experience Enhancer!",
	Long: `Hey there, Developer! 👋

Meet DevBot 🤖, designed to elevate your developer experience with us!
From setting up essential tools 🛠️ to configuring your workspace 🌌,
DevBot has got your back.

Follow along and let's get you up and running in no time! 🚀`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dx-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(context.ContextCmd)
	rootCmd.AddCommand(docker.DockerCmd)
	rootCmd.AddCommand(sdkman.SdkmanCmd)
	rootCmd.AddCommand(gitlab.GitlabCmd)
	rootCmd.AddCommand(git.GitCmd)
	rootCmd.AddCommand(k8s.KubeCmd)
	rootCmd.AddCommand(maven.MavenCmd)
	rootCmd.AddCommand(ide.IdeCmd)
	rootCmd.AddCommand(docker.DockerComposeCmd)
}
