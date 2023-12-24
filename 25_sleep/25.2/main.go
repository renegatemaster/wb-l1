// Разработать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

func sleep(i time.Duration) {
	<-time.After(i * time.Second)
}

func main() {
	sleep(5)
	fmt.Println("End of work")
}
