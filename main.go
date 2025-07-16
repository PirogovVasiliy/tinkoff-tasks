package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(solve(n))
}

func solve(n int) int {
	step := 1
	count := 0
	for step < n {
		step *= 2
		count++
	}
	return count
}
