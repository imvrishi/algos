package contains_common

import "fmt"

func ContainsCommon(arr1, arr2 []rune) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}
	return false
}

func ContainsCommon2(arr1, arr2 []rune) bool {
	mapped := map[rune]rune{}
	for i := 0; i < len(arr1); i++ {
		mapped[arr1[i]] = arr1[i]
	}
	for j := 0; j < len(arr2); j++ {
		if _, ok := mapped[arr2[j]]; ok {
			return true
		}
	}
	return false
}

func main() {
	array1 := []rune{'a', 'b', 'c', 'x'}
	array2 := []rune{'z', 'y', 'd'}
	fmt.Println(ContainsCommon(array1, array2))
	fmt.Println(ContainsCommon2(array1, array2))
}
