package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	Name   string
	Salary float32
}

func main() {

	employees := []Employee{
		Employee{"Adriano", 1000},
		Employee{"Juca", 2000},
		Employee{"Simba", 800},
	}

	for _, name := range []string{"Hermenegildo", "Adriano", "Arquimedes", "Simba"} {
		err, employee := findByName(&employees, name)
		if err == nil {
			fmt.Printf("%s: %v\n", name, employee)
		} else {
			fmt.Printf("%s: %s\n", name, err)
		}
	}

}

func findByName(employees *[]Employee, name string) (error, Employee) {
	for _, employee := range *employees {
		if strings.EqualFold(employee.Name, name) {
			return nil, employee
		}
	}

	return errors.New("Employee not found"), Employee{}
}
