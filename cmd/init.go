/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/realnfcs/didactic-container/internal"
	"github.com/realnfcs/didactic-container/internal/database"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise the whole project",
	Long: `The project have to initialise the database, folders and the workspaces at all. 
	For that, is necessary to run the init command first to have the project ready to use.`,
	Run: func(cmd *cobra.Command, args []string) {
		internal.InitFolders()
		database.CreateTable()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
