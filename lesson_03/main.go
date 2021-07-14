package main

import (
	"fmt"
	"sort"
)

func AverageArray(a [6]float32) float32 {
	var c float32
	var k float32
	k = 0
	for _, v := range a {
		c += v
		k++
	}
	return c / k
}
func Max(a []string) string {
	var longestWord string
	longestLength := 0
	for _, word := range a {
		length := len(word)

		if length > longestLength {
			longestWord = word
			longestLength = length
		}
	}
	return longestWord
}
func Reverse(a []int64) []int64 {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
func PrintSorted(a map[int]string) {
	b := map[int]string{}
	key := []int{}

	for k, val := range a {
		b[k] = val
		key = append(key, k)
	}
	sort.Ints(key)
	for _, k := range key {
		fmt.Print(a[k], " ")
	}
}
func main() {
	a := [6]float32{1, 2, 3, 4, 5, 6}
	fmt.Println(AverageArray(a))

	d := []string{"one", "two"}
	fmt.Println(Max(d))

	b := []int64{1, 2, 5, 15}
	fmt.Println(Reverse(b))

	c := map[int]string{4: "aa", 3: "dd", 1: "dad"}
	PrintSorted(c)
}
