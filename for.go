package main

import "fmt"

func main() {

	fmt.Println("Print from 1 to 10 (using for i <= 10)")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Print from 1 to 10 (for j:=1; j <= 10; j++)")
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	fmt.Println("Print only even values from 1 to 10")
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Println(k)
		}
	}
}
