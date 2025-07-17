package main

import (
	"fmt"
)

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	floors := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&floors[i])
	}
	var employee int
	fmt.Scan(&employee)
	res := solve(n, t, floors, employee)
	fmt.Println(res)

}

func solve(n int, t int, staff []int, emmployee int) int {

	before := staff[emmployee-1] - staff[0]
	after := staff[n-1] - staff[emmployee-1]

	if before <= t || after <= t {
		return before + after
	} else if before < after {
		return before*2 + after
	} else {
		return after*2 + before
	}
}
