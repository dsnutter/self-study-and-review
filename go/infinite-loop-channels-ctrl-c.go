package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	//
	// infinite looping, break and continue are keywords also,
	//  this uses an interrupt signal handler to break on ctrl-c
	//
	c := make(chan os.Signal, 1)
	var closeOk bool = false
	var result os.Signal

	// channel is for a ctrl-c interrupt that is a signal handler
	//   operates async with go channels
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	// handler function for c channel for ctrl-c
	// go statement is an independent concurrent thread of control
	go func() {
		result, closeOk = <-c
		// check from function, receives a value from the channel
		if closeOk && result == os.Interrupt {
			fmt.Println("Cleanup")
			close(c)
			return
		}
	}()

	for {
		fmt.Println("infinite loop")
		//fmt.Printf("%d - %d", len(c), c)
		// we can close the channel and capacity [not the
		//   length off the buffer] is 1. Different than
		//   a length function in other languages, len()
		//   is more like a size() is other languageds for go
		//   channels
		if closeOk && cap(c) >= 1 {
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
