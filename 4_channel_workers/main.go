package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	// Создаём канал для завершение работы по нажатию Ctrl+C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Выбор количества воркеров и создание канала
	num := getNumberOfWorkers()
	ch := make(chan int, num)

	// Создаём заданное количество воркеров-горутин
	// и синхронизируем их
	wg := sync.WaitGroup{}
	for i := 1; i < num+1; i++ {
		wg.Add(1)
		go startWorker(ch, &wg, i)
	}

	// Постоянная запись в канал до получения сигнала
	for i := 0; ; i++ {
		select {
		case <-signalChan:
			fmt.Printf("\nReceived an interrupt, closing channel and stopping workers...\n\n")
			close(ch) // при получении сигнала канал закрывается
			wg.Wait() // основная горутина ждёт воркеры-горутины
			return
		default: // выбераем имеенно этот оператор, чтобы публикация была только в случае, если не нажали Ctrl+C
			ch <- i
			fmt.Fprintf(os.Stdout, "Published %d \n", i)
			time.Sleep(time.Second) // немного тормозим цикл, чтобы сделать его более читаемым
		}
	}
}

func getNumberOfWorkers() int {
	fmt.Println("Choose number of workers:")
	var num int
	fmt.Scanln(&num)
	if num <= 0 {
		fmt.Println("Number of workers must be greater than 0")
		return getNumberOfWorkers()
	}
	return num
}

func startWorker(ch <-chan int, wg *sync.WaitGroup, i int) {
	fmt.Fprintf(os.Stdout, "Worker %d started\n", i)
	defer wg.Done()
	defer fmt.Fprintf(os.Stdout, "Worker %d stopped\n", i)
	for {
		num, ok := <-ch
		if !ok {
			break
		}
		fmt.Fprintf(os.Stdout, "Worker %d prints %d\n", i, num)
	}
}
