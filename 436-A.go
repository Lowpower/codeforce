package main

import (
	"fmt"
)

var num int
var temp int

func main() {
	fmt.Scan(&num)
	res := make(map[int]int)
	for i := 0; i < num; i++ {
		fmt.Scan(&temp)
		if num, ok := res[temp]; ok {
			res[temp] = num + 1
		} else {
			res[temp] = 1
		}
	}
	if len(res) != 2 || res[temp] != num/2 {
		fmt.Println("NO")
		return
	}
	pre := -1
	for key, _ := range res {
		if pre == -1 {
			pre = key
		} else {
			fmt.Printf("yes\n%d %d\n", pre, key)
		}
	}
}
