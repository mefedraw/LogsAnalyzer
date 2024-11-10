package Input

import (
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

func (ui *UserInput) Input() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите команду: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	fmt.Println("Введенная команда:", input)
	_ = ui.ParseInput(input)
}

func (ui *UserInput) ParseInput(command string) error {
	// Убираем префикс команды "analyzer " и разбиваем оставшуюся строку на аргументы
	command = strings.TrimPrefix(command, "analyzer ")
	args := strings.Fields(command)

	// Итерируем по аргументам и парсим ключи и значения
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--path":
			if i+1 < len(args) {
				ui.Path = args[i+1]
				i++ // Переходим к следующему аргументу
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

	// Проверка обязательных параметров
	if ui.Path == "" {
		return fmt.Errorf("обязательный параметр --path отсутствует")
	}
	if ui.Format == "" {
		ui.Format = "markdown" // Устанавливаем формат по умолчанию, если не указан
	}

	return nil
}
