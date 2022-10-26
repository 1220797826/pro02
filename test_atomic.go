package main

import (
	"fmt"
	"sync/atomic"
)

var i int32 = 100

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func load() {
	var i1 int32 = 100
	fmt.Printf("atomic.LoadInt32(&i1): %v\n", atomic.LoadInt32(&i1))

	atomic.StoreInt32(&i1, 200)
}
func main() {

	load()
}
