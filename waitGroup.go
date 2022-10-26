package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func send(i int) {
	defer wp.Done()
	fmt.Printf("i: %v\n", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wp.Add(1)
		go send(i)

	}

	wp.Wait()
	fmt.Println("end.....")
}
