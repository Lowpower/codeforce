package main

import (
	"fmt"
)

var num int
var str string

func main() {
	fmt.Scan(&num)
	fmt.Scan(&str)
	max := 0
	pre := make(map[byte]bool)
	for i := 0; i < num; i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			pre[str[i]] = true
		} else {
			if len(pre) > max {
				max = len(pre)
			}
			pre = make(map[byte]bool)
		}
	}
	if len(pre) > max {
		max = len(pre)
	}
	fmt.Println(max)
}
