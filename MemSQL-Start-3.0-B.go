package main

import (
	"fmt"
	"math"
)

var n int

func main() {
	fmt.Scan(&n)
	sqt := math.Sqrt(float64(n))
	hight := int(sqt)
	width := 0
	if hight*hight == n {
		width = hight
	} else {
		width = hight + 1
		if hight*width < n {
			hight += 1
		}
	}
	fmt.Println(hight*2 + width*2)
}
