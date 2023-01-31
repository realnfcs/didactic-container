package components

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/manifoldco/promptui"
)

type PromptContent struct {
	ErrorMsg string
	Label    string
}

// Funtion to get input from user
func (pc *PromptContent) PromptGetInput() string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.ErrorMsg)
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
		Label:     pc.Label,
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
func (pc *PromptContent)  PromptGetSelect() (result string) {
	items := []string{"Alpine", "Ubuntu", "Cancel"}
	index := -1

	var err error

	for index < 0 {
		prompt := promptui.Select{
			Label: pc.Label,
			Items: items,
		}

		index, result, err = prompt.Run()
	}

	if err != nil {
		log.Printf("Prompt failed %v\n", err)
	}

	return
}
