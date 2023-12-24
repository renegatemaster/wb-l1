// Реализовать все возможные способы
// остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"time"
)

func myGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting goroutine.")
			return
		default:
			fmt.Println("Doing some work...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go myGoroutine(ctx)

	// Имитируем работу main, чтобы дать время myGoroutine
	time.Sleep(3 * time.Second)

	// Останавливаем горутину завершением контекста
	cancel()

	time.Sleep(3 * time.Second)

	fmt.Println("Exiting main.")
}
