package main

import (
	"fmt"
)

var n int
var sum [60]int
var dp [60]int
var pie [60]int

func main() {
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Scan(&pie[i])
	}
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + pie[i]
		if dp[i-1] > sum[i-1]-dp[i-1]+pie[i] {
			dp[i] = dp[i-1]
		} else {
			dp[i] = sum[i-1] - dp[i-1] + pie[i]
		}
	}
	fmt.Println(sum[n]-dp[n], dp[n])
}
