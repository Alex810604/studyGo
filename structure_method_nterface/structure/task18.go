// 18. Динамическая диспетчеризация методов
// Задача: Создайте интерфейс Transport с методом Move(). Реализуйте его для структур Car, Bicycle и Plane.
// Напишите функцию StartRace(transports []Transport), которая вызывает Move() для всех участников.
package structure

import "fmt"

type Transport interface {
	Move()
}

// Согласно условию задачи
type Car struct{}
type Bicycle struct{}
type Plane struct{}

func (Car) Move()        { fmt.Println("движется автомобиль") }
func (bc Bicycle) Move() { fmt.Println("едет велосипед") }
func (pl Plane) Move()   { fmt.Println("летит самолет") }
func StartRace(transports []Transport) {
	for _, transport := range transports {
		transport.Move()
	}
}

// Вариант с добавлением полей в структуры для наглядности и с добавления возвраного значения в метод интерфейса
type Transport1 interface {
	Move1() string
}
type Car1 struct {
	Model string //для наглядности добавим поля
}
type Bicycle1 struct {
	Brand string //для наглядности добавим поля
}
type Plane1 struct {
	FlightNumber string //для наглядности добавим поля
}

func (car Car1) Move1() string {
	return fmt.Sprintf("движется автомобиль %s", car.Model)
}
func (bc Bicycle1) Move1() string {
	return fmt.Sprintf("едет велосипед %s", bc.Brand)
}

func (pl Plane1) Move1() string {
	return fmt.Sprintf("летит самолет %s", pl.FlightNumber)
}
func StartRace1(transports []Transport1) {
	for _, transport := range transports {
		fmt.Println(transport.Move1())
	}

}
