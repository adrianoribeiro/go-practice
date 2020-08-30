package main

import "fmt"

/*
Type assertion is used to extract the underlying value of the interface.

i.(T) is the syntax which is used to get the underlying value of interface i whose concrete type is T.
*/

func main() {

	getInt(38)

	var newNumber interface{} = 39
	getInt(newNumber)

	//Trying to pass a string intead of int
	getInt("Adriano")
}

func getInt(value interface{}) {
	extractValue, ok := value.(int) //get the underlying int value from value
	if ok {
		fmt.Println(extractValue)
	} else {
		fmt.Println("Invalid number.")
	}

}
