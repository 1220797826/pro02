package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	chanInt := make(chan int)
	go func() {
		for _ = range ticker.C {
			select {
			case chanInt <- 1:
				fmt.Println("1111")
			case chanInt <- 2:
				fmt.Println("2222")
			case chanInt <- 3:
				fmt.Println("3333")
			}
		}

	}()
	sum := 0
	for v := range chanInt {
		fmt.Println("收到:", v)
		sum += v
		if sum > 10 {
			break
		}
	}
}
