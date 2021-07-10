package insertion_sort

import "fmt"

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	InsertionSort(array)
	fmt.Println(array)
}

func InsertionSort(array []int) {
	size := len(array)
	insertion := []int{array[0]}
	insertion_count := 1
	for i := 1; i < size; i++ {
		if array[i] > insertion[0] && array[i] < insertion[insertion_count-1] {
			for j := 1; j < i; j++ {
				if array[i] > insertion[j-1] && array[i] < insertion[j] {
					insertion = append(insertion[:j], insertion[j-1:]...)
					insertion[j] = array[i]
					insertion_count++
					break
				}
			}
		} else if array[i] < insertion[0] {
			insertion = append([]int{array[i]}, insertion...)
			insertion_count++
		} else {
			insertion = append(insertion, array[i])
			insertion_count++
		}
	}

	copy(array, insertion)
}
