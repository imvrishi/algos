package counting_sort

import (
	"fmt"
)

func Max(array []int) int {
	size := len(array)
	max := array[0]
	for i := 0; i < size; i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func CountSort(array []int) []int {
	max := Max(array)
	bucket := make([][]int, max+1)
	count := len(array)

	for i := 0; i < count; i++ {
		a := array[i]
		bucket[a] = append(bucket[a], a)
	}

	output := make([]int, count)
	k := 0
	for i := 0; i < max+1; i++ {
		size := len(bucket[i])
		for j := 0; j < size; j++ {
			output[k] = bucket[i][j]
			k++
		}
	}

	return output
}

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println("Before Sort:", array)
	fmt.Println("After Sort :", CountSort(array))
}
