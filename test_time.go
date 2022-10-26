package main

import (
	"fmt"
	"time"
)

func test2() {
	now := time.Now()

	fmt.Printf("now.Unix(): %v\n", now.Unix())
	fmt.Printf("now.UnixNano(): %v\n", now.UnixNano())
	fmt.Printf("time.Unix(now.Unix()): %v\n", time.Unix(now.Unix(), 0))
}

func test3() {
	t := time.Now()
	fmt.Printf("t.Format(\"2006-01-02\"): %v\n", t.Format("2006-01-02"))
}
func main() {
	test3()
}
