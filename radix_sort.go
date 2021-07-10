package radix_sort

import (
	"fmt"
	"math"
	"strconv"
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

func RadixSort(array []int) []int {
	max := Max(array)
	num_length := len(strconv.Itoa(max))
	count := len(array)
	bucket := [10][]int{}
	output := array

	for i := 0; i < num_length; i++ {
		for j := 0; j < count; j++ {
			pow := int(math.Pow10(i))
			rem := output[j] % (10 * pow)
			num := rem / (1 * pow)
			bucket[num] = append(bucket[num], output[j])
		}

		n := 0
		for j := 0; j < 10; j++ {
			size := len(bucket[j])
			for k := 0; k < size; k++ {
				output[n] = bucket[j][k]
				n++
			}
		}
		bucket = [10][]int{}
	}

	return array
}

func main() {
	array := []int{66, 99, 0, 13, 57, 101, 97, 84}
	fmt.Println("Before Sort:", array)
	fmt.Println("After Sort :", RadixSort(array))
}
