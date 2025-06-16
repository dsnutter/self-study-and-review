package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var input string

	inputInterrupt := map[string]os.Signal{
		"c": syscall.SIGINT,
		// 0x04 => Ctrl-D
		"d": syscall.Signal(0x04),
		"z": syscall.SIGTSTP,
	}

	fmt.Println("Capture which? Ctrl-C [c], Ctrl-D [d] or Ctrl-Z [z]?")
	fmt.Scanf("%s", &input)

	switch {
	case input[0:1] == "C":
	case input[0:1] == "c":
		fmt.Println("Breaking out of loop on Ctrl-C")
	case input[0:1] == "D":
	case input[0:1] == "d":
		fmt.Println("Breaking out of loop on Ctrl-D")
	case input[0:1] == "Z":
	case input[0:1] == "z":
		fmt.Println("Breaking out of loop on Ctrl-Z")
	default:
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, inputInterrupt["c"], inputInterrupt["d"], inputInterrupt["z"])

	// separate closeOk necessary due to closing the interrupt in the callback
	closeOk := make(chan bool, 1)

	// concurrent thread with go keyword, with callback for closing
	go interrupt(closeOk, c, closeInterrupt)

	for {
		fmt.Println("infinite loop")

		// if cap(closeOk) >= 1 {
		// 	select {
		// 	case ok := <-closeOk:
		// 		if ok {
		// 			close(closeOk)
		// 			break
		// 		}
		// 	}
		// }
		if cap(closeOk) >= 1 {
			for ok := range closeOk {
				if ok {
					close(closeOk)
					break
				}
			}
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func closeInterrupt(c chan os.Signal) {
	fmt.Println("Closing and stopping")
	close(c)
}

func interrupt(closeOk chan bool, c chan os.Signal, callback func(c chan os.Signal)) {
	var result os.Signal
	var cOk bool
	result, cOk = <-c

	if cOk && result == os.Interrupt && cap(c) >= 1 {
		callback(c)
		go func() {
			closeOk <- true
		}()
	}
}
