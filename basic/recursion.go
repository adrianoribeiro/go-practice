package main

import "fmt"

func main() {

	for i := 0; i < 6; i++ {
		fmt.Printf("The factorial of %v is %v\n", i, factorial(i))
	}

}

func factorial(number int) int {

	if number <= 2 {
		return number
	}

	return number * factorial(number-1)
}
