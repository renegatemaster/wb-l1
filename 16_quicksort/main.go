// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	array := []int{18, 2, 11, 7, 4, 5, 1}
	fmt.Println(array)

	quickSort(array)

	fmt.Println(array)
}

func quickSort(array []int) []int {

	if len(array) < 2 {
		return array
	}

	left, right := 0, len(array)-1
	pivot := rand.Int() % len(array)

	array[pivot], array[right] = array[right], array[pivot]

	for i := range array {
		if array[i] < array[right] {
			array[left], array[i] = array[i], array[left]
			left++
		}
	}
	array[left], array[right] = array[right], array[left]

	quickSort(array[:left])
	quickSort(array[left+1:])

	return array
}
