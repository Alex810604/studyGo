// 10. Уникальные элементы
// Создайте программу, которая принимает массив и возвращает новый массив, содержащий только уникальные элементы.
package main

import "fmt"

func main() {
	var num []int
	var n int       //число в массиве
	var numbers int //количесло чисел в массиве
	fmt.Print("Введите количество чисел в массиве: ")
	fmt.Scan(&numbers)
	for i := 0; i < numbers; i++ {
		fmt.Printf("Введите число №%d: ", i+1)
		fmt.Scanln(&n)       //создаем скан и вводим числа в массиве
		num = append(num, n) // добавляем сканированные числа в срез num
	}
	fmt.Printf("Массив, содержащий только уникальные элементы от массива '%d': ", num)

	// для решения 1-м способом
	fmt.Println(cherezFor(num))

	// для решения 2-м способом
	findUnique(num)

	// Решение с удалением дубликатов в массиве
	fmt.Printf("Массив без дублирующих чисел в отличие от массива'%d': %d.\n", num, dublikat(num))
}

// Первый способ решения
func cherezFor(numbers1 []int) []int {
	var unikMassiv []int // создаем новый как бы уникальный массив
	for i, chislo1 := range numbers1 {
		unikChiclo := true // допустим это число уникальное в массиве numbers1 []int
		for j, chislo2 := range numbers1 {
			if i != j && chislo1 == chislo2 {
				unikChiclo = false
				break
			}
		}
		if unikChiclo == true {
			unikMassiv = append(unikMassiv, chislo1)
		}
	}
	return unikMassiv
}

// Второй способ решения
func findUnique(numbers2 []int) {
	freq := make(map[int]int)
	for _, num := range numbers2 {
		freq[num]++
	}
	var unique []int
	for _, num := range numbers2 {
		if freq[num] == 1 {
			unique = append(unique, num)
		}
	}
	fmt.Println(unique)
}

// Трейтий способ решения - устранение чисел-дубликатов
func dublikat(numbers3 []int) []int {
	var unMassiv []int
	for _, num1 := range numbers3 {
		chiclo := false
		for _, num2 := range unMassiv {
			if num1 == num2 {
				chiclo = true
				break
			}
		}
		if !chiclo {
			unMassiv = append(unMassiv, num1)
		}
	}
	return unMassiv
}
