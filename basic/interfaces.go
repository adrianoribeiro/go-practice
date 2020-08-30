package main

import "fmt"

func init() {
	fmt.Println("It is call before.")
}

func main() {

	cat := Cat{}
	bird := Bird{}

	fmt.Println(print(cat))
	fmt.Println(print(bird))
}

type Animal interface {
	name() string
	eat() string
	fly() bool
}

type Cat struct{}

func (cat Cat) name() string {
	return "Cat"
}

func (cat Cat) eat() string {
	return "Fish"
}

func (cat Cat) fly() bool {
	return false
}

type Bird struct{}

func (bird Bird) name() string {
	return "Bird"
}

func (bird Bird) eat() string {
	return "earthworm"
}

func (bird Bird) fly() bool {
	return true
}

func print(animal Animal) string {
	if animal.fly() {

		return fmt.Sprintf("The %v eats %v and can fly", animal.name(), animal.eat())
	}

	return fmt.Sprintf("The %v eats %v and can't fly", animal.name(), animal.eat())
}
