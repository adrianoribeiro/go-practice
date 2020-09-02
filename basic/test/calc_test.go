package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	var myTest = []struct {
		number1 int
		number2 int
		result  int
	}{
		{1, 1, 2},
		{5, 5, 10},
		{7, 5, 12},
		{400, 55, 455},
	}

	for _, testElement := range myTest {
		assert.EqualValues(t, testElement.result, Sum(testElement.number1, testElement.number2))
	}
}
