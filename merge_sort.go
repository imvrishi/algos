package merge_sort

import (
	"fmt"
	"math"
)

func MergeSort(array []int) []int {
	size := len(array)
	if size < 2 {
		return array
	}

	mid := int32(math.Round(float64(size) / 2))
	left := array[:mid]
	right := array[mid:]

	return Merge(
		MergeSort(left),
		MergeSort(right),
	)
}

func Merge(left, right []int) []int {
	temp := []int{}
	leftSize := len(left)
	rightSize := len(right)
	leftIndex := 0
	rightIndex := 0

	for leftIndex < leftSize && rightIndex < rightSize {
		if left[leftIndex] < right[rightIndex] {
			temp = append(temp, left[leftIndex])
			leftIndex++
		} else {
			temp = append(temp, right[rightIndex])
			rightIndex++
		}
	}

	if leftIndex < leftSize {
		temp = append(temp, left[leftIndex:]...)
	}

	if rightIndex < rightSize {
		temp = append(temp, right[rightIndex:]...)
	}

	return temp
}

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(MergeSort(array))
}
