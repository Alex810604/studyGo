package animal

import "fmt"

// Принцип разделения интерфейсов. "Лучше много узких интерфейсов, чем один универсальный."
// "Клиенты" не должны зависеть от методов, которые они не используют.
type Movement interface{ Move() }  // Только движение
type Speaking interface{ Speak() } //	Только звуки

// Каждый класс/структура должна иметь только одну причину для изменения (одну ответственность)
// - принцип единственной ответственности
type Dog struct{ Name string }

func (d Dog) Move()  { fmt.Println("бежит " + d.Name) }
func (d Dog) Speak() { fmt.Println("громко лающий " + d.Name) }

type Cat struct{ Name string }

func (c Cat) Speak() { fmt.Println("мяукает " + c.Name) }

type LoudCat struct{ Cat } // Добавление новых животных через встраивание
func (l LoudCat) Speak()   { fmt.Println("ГРОМКО мяукает", l.Name) }

func MoveAnimal(m Movement) { // сигнатура
	m.Move() // реализация
}
func SpeakAnimal(s Speaking) { // сигнатура
	s.Speak() // реализация
}
