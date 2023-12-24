// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

func main() {

	data := make(map[int]int)
	// Мьютекс, чтобы избежать datarace'a
	var mu sync.Mutex
	// WaitGroup, чтобы дождаться всех горутин
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(data map[int]int, mu *sync.Mutex, key int) {
			mu.Lock()
			data[key]++
			fmt.Println(data)
			mu.Unlock()
			wg.Done()
		}(data, &mu, i)
	}

	wg.Wait()
	fmt.Println(data)
}
