//3. Метод с получателем по значению и указателю
//Задача: Создайте структуру Counter с полем Value. Добавьте два метода:
//Increment() (увеличивает Value на 1, получатель по указателю)
//GetValue() int (возвращает Value, получатель по значению)

package structure

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value += 1
}

func (c Counter) GetValue() int {
	return c.Value
}
