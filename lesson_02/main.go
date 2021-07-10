package main

import (
	"fmt"
	"sequence/fibonacci"
)

func main() {
	defer fmt.Println("sequence")
	fmt.Println("Fibonacci sequence:")
	fibonacci.Print(5)
}
