package main

import (
	"fmt"
	"sort"
	_ "strconv"
	"strings"
)

func main() {
	var a string
	var b string
	fmt.Scanln(&a)
	strArray := strings.Split(a, "+")
	sort.Strings(strArray)
	len := len(strArray)
	for i := 0; i < len-1; i++ {
		b += strArray[i]
		b += "+"
	}
	b += strArray[len-1]
	fmt.Println(b)
}
