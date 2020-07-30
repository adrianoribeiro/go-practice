package main

import "fmt"

func main() {

	describe("Hello")
	describe(1)
	describe(1.0)
	describe(false)
}

func describe(value interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}
