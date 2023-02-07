package cmd

import (
	"fmt"
	"log"

	"github.com/realnfcs/didactic-container/internal/image"
	"github.com/spf13/cobra"
)

var (
	local    string
	url      string
	imgName  string
	filename string
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "new command will download a image and insert in the database",
	Long: `The new command will download the choice image and insert into
the database.
`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			fmt.Println("Checking args...")
			err := cmd.ValidateArgs(args)
			if err != nil {
				log.Fatalln(err)
			}

			for _, v := range args {
				switch v {
				case "ubuntu":
					image.UbuntuImage()
				case "alpine":
					image.AlpineImage()
				}
			}

			return
		}

		if url != "" {
			if filename == "" || imgName == "" {
				cmd.Help()
				log.Fatalln("error: you have the specified the filename and the image name")
			}

			fs := image.Filesystem{
				URL:      url,
				FileName: filename,
				Name:     imgName,
			}

			err := fs.PullImage()
			if err != nil {
				log.Fatalln(err)
			}

			return
		}

		if local != "" {
			if filename == "" || imgName == "" {
				cmd.Help()
				log.Fatalln("error: you have the specified the filename and the image name")
			}

			fs := image.Filesystem{
				URL:      local,
				FileName: filename,
				Name:     imgName,
			}

			err := fs.PullLocalImage()
			if err != nil {
				log.Fatalln(err)
			}

			return
		}
	},
}

func init() {

	newCmd.ValidArgs = []string{"ubuntu", "alpine"}

	newCmd.PersistentFlags().StringVarP(&url, "url", "u", "", "Indicate the image/filesystem url for download")
	newCmd.PersistentFlags().StringVarP(&local, "local", "l", "", "Indicate the custom image/filesystem local path")
	newCmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "Indicate the filename of the image")
	newCmd.PersistentFlags().StringVarP(&imgName, "name", "n", "", "Indicate the name of the image")

	imageCmd.AddCommand(newCmd)
}
