package main

import (
	"fmt"
)

func main() {

	fmt.Println("Sample 1")
	var team [3]string
	team[0] = "Flamengo"
	team[1] = "Fluminense"
	team[2] = "Botafogo"

	for _, element := range team {
		fmt.Println(element)
	}

	fmt.Println("\nSample 2")
	newTeam := []string{"Flamengo", "Botafogo", "Fluminense", "Vasco"}
	for index, element := range newTeam {
		fmt.Println(index, ": ", element)
	}

	fmt.Println("\nSample 3 (with indexes)")
	newTeam2 := [5]string{0: "Flamengo", 4: "Botafogo"}
	for index, element := range newTeam2 {
		fmt.Println(index, ": ", element)
	}
}
