// 4. Thread-safe булев флаг
// Реализуйте булев флаг, который можно безопасно устанавливать и проверять из разных горутин, используя только atomic.
package goroutine2

import "sync/atomic"

type Flag struct {
	state int64
}

func (f *Flag) Set(value bool) {
	var i int64
	if value {
		i = 1
	}
	atomic.StoreInt64(&f.state, i)
}

func (f *Flag) Get() bool {
	return (atomic.LoadInt64(&f.state)) != 0

}
