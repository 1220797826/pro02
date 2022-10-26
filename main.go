package main

import (
	"fmt"
	"time"
)

var values = make(chan int)

func send() {
	vaule := 1
	fmt.Println(vaule)
	time.Sleep(time.Second * 5)
	values <- vaule
}
func main() {
	// services.Test()
	defer close(values)
	go send()
	fmt.Println("wait .....")
	a := <-values
	fmt.Println(a)
}
