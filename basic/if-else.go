package main

import "fmt"

func main() {

	fmt.Println("Adriano is 21, are Adriano able to driver?")
	age := 21
	if age >= 18 {
		fmt.Println("Yes, Adriano are able to driver")
	} else {
		fmt.Println("No, Adriano aren't able to driver")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 2")
	}
}
