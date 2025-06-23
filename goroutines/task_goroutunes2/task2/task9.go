// 9. Хранилище временных меток
// Создайте потокобезопасное хранилище последних временных меток событий, где:
// Запись происходит редко (обновление метки)
// Чтение происходит часто
// Используйте sync.RWMutex.
package goroutine2

import (
	"sync"
	"time"
)

type Events struct {
	storage map[string]time.Time
	rwmu    sync.RWMutex
}

// создаём пустое хранилище
func NewEvents() *Events {
	return &Events{storage: make(map[string]time.Time)}
}
func (e *Events) SetEvents(key string) {
	e.rwmu.Lock()
	defer e.rwmu.Unlock()
	e.storage[key] = time.Now()
}
func (e *Events) GetEvents(key string) time.Time {
	e.rwmu.RLock()
	defer e.rwmu.RUnlock()
	return e.storage[key]
}
