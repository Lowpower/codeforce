package main

import (
	"fmt"
)

var lucky = []int{4, 7, 44, 77, 47, 74, 444, 777, 477, 747, 774, 744, 474, 447}

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	var a = num
	var fuck = true
	for {
		if a == 0 {
			break
		}
		if a%10 == 4 || a%10 == 7 {
			fuck = true
			a = a / 10
		} else {
			fuck = false
			break
		}
	}
	if !fuck {
		for _, i := range lucky {
			if num%i == 0 {
				fuck = true
				break
			}
		}
	}
	if fuck {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
