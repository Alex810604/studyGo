// 6. Разделение слайса на пакеты (чанки)
// Напишите функцию ChunkSlice, которая принимает слайс целых чисел и размер чанка,
// а возвращает слайс слайсов, разбитый на части указанного размера.
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunkSize := 3

	chunks := ChunkSlice(numbers, chunkSize)
	fmt.Println(chunks) // [[1 2 3] [4 5 6] [7 8 9] [10]]
}
func ChunkSlice[T any](numbers []T, size int) [][]T {
	var slices [][]T
	for i := 0; i < len(numbers); i += size {
		end := i + size
		if end > len(numbers) {
			end = len(numbers)
		}
		slices = append(slices, numbers[i:end])
	}
	return slices
}
