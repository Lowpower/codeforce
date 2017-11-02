package main

import (
	"fmt"
)

func main() {
	var (
		inputnum int
		input1   int
		input2   int
		input3   int
	)
	fmt.Scan(&inputnum)
	output := 0
	for i := 0; i < inputnum; i++ {
		fmt.Scan(&input1, &input2, &input3)
		if input1+input2+input3 >= 2 {
			output += 1
		}
	}
	fmt.Println(output)
}
