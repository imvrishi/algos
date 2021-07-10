package quick_sort

import "fmt"

func Partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}

func QuickSort(array []int, low, high int) []int {
	if low < high {
		mid := Partition(array, low, high)

		QuickSort(array, low, mid-1)
		QuickSort(array, mid+1, high)
	}

	return array
}

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Println(QuickSort(array, 0, len(array)-1))
}
