package Input

import (
	"NginxLogsAnalyzer/Errors/UserInputError"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserInput struct {
	Path        string
	FromDate    string
	ToDate      string
	Filter      string
	FilterValue string
	Format      string
}

func NewUserInput() *UserInput {
	return &UserInput{}
}

func (ui *UserInput) Input() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите команду: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	err := ui.ParseInput(input)
	return err
}

func (ui *UserInput) ParseInput(command string) error {
	if strings.Split(command, " ")[0] != "analyzer" {
		return UserInputError.NewErrUserInput("incorrect command")
	}
	command = strings.TrimPrefix(command, "analyzer ")
	args := strings.Fields(command)

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--path":
			if i+1 < len(args) {
				ui.Path = args[i+1]
				i++
			}
		case "--from":
			if i+1 < len(args) {
				ui.FromDate = args[i+1]
				i++
			}
		case "--to":
			if i+1 < len(args) {
				ui.ToDate = args[i+1]
				i++
			}
		case "--filter-field":
			if i+1 < len(args) {
				ui.Filter = args[i+1]
				i++
			}
		case "--filter-value":
			if i+1 < len(args) {
				ui.FilterValue = args[i+1]
				i++
			}
		case "--format":
			if i+1 < len(args) {
				ui.Format = args[i+1]
				i++
			}
		}
	}

	if ui.Path == "" {
		return UserInputError.NewErrUserInput("--path does not exist")
	}
	if ui.Format == "" {
		ui.Format = "markdown"
	}

	return nil
}
