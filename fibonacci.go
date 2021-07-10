package fibonacci

import "fmt"

func fibonacciIterative(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = a+b, a
	}
	return a
}

func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func main() {
	fmt.Println(fibonacciIterative(8))
	fmt.Println(fibonacciRecursive(8))
}
