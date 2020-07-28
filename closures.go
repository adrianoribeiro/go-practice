package main

import "fmt"

func main() {

	fmt.Println("Go functions may be closures. A closure is a function value that references ")
	fmt.Println("variables from OUTSIDE its body. The function may access and assign to the")
	fmt.Println("referenced variables; in this sense the function is \"bound\" to the variables.\n\n")

	nextValue := getNextValue(5)

	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())

	otherNextValue := getNextValue(10)

	fmt.Println(otherNextValue())
	fmt.Println(otherNextValue())
}

func getNextValue(number int) func() int {
	i := number

	return func() int {
		i++
		return i
	}
}
