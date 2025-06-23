package main

import (
	"fmt"
	"time"
)

//
//type Animal1 interface {
//	Move()  //бежит
//	Speak() // говорит
//}
//
//type Dog1 struct {
//	Name1 string
//}
//
//type Cat1 struct {
//	Name1 string
//}
//
//func (d Dog1) Move() {
//	fmt.Println("бежит " + d.Name1)
//}
//
//func (d Dog1) Speak() { fmt.Println("говорит " + d.Name1) } // гавкает
//
//func (c Cat1) Move() { // сигнатура
//	fmt.Println("бежит " + c.Name1) // реализация
//}
//
//func (c Cat1) Speak() { fmt.Println("говорит " + c.Name1) }
//
//func Animal2(a Animal1) { // сигнатура
//
//	a.Speak() // реализация
//	a.Move()

func main() {
	//dog := Dog1{Name1: "Барсик"}
	//cat := Cat1{Name1: "Мурзик"}
	//Animal2(dog)
	//Animal2(cat)
	// := animal.Movement()
	//dogAnimal := animal.Dog{Name: "Барсик"}
	//catAnimal := animal.Cat{Name: "Мурзик"}
	//loudCat := animal.LoudCat{Cat: animal.Cat{Name: "Васька"}}
	//animal.MoveAnimal(dogAnimal)
	//animal.SpeakAnimal(dogAnimal)
	//animal.SpeakAnimal(catAnimal)
	//animal.SpeakAnimal(loudCat)

	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("Горутина завершилась") // не выполнится, если main() завершится раньше
	//}()
	//time.Sleep(1 * time.Second)
	// программа завершится через 1 секунду, не дожидаясь горутины
	go func() {
		for {
			fmt.Println("Бесконечный цикл...")
			time.Sleep(1 * time.Second)
		}
	}()

	// Ждём Ctrl+C...
	select {}

}
