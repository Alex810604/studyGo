// 8. Интерфейс для работы с БД
// Задача: Создайте интерфейс Database с методами Save(data string) и Read() string.
// Реализуйте его в структуре MockDB, которая сохраняет данные в памяти (в слайс).
package structure

import (
	"strings"
)

type Database interface {
	Save(data string)
	Read() string
}
type MockDB struct {
	storage []string
}

func (db *MockDB) Save(data string) {
	if data != "" {
		db.storage = append(db.storage, data)
	}
}
func (db MockDB) Read() string {
	if len(db.storage) == 0 {
		return "Нет данных"
	}
	return strings.Join(db.storage, ", ") // Данные 1, Данные 2
	//return fmt.Sprintf("%v", db.storage)	// [Данные 1 Данные 2]

}
