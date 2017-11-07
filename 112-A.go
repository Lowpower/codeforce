package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	var b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println(strings.Compare(strings.ToLower(a), strings.ToLower(b)))
}
