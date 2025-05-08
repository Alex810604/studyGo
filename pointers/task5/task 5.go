// Дана структура:
//type Person struct {
//    Name string
//    Age  int
//}
//Напишите функцию Birthday(p *Person), которая увеличивает возраст (Age) на 1.

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alice", Age: 25}
	Birthday(&person)
	fmt.Println(person)
	//или
	fmt.Println(person.Name)
	//или
	fmt.Println(person.Age)
}
func Birthday(p *Person) {
	p.Age += 1

}
