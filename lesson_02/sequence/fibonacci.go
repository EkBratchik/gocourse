package fibonacci

import "fmt"

func Fib(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}
func Print(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", Fib(i))
	}
}
