// 15. Интерфейс для клонирования объектов
// Задача: Создайте интерфейс Cloner с методом Clone() Cloner.
// Реализуйте его для структуры Document (поле Content string), чтобы метод возвращал копию объекта.
package structure

type Cloner interface {
	Clone() Cloner
}
type Document struct {
	Content string
}

func (d *Document) Clone() Cloner {
	return &Document{Content: d.Content}
}
