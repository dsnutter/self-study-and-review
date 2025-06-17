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
		"z": syscall.SIGTSTP,
	}

	fmt.Println("Capture which? Ctrl-C [c], Ctrl-Z [z]?")
	fmt.Scanf("%s", &input)
	var desiredResult os.Signal = inputInterrupt["c"]

	switch {
	case input[0:1] == "C":
	case input[0:1] == "c":
		fmt.Println("Breaking out of loop on Ctrl-C")
	case input[0:1] == "Z":
	case input[0:1] == "z":
		fmt.Println("Breaking out of loop on Ctrl-Z")
		desiredResult = inputInterrupt["z"]
	default:
		os.Exit(1)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, inputInterrupt["c"], inputInterrupt["z"])
	var result os.Signal
	var closeOk bool

	go func() {
		result, closeOk = <-c
		if result == desiredResult && cap(c) >= 1 {
			closeInterrupt(c)
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		if closeOk && cap(c) >= 1 {
			break
		} else {
			fmt.Println("infinite loop")
		}
	}
}

func closeInterrupt(c chan os.Signal) {
	fmt.Println("Closing and stopping")
	close(c)
}
