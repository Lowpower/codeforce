package main

import (
	"bytes"
	"fmt"
	"strings"
)

var byteMap = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
	"y": true,
}

func main() {
	var input string
	var output bytes.Buffer
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	inputByte := []byte(input)
	for _, value := range inputByte {
		if _, ok := byteMap[string(value)]; ok {
			continue
		}
		output.WriteString(".")
		output.WriteByte(value)
	}
	fmt.Println(output.String())
}
