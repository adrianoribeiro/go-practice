package main

import "fmt"

func main() {

	user := make(map[int]string)
	user[1] = "Adriano"
	user[2] = "Barnabe"
	user[3] = "Ciro"

	fmt.Println(user)
	fmt.Printf("Len: %v\n", len(user))

	fmt.Println("//Removing id 2")
	delete(user, 2)
	fmt.Println(user)
	fmt.Printf("Len: %v\n", len(user))

	weekdays := map[int]string{1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 7: "Sun"}
	fmt.Println(weekdays)
}
