package account

import "fmt"

type Person struct {
	FirstName, LastName string
}

type Account struct {
	Person Person
	Amount float32
}

func (account Account) Print() {
	fmt.Printf("{\n\t\"name\": \"%v %v\",\n\t\"amount \": %v\n}\n", account.Person.FirstName, account.Person.LastName, account.Amount)
}

func (account *Account) Transfer(accountSource *Account, amount float32) {
	account.Amount -= amount
	accountSource.Amount += amount
}
