package main

import (
	"fmt"
)

func main() {
	// Создаем map для хранения информации о выданных изданиях
	readerPublications := map[string]map[string][]string{
		"Иванов": {
			"книги":   {"Война и мир", "Преступление и наказание"},
			"журналы": {"Наука и жизнь", "Вокруг света"},
		},
		"Петрова": {
			"книги":   {"1984", "Мастер и Маргарита"},
			"журналы": {"Forbes"},
		},
		"Сидоров": {
			"книги": {"Тихий Дон"},
		},
	}
	// Добавление новых изданий существующему читателю
	readerPublications["Иванов"]["книги"] = append(readerPublications["Иванов"]["книги"], "Анна Каренина")
	readerPublications["Петрова"]["периодика"] = append(readerPublications["Петрова"]["периодика"], "Time")
	//удаление
	delete(readerPublications["Петрова"], "книги")
	// 1. Выводим количество читателей с изданиями
	fmt.Printf("Количество читателей с изданиями: %d\n\n", len(readerPublications))

	// 2. Выводим общее количество изданий у каждого читателя
	for reader, publications := range readerPublications {
		total := 0
		// Подсчитываем общее количество изданий для каждого типа
		for _, items := range publications {
			total += len(items)
		}
		fmt.Printf("Читатель %s имеет %d изданий\n", reader, total)
	}
}
