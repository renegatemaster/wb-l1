// Удалить i-ый элемент из слайса.

package main

import (
	"errors"
	"fmt"
)

func main() {

	slice := []int{2, 4, 6, 8, 10}

	newSlice, err := rmElem(slice, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newSlice)
}

func rmElem(slice []int, i int) ([]int, error) {

	// Не будем обрабатывать пустые слайсы
	if len(slice) == 0 {
		err := errors.New("given slice is empty")
		return nil, err
	}
	// На случай, если индекс вне слайса
	if i >= len(slice) || i < 0 {
		err := errors.New("given index is out of range")
		return nil, err
	}

	newSlice := slice[:i+copy(slice[i:], slice[i+1:])]
	return newSlice, nil
}
