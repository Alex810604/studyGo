//2. Интерфейс "Фигура"
//Задача: Создайте интерфейс Shape с методом Area() float64.
//Реализуйте этот интерфейс для структур Circle (с полем Radius) и Rectangle (из первой задачи).

package structure

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct { // круг
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height

}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
func Figure(s Shape) {
	fmt.Println("Площадь:", s.Area())
}

//func main() {
//	r := Rectangle{5, 4.36}
//	c := Circle{5}
//	Figure(r)
//	Figure(c)
//}
