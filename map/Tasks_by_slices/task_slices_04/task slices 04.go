// 4. Сравнение длины и capacity
// Напишите программу, которая:
// Создает срез с помощью make([]int, 3, 5),
// Добавляет элементы, пока len не превысит cap,
// Выводит len и cap после каждого добавления.
// Пример вывода:
// Исходный: len=3, cap=5
// После добавления 4: len=4, cap=5
// После добавления 6: len=6, cap=10 (удвоение capacity)
package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)

	// Заполняем существующие  значения
	for i := 0; i < len(slice); i++ {
		slice[i] = i + 1
	}
	fmt.Printf("Исходный: %v\n==================================\n", slice)

	// Добавляем элементы, пока len не превысит cap
	for i := len(slice) + 1; ; i++ {
		Cap := cap(slice)
		slice = append(slice, i)

		fmt.Printf("Добавлен элемент: %d\n", i)
		fmt.Printf("Получен срез: %v\n", slice)
		fmt.Printf("len=%d, cap=%d", len(slice), cap(slice))

		// Проверяем увеличение capacity
		if Cap != cap(slice) {
			fmt.Printf(" (произошло удвоение capacity с %d до %d)", Cap, cap(slice))
		}
		fmt.Println("\n-------------------------------------")

		// Цикл бесконечный, поэтому ограничиваем его длину (например, когда длина превысит его исходную вместимость)
		if len(slice)+1 > 6 {
			break
		}
	}
}
