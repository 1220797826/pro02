package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100
var lock sync.Mutex

func add() {
	// lock.Lock()
	i++
	fmt.Printf("i++: %v\n", i)
	// lock.Unlock()
}

func sub() {
	// lock.Lock()
	i--
	fmt.Printf("i--: %v\n", i)
	// lock.Unlock()
}
func main() {

	for j := 0; j < 100; j++ {
		go add()
		go sub()
	}
	time.Sleep(time.Second * 3)
	fmt.Printf("end......: %v\n", i)

}
