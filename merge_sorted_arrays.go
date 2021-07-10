package main

import "fmt"

func MergeSortedArrays(array1, array2 []int) []int {
	size1 := len(array1)
	size2 := len(array2)

	if size1 == 0 {
		return array2
	}

	if size2 == 0 {
		return array1
	}

	index1 := 0
	index2 := 0
	mergedArrIndex := 0
	mergedArr := make([]int, size1+size2)
	for index1 < size1 && index2 < size2 {
		if array1[index1] < array2[index2] {
			mergedArr[mergedArrIndex] = array1[index1]
			index1++
		} else {
			mergedArr[mergedArrIndex] = array2[index2]
			index2++
		}
		mergedArrIndex++
	}

	if index1 < size1 {
		mergedArr = append(mergedArr[:mergedArrIndex], array1[index1:]...)
	}

	if index2 < size2 {
		mergedArr = append(mergedArr[:mergedArrIndex], array2[index2:]...)
	}

	return mergedArr
}

func main() {
	fmt.Println(MergeSortedArrays([]int{0, 3, 4, 31}, []int{4, 6, 30}))
}
