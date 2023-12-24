// Реализовать структуру-счетчик,
// которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Структура счётчика
type Counter struct {
	value int
	mu    sync.Mutex
}

// Увеличиваем счётчик на 1
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Выводим значение счётчика
func (c *Counter) Result() int {
	return c.value
}

func main() {

	wg := sync.WaitGroup{}
	counter := Counter{value: 0}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Запускаем 10 горутин
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// Каждая по очереди увеличивает счётчик до завершения контекста
		go func(ctx context.Context, counter *Counter) {
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					counter.Increment()
				}
			}
		}(ctx, &counter)
	}

	wg.Wait()
	fmt.Println(counter.Result())
}
