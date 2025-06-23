// 1. Кэш с RW Mutex
// Реализуйте потокобезопасный кэш, где операции чтения происходят значительно чаще, чем записи.
// Используйте sync.RWMutex для оптимизации производительности.
package goroutine2

import (
	"fmt"
	"sync"
	"time"
)

var (
	cache  = make(map[string]string)
	rwLock sync.RWMutex
	key    = ""
)

// Чтение (многопоточное)
func Get(key string) string {
	rwLock.RLock() // Блокировка для чтения (не эксклюзивная)
	defer rwLock.RUnlock()
	return cache[key]
}

func Writer(wg *sync.WaitGroup, i int, word string) {
	defer wg.Done()
	key = fmt.Sprintf("писателя №%d", i+1)
	rwLock.Lock() // Полная блокировка
	cache[key] = word
	rwLock.Unlock()
	fmt.Printf("Текст %s -> '%s'\n", key, word)
	time.Sleep(500 * time.Millisecond) // Имитируем редкие записи
}

func Reader(wg *sync.WaitGroup, j int) {
	defer wg.Done()
	value := Get(key)
	fmt.Printf("Читатель №%d читает %s -> '%s'\n", j, key, value)
	// Имитация чтения
	time.Sleep(100 * time.Millisecond)

}
