package main

import "fmt"

func main() {

	/*
		Go supports pointers, allowing you to pass references
		to values and records within your program
	*/

	value := 1
	fmt.Println(value)

	changeValue(value)
	fmt.Println(value) //Print 1 (value was not change)

	//The &value syntax gives the memory address of value.
	changeValuePointer(&value)
	fmt.Println(value) //Print 2
}

func changeValue(value int) {
	value = 2
}

//int pointer
func changeValuePointer(value *int) {
	*value = 2
}
