package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counterMutex int
	wgMutex      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final counter: ", counter)
}

func incCounterMutex(id int) {

	defer wg.Done()

	for cont := 0; cont < 2; cont++ {
		mutex.Lock()
		{
			value := counter

			runtime.Gosched()

			value++

			counter = value
		}
		mutex.Unlock()
	}

}
