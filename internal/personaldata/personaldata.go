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
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f кг.\n", p.Weight)
	fmt.Printf("Рост: %.2f м.\n\n", p.Height)

}
