package factorial

import "fmt"

func findFactorialIterative(n int) int {
	fact := 1
	for i := n; i > 1; i-- {
		fact *= i
	}
	return fact
}

func findFactorialRecursive(n int) int {
	if n <= 2 {
		return n
	}

	return n * findFactorialRecursive(n-1)
}

func main() {
	fmt.Println(findFactorialIterative(5))
	fmt.Println(findFactorialRecursive(5))
}
