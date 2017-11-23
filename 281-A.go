package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scanf("%s\n", &str)
	sstr := []byte(str)
	if sstr[0] >= 'a' && sstr[0] <= 'z' {
		sstr[0] -= 32
	}
	fmt.Println(string(sstr))
}
