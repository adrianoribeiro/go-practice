package main

import "fmt"

func main() {

	fruits := []string{"Apple", "Orange", "Banana"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	fruitsWithIndex := map[int]string{0: "Apple", 1: "Orange", 2: "Banana"}
	for idx, fruit := range fruitsWithIndex {
		fmt.Printf("%v: %v\n", idx, fruit)
	}
}
