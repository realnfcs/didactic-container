package cmd

import (
	"fmt"

	"github.com/realnfcs/didactic-container/internal/image"
	"github.com/spf13/cobra"
)

// Var who will storage the flags values of del command
var (
	id   string
	name string
	path string
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "del command will delete the specified image",
	Long: `The del command will delete the specified image with the id, name,
or path of the image.
    `,
	Run: func(cmd *cobra.Command, args []string) {
		if id != "" || name != "" {
			image.DeleteImage(id, name, path)
		} else {
			fmt.Println("value error: you have to pass a value with flags")
			cmd.Help()
		}
	},
}

func init() {

	delCmd.PersistentFlags().StringVarP(&id, "id", "i", "", "flag to specified the image with id")
	delCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "flag to specified the image with name")
	delCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "flag to specified the image with path")

	imageCmd.AddCommand(delCmd)
}
