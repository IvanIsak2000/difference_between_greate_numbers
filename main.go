package main

import (
	"fmt"
)

func add(first, second int64) int64 {
	var diff int64
	if first > second {
		diff = first - second
	}

	if first < second {

		diff = second - first

	}
	if first == second {
		diff = 0
	}
	return diff

}

func main() {
	var first, second int64
	fmt.Print("Write first number:")
	fmt.Scan(&first)
	fmt.Print("Write second number:")
	fmt.Scan(&second)
	var output = add(first, second)

	fmt.Print("Difference between: ", first, " and ", second, " is ", output)

}
