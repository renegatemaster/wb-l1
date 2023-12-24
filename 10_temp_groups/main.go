// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}
package main

import (
	"fmt"
	"sync"
)

func main() {

	var temps []float64 = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemps(temps)

	fmt.Println(groups)
}

func groupTemps(temps []float64) map[int][]float64 {

	groups := make(map[int][]float64)
	wg := sync.WaitGroup{}

	for _, temp := range temps {
		wg.Add(1)
		go func(temp float64) {
			key := int(temp/10) * 10
			groups[key] = append(groups[key], temp)
			wg.Done()
		}(temp)
	}

	wg.Wait()
	return groups
}
