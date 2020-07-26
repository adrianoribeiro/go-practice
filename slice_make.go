package main

import "fmt"

func main() {

	values := make([]int, 3)
	fmt.Printf("Values len/cap: %v/%v\n", len(values), cap(values))
	fmt.Printf("Values: %v\n", values)

	fmt.Println("Change the first value to 10")
	values[0] = 10
	fmt.Printf("Values: %v\n", values)

	fmt.Println("Change the second value to 5")
	values[1] = 5
	fmt.Printf("Values: %v\n", values)
}
