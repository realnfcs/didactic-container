package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/realnfcs/didactic-container/internal/image"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		newImage()
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	imageCmd.AddCommand(newCmd)
}

// Funtion to get input from user
func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}

		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

// Function to specified the prompt select UI
func promptGetSelect(pc promptContent) (result string) {
	items := []string{"Alpine", "Ubuntu", "Cancel"}
	index := -1

	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.label,
			Items: items,
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
	}

	return
}

// Function to download a image if them is not on workspace/fs
func newImage() {
	imagePromptContent := promptContent{
		"Please select a filesystem",
		"What of these filesystem you want?",
	}

	result := promptGetSelect(imagePromptContent)

	switch result {

	case "ubuntu":
		image.UbuntuImage()
		break

	case "alpine":
		image.AlpineImage()
		break

	default:
		os.Exit(1)
	}

}
