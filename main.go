package main

import (
	"fmt"
	"time"
)

func countToTen(name string) {
	for i := range [10]int{} {
		fmt.Println(i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go countToTen("first")
	go countToTen("second")
	for {
	}
}
