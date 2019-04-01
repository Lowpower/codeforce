package main

import (
	"fmt"
)

var num int

func main() {
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		var s string
		fmt.Scanln(&s)
		low := -1
		high := -1
		m := make(map[rune]bool)
		fail := false
		for _, j := range s {
			if int(j) < low || low == -1 {
				low = int(j)
			}
			if int(j) > high || high == -1 {
				high = int(j)
			}
			if m[j] != false {
				fail = true
				break
			}
			m[j] = true
		}
		if high-low+1 != len(s) || fail {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
