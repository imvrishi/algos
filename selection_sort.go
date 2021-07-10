package selection_sort

import "fmt"

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	SelectionSort(array)
	fmt.Println(array)
}

func SelectionSort(array []int) {
	size := len(array)
	for i := 0; i < size; i++ {
		min := array[i]
		pos := i
		for j := i; j < size; j++ {
			if array[j] < min {
				min = array[j]
				pos = i
			}
		}

		array[i], array[pos] = array[pos], array[i]
	}
}
