package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	max := 0
	curr := 0
	var indoor, outdoor int
	for i := 0; i < num; i++ {
		fmt.Scanf("%d %d\n", &outdoor, &indoor)
		curr -= outdoor
		curr += indoor
		if curr > max {
			max = curr
		}
	}
	fmt.Println(max)
}
