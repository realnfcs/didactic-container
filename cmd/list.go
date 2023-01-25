/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/realnfcs/didactic-container/internal/database"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "See list of all images you have download",
	Long:  `Param to see a list of all images you have download.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.InfoImages()
	},
}

func init() {
	imageCmd.AddCommand(listCmd)
}
