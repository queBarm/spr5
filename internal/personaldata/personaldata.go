package personaldata

import "fmt"

type Personal struct {
	//Структура данных пользователя
	Name   string
	Weight float64
	Height float64
}

func (p Personal) Print() {
	//Выводим данные пользователя
	info := (fmt.Sprintf("Имя: %s\n", p.Name) +
		fmt.Sprintf("Вес: %.2f кг.\n", p.Weight) +
		fmt.Sprintf("Рост: %.2f м.\n\n", p.Height))

	fmt.Print(info)
}
