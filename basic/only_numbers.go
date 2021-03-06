package main

import (
	"fmt"
)

type UserInput interface {
	Add(rune)
	GetValue() string
}

type NumericInput struct {
	input string
}

var onlyNumbers = make([]rune, 0)
func (numericInput *NumericInput) Add(character rune) {
	if character >= 48 && character <= 57 {
		onlyNumbers = append(onlyNumbers, character)
	}
}

func (numericInput *NumericInput) GetValue() string {
	return string(onlyNumbers)
}

func main() {
	var input UserInput = &NumericInput{}
	input.Add('1')
	input.Add('a')
	input.Add('0')
	fmt.Println(input.GetValue())
}
