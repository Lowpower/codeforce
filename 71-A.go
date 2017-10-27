package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var num int
	var output bytes.Buffer
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		var input string
		fmt.Scanln(&input)
		if len(input) <= 10 {
			fmt.Println(input)
		} else {
			inputByte := []byte(input)
			output.WriteByte(inputByte[0])
			output.WriteString(strconv.Itoa(len(input) - 2))
			output.WriteByte(inputByte[len(input)-1])
			fmt.Println(output.String())
		}
		output.Reset()
	}
}
