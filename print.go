package main

import "fmt"

func main() {

	fmt.Printf("%v is divisible by %v", 10, 5)
	fmt.Print("\n")
	fmt.Printf("Is %v divisible by %v? %v", 10, 3, 10%3 == 0)
	fmt.Print("\n")
}
