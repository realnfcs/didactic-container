package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	local    string
	url      string
	imgName  string
	filename string
)

type NewCommand struct {
    image imageEngineInterface
    cmd *cobra.Command
}

func (n *NewCommand) New() *NewCommand {
    return &NewCommand{
        // TODO: imageController packager to manager image engine
        image: imageController.GetImageEngine(),
        cmd: &cobra.Command{
        	Use:   "new",
        	Short: "new command will download a image and insert in the database",
        	Long: `The new command will download the choice image and insert into
the database.
`,
            // TODO: function return a slice of strings containing the filesystem available
            ValidArgs: n.image.Filesystem(),
        	Run: func(cmd *cobra.Command, args []string) {
    
            	if len(args) > 0 {
    
        	    	fmt.Println("Checking args...")
            		err := cmd.ValidateArgs(args)

        	    	if err != nil {
        		    	log.Fatalln(err)
            		}
                }
            
               // This part will call the image engine interface and download the fs
                if local != "" {
                    err := n.image.GetLocalImage(local, imgName, filename, args...)
                    if err != nil {
                        log.Fatalln(err)
                    }

                    return
                }

                if url != "" {
                    err := n.image.DownloadImage(args...)
                    if err != nil {
                        log.Fatalln(err)
                    }

                    return
                }

                fmt.Println("error: null flags")

            },
	    },
    }
}

func init() {

    new := new(NewCommand).New()

	new.cmd.PersistentFlags().StringVarP(&url, "url", "u", "", "Indicate the image/filesystem url for download")
	new.cmd.PersistentFlags().StringVarP(&local, "local", "l", "", "Indicate the custom image/filesystem local path")
	new.cmd.PersistentFlags().StringVarP(&filename, "filename", "f", "", "Indicate the filename of the image")
	new.cmd.PersistentFlags().StringVarP(&imgName, "name", "n", "", "Indicate the name of the image")

	imageCmd.AddCommand(new.cmd)
}
