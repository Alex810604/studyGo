// 6. Пустой интерфейс и type assertion
// Задача: Напишите функцию PrintType, которая принимает пустой интерфейс (interface{}),
// определяет тип переданного значения и печатает его.
package structure

import "fmt"

func PrintType(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("string", v)
	case int:

		fmt.Println("int", v)
	case float64:

		fmt.Println("float64", v)
	case bool:
		fmt.Println("bool", v)
	case nil:
		fmt.Println("nil", v)
	case []int:
		fmt.Println("[]int{}", v, len(v))
	default:
		fmt.Println("Неизвестный тип:", v)

	}
}
