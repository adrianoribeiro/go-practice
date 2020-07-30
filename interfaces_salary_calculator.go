package main

import "fmt"

func main() {

	employees := []SalaryCalculator{
		Programmer{"Adriano", 1000},
		Manager{"Juca", 2000},
		Programmer{"Simba", 800},
	}

	expense := float32(0)
	for _, employee := range employees {
		expense += employee.SalaryCalculator()
	}

	fmt.Printf("Total Expense Per Month is %v\n", expense)

}

type SalaryCalculator interface {
	SalaryCalculator() float32
}

type Programmer struct {
	name   string
	salary float32
}

type Manager struct {
	name   string
	salary float32
}

func (programmer Programmer) SalaryCalculator() float32 {
	return programmer.salary
}

//The manager earn a % over its salary
func (manager Manager) SalaryCalculator() float32 {
	return manager.salary * 1.2
}
