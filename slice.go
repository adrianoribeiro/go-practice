package main

import "fmt"

func main() {

	fmt.Println("#Print all days of the week")
	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	fmt.Printf("days=%v\n\n", days)

	fmt.Println("#Print only weekdays")
	weekdays := days[:5]
	fmt.Printf("days=%v\n\n", weekdays)

	fmt.Println("#Print only weekend")
	weekend := days[5:]
	fmt.Printf("days=%v\n\n", weekend)

	fmt.Println("#Print from Wed to Sat")
	wedToSat := days[2:6]
	fmt.Printf("days=%v\n\n", wedToSat)

	fmt.Printf("#Wed to Sat (len): %v ( Mon, Tue, [Wed, Thu, Fri, Sat], Sun)\n", len(wedToSat))
	fmt.Printf("#Wed to Sat (cap): %v ( Mon, Tue, [Wed, Thu, Fri, Sat, Sun])\n", cap(wedToSat))
}
