package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(numbers...))

}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}
