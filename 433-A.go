package main

import (
	"fmt"
)

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	if a < b {
		return gcd(b, a)
	} else {
		return gcd(b, a%b)
	}
}

var num int

func main() {
	fmt.Scan(&num)
	div := num / 2
	for i := div; div >= 1; i-- {
		if gcd(i, num-i) == 1 {
			fmt.Println(i, num-i)
			break
		}
	}
}
