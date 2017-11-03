package main

import (
	"fmt"
)

var fuckMap = map[string]int{
	"X++": 1,
	"++X": 1,
	"X--": -1,
	"--X": -1,
}

func main() {
	var (
		inputnum int
		input    string
	)
	fmt.Scan(&inputnum)
	output := 0
	for i := 0; i < inputnum; i++ {
		fmt.Scan(&input)
		output += fuckMap[input]
	}
	fmt.Println(output)
}
