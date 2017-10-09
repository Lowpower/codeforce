package main

import (
	"fmt"
)

var num int
var temp int

func main() {
	fmt.Scan(&num)
	max := -1
	for i := 0; i < num; i++ {
		fmt.Scan(&temp)
		if temp > max {
			max = temp
		}
	}
	if max > 25 {
		fmt.Println(max - 25)
	} else {
		fmt.Println(0)
	}
}
