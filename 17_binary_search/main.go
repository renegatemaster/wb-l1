// Реализовать бинарный поиск встроенными методами языка
package main

import (
	"errors"
	"fmt"
)

func main() {

	array := []int{1, 3, 5, 7, 9}

	index, err := binarySearch(array, 7)

	if index < 0 {
		fmt.Println(err)
	} else {
		fmt.Printf("The index of item is %d\n", index)
	}
}

func binarySearch(array []int, item int) (int, error) {

	// Индексы границ поиска
	low, high := 0, len(array)-1

	for low <= high {
		mid := (low + high) / 2
		guess := array[mid]
		if guess == item {
			return mid, nil
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	err := errors.New("no such item in provided array")
	return -1, err
}
