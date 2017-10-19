package main

import (
	"fmt"
	"sort"
)

var a [3]int
var t1 int
var t2 int

func main() {
	fmt.Scan(&a[0], &a[1], &a[2], &t1, &t2)
	a[1] = a[1] / 5
	a[2] = a[2] / 5
	s := a[:]
	sort.Ints(s)
	tt1 := 3
	tt2 := 3
	if t1 > s[0] && t1 <= s[1] {
		tt1 = 1
	} else if t1 > s[1] && t1 <= s[2] {
		tt1 = 2
	}
	if t2 > s[0] && t2 <= s[1] {
		tt2 = 1
	} else if t2 > s[1] && t2 <= s[2] {
		tt2 = 2
	}
	if tt1 == tt2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
