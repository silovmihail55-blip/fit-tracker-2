package actioninfo

import (
	"fmt"
	"log"
)

// DataParser — интерфейс для парсинга данных и получения информации о действии
type DataParser interface {
	Parse(string) error          // Парсит строку, возвращает ошибку при сбое
	ActionInfo() (string, error) // Возвращает информацию о действии и возможную ошибку
}

// Info обрабатывает набор строк через парсер, выводит результат при успешном парсинге
func Info(dataset []string, dp DataParser) {
	parsedOK := false

	for _, value := range dataset {
		if err := dp.Parse(value); err != nil {
			log.Println(err)
			continue
		}
		parsedOK = true
	}

	if !parsedOK {
		return
	}

	str, err := dp.ActionInfo()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Print(str)
}
