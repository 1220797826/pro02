package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 10)

	fmt.Println("begin")

	go func() {
		<-timer.C
		fmt.Println("end.....")
	}()
	// timer.Stop()
	timer.Reset(time.Second)
	time.Sleep(time.Second * 20)
}
