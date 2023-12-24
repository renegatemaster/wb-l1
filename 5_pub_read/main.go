// Разработать программу, которая будет
// последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Запрашиваем время работы программы
	timer := getRuntime()

	// Используем контекст WithTimeout,
	// чтобы ограничитить исполнение по времени
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), timer)
	defer fmt.Println("Context cancelled")
	defer cancel()

	// Читаем канал
	go read(ch)

	// Пишем в канал до истечения заданного времени
	write(ch, ctx)
	fmt.Println("Channel closed")
}

func read(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received value: ", value)
	}
}

func write(ch chan<- int, ctx context.Context) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done(): // по истечении заданного времени закрываем канал
			fmt.Println("Time is over, closing channel...")
			close(ch)
			return
		default:
			ch <- i
		}
		time.Sleep(500 * time.Millisecond) // тормозим цикл для наглядности
	}
}

func getRuntime() time.Duration {
	fmt.Println("Enter run time in seconds:")
	var timer time.Duration
	fmt.Scanln(&timer)
	if timer <= 0 {
		fmt.Println("Run time must be greater than 0")
		return getRuntime()
	}
	return timer * time.Second
}
