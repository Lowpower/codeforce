package main

import (
	"fmt"
	"sort"
)

var num int

func main() {
	fmt.Scanln(&num)
	a := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scan(&a[i])
	}
	for i := num; i >= 1; i /= 2 {
		for j := 0; j < num; j += i {
			if sort.IntsAreSorted(a[j : j+i]) {
				fmt.Println(i)
				return
			}
		}
	}
}
