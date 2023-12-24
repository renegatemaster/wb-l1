// Реализовать все возможные способы
// остановки выполнения горутины.
package main

import (
	"fmt"
	"sync"
	"time"
)

func read(ch chan int, wg *sync.WaitGroup) {

	for {
		val, ok := <-ch
		if !ok {
			wg.Done()
			fmt.Println("Exiting goroutine 'read'")
			return
		}
		fmt.Println("value: ", val)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	// Горутина для чтения
	// Завершится, когда отработает команад
	// wg.Done()
	go read(ch, &wg)
	// Горутина для записи
	// Завершится, когда будет нечего читать
	// из канала и можно будет его закрыть
	go write(ch)

	// Ждём завершения горутины
	wg.Wait()

	fmt.Println("Exiting main.")
}

func write(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}
