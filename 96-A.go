package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		input string
	)
	fmt.Scan(&input)
	if strings.Contains(input, "0000000") || strings.Contains(input, "1111111") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
