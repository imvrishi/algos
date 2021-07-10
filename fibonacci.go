package main

import (
	"fmt"
	"time"
)

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

func fibonacciRecursiveMemoized(n int, cache *map[int]int) int {
	if n < 2 {
		return n
	}
	if v, ok := (*cache)[n]; ok {
		return v
	} else {
		(*cache)[n] = fibonacciRecursiveMemoized(n-1, cache) + fibonacciRecursiveMemoized(n-2, cache)
		return (*cache)[n]
	}
}

func main() {
	// start := time.Now()
	// fmt.Println(fibonacciRecursive(40))
	// fmt.Println(time.Since(start))
	cache := &map[int]int{}
	start := time.Now()
	fmt.Println(fibonacciRecursiveMemoized(40, cache))
	fmt.Println(time.Since(start))
}
