package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanf("%d\n", &num)
	var str string
	fmt.Scanf("%s\n", &str)
	strbyte := []byte(str)
	sum := 0
	for i := 0; i < len(strbyte)-1; i++ {
		if strbyte[i+1] == strbyte[i] {
			sum++
		}
	}
	fmt.Println(sum)
}
