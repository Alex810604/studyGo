// 16. Работа с временем в структуре
// Задача: Создайте структуру Event с полями Title и Time (типа time.Time).
// Добавьте метод IsAfter(now time.Time) bool, проверяющий, что событие еще не наступило.
package structure

import "time"

type Event struct {
	Title string
	Time  time.Time
}

func (ev *Event) IsAfter(now time.Time) bool {
	return ev.Time.After(now)

}
