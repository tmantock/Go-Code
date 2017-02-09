package main

import (
	"fmt"
)

func main() {
	var num int
	num = 4
	for i := 1; i <= num; i++ {
		fib := fibonacci(i)
		fmt.Printf("Fibonacci of %d is = %d\n", i, fib)
	}
}

func fibonacci(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	fib := fibonacci(x-1) + fibonacci(x-2)
	fmt.Println(fib)
	return fib
}
