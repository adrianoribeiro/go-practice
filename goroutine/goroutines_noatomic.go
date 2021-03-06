package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counterNA int
	wgNA      sync.WaitGroup
)

func main() {

	wgNA.Add(2)

	go incCounterNA(1)
	go incCounterNA(2)

	wgNA.Wait()
	fmt.Println("Final counter: ", counter)
}

func incCounterNA(id int) {

	defer wgNA.Done()

	for count := 0; count < 2; count++ {

		value := counterNA

		runtime.Gosched()

		value++

		counterNA = value
		// fmt.Printf("Temporary counter (%v): %v\n", id, counter)
	}
}
