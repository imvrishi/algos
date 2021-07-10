package bubble_sort

import "fmt"

func main() {
	array := []int{6, 5, 3, 1, 8, 7, 2, 4}
	BubbleSort(array)
	fmt.Println(array)
}

func BubbleSort(array []int) {
	size := len(array)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
