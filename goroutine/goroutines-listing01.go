package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	//Try to change to 2
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}

		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}

		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()

	fmt.Println("\n\nTerminating program")
}
