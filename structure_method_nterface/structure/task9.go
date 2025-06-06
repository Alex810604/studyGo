// 9. Полиморфизм с интерфейсами
// Задача: Создайте интерфейс Animal с методом Sound() string. Реализуйте его для структур Dog и Cat.
// Затем создайте функцию MakeSound(animal Animal), которая вызывает Sound().
package structure

import "fmt"

type Animal interface {
	Sound() string //звук
}
type Dog struct{}

func (d Dog) Sound() string {
	return "гав"
}

type Cat struct{}

func (c Cat) Sound() string {
	return "мяу"
}
func MakeSound(animal Animal) {
	fmt.Println(animal.Sound())
}
