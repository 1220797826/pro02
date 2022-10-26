package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 4, 1, 7, 6}
	sort.Ints(s)
	fmt.Printf("s: %v\n", s)
}
