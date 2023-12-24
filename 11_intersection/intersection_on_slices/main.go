package main

import "fmt"

func main() {

	first := []int{1, 0, 17, 28, 84, 42, 78}
	second := []int{4, 78, 10, 42, 66, 17}

	third := intersect(first, second)

	fmt.Println(third)
}

func intersect(first []int, second []int) []int {
	third := []int{}

	for i := range first {
		for j := range second {
			switch {
			case first[i] == second[j]:
				third = append(third, first[i])
			default:
				continue
			}
		}
	}

	return third
}
