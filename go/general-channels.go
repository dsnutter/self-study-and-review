package main

import (
	"fmt"
	"time"
)

func main() {
	// channels blocking
	// := declares var at same time
	c1 := make(chan string)
	c2 := make(chan string)

	// go statement is an independent concurrent thread of control
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	// go statement is an independent concurrent thread of control
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("recieved", msg2)
		}
	}
}
