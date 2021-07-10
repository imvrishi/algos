package check_sort

import (
	"fmt"
	"sort"
)

func main() {
	letters := []string{"a", "d", "z", "e", "r", "b"}
	sort.Strings(letters)

	basket := []int{2, 65, 32, 2, 1, 7, 8}
	sort.Ints(basket)

	spanish := []string{"réservé", "Premier", "Cliché", "communiqué", "café", "Adieu"}
	sort.Strings(spanish)

	fmt.Println(letters)
	fmt.Println(basket)
	fmt.Println(spanish)
}
