package main

import "fmt"

func main() {
	number := 5
	previous, next := previousAndNext(5)
	fmt.Printf("The previous/next value of %v is: %v and %v\n", number, previous, next)
}

func previousAndNext(value int) (int, int) {
	return value - 1, value + 1
}
