package main

import (
	"fmt"
	"runtime"
)

func show(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str)
	}
}

func main() {

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	return
	go show("go .....")

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("main ....")
	}
}
