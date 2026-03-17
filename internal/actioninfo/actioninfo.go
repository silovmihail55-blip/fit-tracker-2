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
	for _, value := range dataset {
		if err := dp.Parse(value); err != nil {
			log.Println(err)
			continue
		}

		// Переместили сюда обработку ActionInfo прямо после успешного парсинга
		str, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Print(str)
	}
}
