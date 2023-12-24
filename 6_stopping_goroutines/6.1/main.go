// Реализовать все возможные способы
// остановки выполнения горутины.
package main

import (
	"fmt"
	"time"
)

func myGoroutine(stopCh chan int) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Exiting goroutine.")
			return
		default:
			fmt.Println("Doing some work...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	stopCh := make(chan int)

	go myGoroutine(stopCh)

	// Имитируем работу main, чтобы дать время myGoroutine
	time.Sleep(3 * time.Second)

	// Останавливаем сигналом в канал
	stopCh <- 1

	// Также можно остановить закрытием канала
	// close(stopCh)

	time.Sleep(3 * time.Second)

	fmt.Println("Exiting main.")
}
