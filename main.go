package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int) {
	for i := range [10]int{} {
		time.Sleep(2 * time.Second)
		fmt.Printf("sending %d\n", i)
		c <- i
	}
}

func main() {
	c := make(chan int)
	go countToTen(c)
	for {
		a := <-c
		fmt.Printf("received %d\n", a)
	}

}
