package main

import (
	"fmt"
	"runtime"
)

const (
	outputMessage = "%s has worked %d weeks and %d days in the company"
	outputMessageWithoutDays = "%s has worked %d weeks in the company"
)

func main() {

	people := make(map[string]int, 0)
	people["Robert"] = 30
	people["John"] = 475
	people["Elly"] = 1022
	people["Bob"] = 99

	processInParallel(people, 2)
}

func processInParallel(people map[string]int, numberParallelProcess int) {
	runtime.GOMAXPROCS(numberParallelProcess)
	wg.Add(len(people))

	for name, workDays := range people {

		go print(name, workDays)
	}
	wg.Wait()
}

func print(name string, workDays int) {
	defer wg.Done()

	weeks, restDays := numOfWeeks(workDays)
	if restDays == 0 {
		fmt.Println(fmt.Sprintf(outputMessageWithoutDays, name, weeks))
	} else {
		fmt.Println(fmt.Sprintf(outputMessage, name, weeks, restDays))
	}
}

func numOfWeeks(workDays int) (int, int) {
	weeks := workDays/7
	restDays := workDays - (weeks*7)
	return weeks, restDays
}
