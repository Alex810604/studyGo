// 14. Метод с variadic-аргументами
// Задача:
// Создайте структуру Calculator с методом Sum(numbers ...int) int, который возвращает сумму переданных чисел.
package structure

type Calculator struct{}

func (cal *Calculator) Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
