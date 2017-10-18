package main

import (
	"fmt"
)

var s string
var num int
var str [120]string

func main() {
	fmt.Scan(&s)
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Scan(&str[i])
	}
	for i := 0; i < num; i++ {
		if str[i] == s {
			fmt.Println("YES")
			return
		}
		for j := 0; j < num; j++ {
			if str[i][0] == s[1] && str[j][1] == s[0] {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}
