package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "didactic-container",
	Short: "Repository for learning how containers works in Go.",
	Long:  `The objective of the project is to create an environment/infrastructure like a Docker to understand concepts and methods such as the creation of containers and their configuration, management and download of images, among others. This repository is learning purpose only, any contribution is welcome.`,
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
