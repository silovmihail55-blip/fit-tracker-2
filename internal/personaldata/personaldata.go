package personaldata

import "fmt"

// Personal хранит личные данные пользователя: имя, вес и рост
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит на экран личные данные в формате:
// Имя: <имя>
// Вес: <вес> кг.
// Рост: <рост> м.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
