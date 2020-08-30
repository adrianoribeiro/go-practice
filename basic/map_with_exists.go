package main

import "fmt"

func main() {
	numberOccurs := map[string]int{}
	numbers := []string{"four", "one", "three", "one", "three", "one", "one", "three", "one", "four", "two"}
	for _, number := range numbers {
		_, exist := numberOccurs[number]

		if exist {
			numberOccurs[number]++
		} else {
			numberOccurs[number] = 1
		}
	}

	printMap(&numberOccurs)
	fmt.Println("------------------------")
	delete(numberOccurs, "four")
	printMap(&numberOccurs)
}

func printMap(numberOccurs *map[string]int) {
	for key, value := range *numberOccurs {
		fmt.Printf("Key: %s | Value: %v\n", key, value)
	}
}
