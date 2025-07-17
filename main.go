package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	students := make([]int, n)
	for i := 0; i < n; i++ {

		fmt.Scan(&students[i])

	}

	fmt.Println(solve(n, students))
}

func solve(n int, students []int) (int, int) {
	even := 0
	odd := 0

	for i := 0; i < n; i += 2 {
		if students[i]%2 != 1 {
			if odd != 0 {
				return -1, -1
			}
			odd = i + 1
		}
	}
	for i := 1; i < n; i += 2 {
		if students[i]%2 != 0 {
			if even != 0 {
				return -1, -1
			}
			even = i + 1
		}
	}
	if even == 0 && odd == 0 {
		if n >= 3 {
			return 1, 3
		} else {
			return -1, -1
		}
	}
	if even == 0 || odd == 0 {
		return -1, -1
	}

	if odd > even {
		return even, odd
	}

	return odd, even
}
