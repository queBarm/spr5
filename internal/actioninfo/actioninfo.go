package actioninfo

import "fmt"

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		// Парсим строку
		if err := dp.Parse(data); err != nil {
			fmt.Printf("Error parsing string [%s]: %v\n", data, err)
			continue
		}

		// Получаем информацию о действии
		info, err := dp.ActionInfo()
		if err != nil {
			fmt.Printf("Error while generating report: %v\n", err)
			continue
		}

		// Выводим результат
		fmt.Println(info)
	}
}
