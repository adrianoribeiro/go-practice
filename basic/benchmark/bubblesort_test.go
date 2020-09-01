package benchmark

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubblesort(t *testing.T) {
	values := getElements(5)
	BubbleSort(values)
	fmt.Println(values)

	assert.EqualValues(t, 1, values[0])
	assert.EqualValues(t, 2, values[1])
	assert.EqualValues(t, 3, values[2])
	assert.EqualValues(t, 4, values[3])
	assert.EqualValues(t, 5, values[4])
}

func getElements(n int) []int {
	values := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		values[i] = i + 1
	}

	return values
}

func BenchmarkBubblesort10(b *testing.B) {
	values := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(values)
	}
}

func BenchmarkBubblesort100(b *testing.B) {
	values := getElements(100)
	for i := 0; i < b.N; i++ {
		BubbleSort(values)
	}
}

func BenchmarkBubblesort1000(b *testing.B) {
	values := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(values)
	}
}
