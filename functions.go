package main

import "fmt"

func main() {
	fmt.Println(plus(5, 7))
	fmt.Println(concat("Adriano", "Ribeiro"))
}

func plus(value1, value2 int) int {
	return value1 + value2
}

func concat(name, lastName string) string {
	return name + " " + lastName
}
