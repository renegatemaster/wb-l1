// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
package main

import "fmt"

func main() {

	stringSlice := []string{"cat", "cat", "dog", "cat", "tree"}

	stringArray := sliceToArray(stringSlice)

	// Выводим значения множества в одну строку
	for i := range stringArray {
		fmt.Printf("%s ", i)
	}
	fmt.Println()
}

func sliceToArray(slice []string) map[string]bool {

	array := make(map[string]bool)

	// Дублирование ключей невозможно,
	// поэтому остаются только уникальные элементы
	for _, str := range slice {
		array[str] = true
	}

	return array
}
