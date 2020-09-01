package benchmark

func BubbleSort(elements []int) {

	for i := 0; i < len(elements); i++ {
		for j := i; j < len(elements); j++ {
			if elements[i] > elements[j] {
				aux := elements[i]
				elements[i] = elements[j]
				elements[j] = aux
			}
		}
	}
}
