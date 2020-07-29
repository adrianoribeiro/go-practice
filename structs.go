package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Age                 int
}

func main() {

	person1 := Person{"Adriano", "Ribeiro", 38}
	fmt.Printf("Person1: %v\n", person1)

	person2 := Person{
		FirstName: "Juca",
		LastName:  "Jin",
	}
	fmt.Printf("Person2: %v\n", person2)
}
