package main

import (
	"fmt"
	"time"
)

func main() {

	var c = make(chan string)

	start := time.Now()

	//No blocking
	go message1(c)
	go message2(c)
	go message3(c)

	//blocking
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	close(c)

	elapse := time.Since(start)
	fmt.Printf("\nTime elapsed: %s\n", elapse)
}

func message1(c chan string) {

	time.Sleep(3 * time.Second)
	c <- "msg2, String delay 3sec"
}

func message2(c chan string) {

	time.Sleep(4 * time.Second)
	c <- "msg3, String delay 4sec"
}

func message3(c chan string) {

	time.Sleep(2 * time.Second)
	c <- "msg1, String delay 2sec"
}
