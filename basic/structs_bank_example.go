package main

import (
	"fmt"
	"go-practice/basic/account"
)

func main() {

	personAdriano := account.Person{"Adriano", "Ribeiro"}
	personJuca := account.Person{"Juca", "Jin"}

	account1 := account.Account{
		Person: personAdriano,
		Amount: 1000,
	}

	account2 := account.Account{
		Person: personJuca,
		Amount: 1000,
	}

	fmt.Println("Initial state *************************************")
	account1.Print()
	account2.Print()

	fmt.Println("\nTransfer 100 from Adriano to Juca *****************")
	account1.Transfer(&account2, 100)
	account1.Print()
	account2.Print()

	fmt.Println("\nTransfer 500 from Juca to Adriano *****************")
	account2.Transfer(&account1, 500)
	account1.Print()
	account2.Print()
}
