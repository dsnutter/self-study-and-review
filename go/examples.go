package main

import (
	"fmt"
)

func main() {
	{
		var msg string = "World - stuff... SUM:"
		var x int = 12
		var y int = 45
		var sum int

		if sum = x + y; sum > 46 {
			fmt.Println(msg, sum)
		}
	}

	// typing testing
	var x interface{} = 53.2
	//	var x interface{} = 53
	//	var x interface{} = "Hello"

	switch v := x.(type) {
	case string:
		fmt.Println("Type is string...")
	case int:
		fmt.Println("Type is int...")
	default:
		fmt.Printf("Type is %T\n", v)
	}

	// looping, no such thing as a while loop
	numbers := []int{2, 4, 6, 8}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	t := 0
	// closest to while loop
	for t < 5 {
		fmt.Println(t)
		t++
	}

}
