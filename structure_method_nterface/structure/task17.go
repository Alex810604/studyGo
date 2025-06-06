// 17. Интерфейс для сортировки по разным полям
// Задача: //Создайте структуру Product (Name, Price, Rating). Реализуйте интерфейс sort.Interface для
// сортировки слайса []Product по любому из полей (выбор поля — через вложенную структуру-компаратор).
package structure

type Product struct {
	Name   string
	Price  float64
	Rating float64
}

type Comparator struct {
	Product  []Product
	LessFunc func(p1, p2 Product) bool // Функция сравнения двух товаров
}

func (c Comparator) Len() int {
	return len(c.Product) // Возвращает количество элементов в слайсе
}
func (c Comparator) Less(i, j int) bool {
	return c.LessFunc(c.Product[i], c.Product[j]) // Использует внешнюю LessFunc для сравнения
}
func (c Comparator) Swap(i, j int) {
	c.Product[i], c.Product[j] = c.Product[j], c.Product[i] // Меняет элементы местами
}
