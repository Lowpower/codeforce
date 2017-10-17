package main

import (
	"fmt"
)

var num int
var str string

func main() {
	fmt.Scan(&num)
	fmt.Scan(&str)
	if str[0] == 'S' && str[len(str)-1] == 'F' {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
