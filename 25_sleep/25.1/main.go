// Разработать собственную функцию sleep

package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(i time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), i*time.Second)
	defer cancel()

	<-ctx.Done()
}

func main() {
	sleep(5)
	fmt.Println("End of work")
}
