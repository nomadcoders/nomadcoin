package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	for i := range [10]int{} {
		fmt.Printf(">> sending %d <<\n", i)
		c <- i
		fmt.Printf(">> sent %d <<\n", i)
	}
	close(c)
}

func receive(c <-chan int) {
	for {
		time.Sleep(10 * time.Second)
		a, ok := <-c
		if !ok {
			fmt.Println("we are done.")
			break
		}
		fmt.Printf("|| received %d ||\n", a)
	}
}

func main() {
	c := make(chan int, 10)
	go send(c)
	receive(c)
}
