package main

import (
	"fmt"
	"sort"
)

type testSlice []int

func (x testSlice) Len() int           { return len(x) }
func (x testSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x testSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func main() {
	a := []int{4, 6, 3, 8, 1}

	sort.Sort(testSlice(a))

	fmt.Printf("a: %v\n", a)

	i2 := sort.Reverse(testSlice(a))
	fmt.Printf("a: %v\n", i2)
}
