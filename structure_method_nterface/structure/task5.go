// 5. Интерфейс Stringer
// Задача: Реализуйте интерфейс fmt.Stringer для структуры Book (с полями Title, Author), чтобы при печати выводилось:
// "Книга: {Title}, Автор: {Author}".
package structure

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Книга: {%s}, Автор: {%s}", b.Title, b.Author)
}
func PrintBook(s fmt.Stringer) {
	fmt.Println(s)
}
