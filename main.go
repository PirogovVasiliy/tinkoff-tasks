package main

import (
	"fmt"
)

func main() {
	var l, r uint64
	fmt.Scan(&l, &r)

	fmt.Println(solve(l, r))
}

func solve(l uint64, r uint64) int {
	count := 0

	digitNumber := getDigitNumber(l)
	firstDigit := int(l / pow(10, digitNumber-1))
	for i := firstDigit; i <= 9; i++ {
		test := formTest(i, digitNumber)

		if test < l {
			continue
		} else if test < r {
			count++
		} else if test == r {
			return count + 1
		} else {
			return count
		}
	}

	for {
		digitNumber++
		for i := 1; i <= 9; i++ {
			test := formTest(i, digitNumber)

			if test < r {
				count++
			} else if test == r {
				return count + 1
			} else {
				return count
			}
		}
	}
}

func getDigitNumber(num uint64) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func formTest(digit int, digitNumbers int) uint64 {
	res := uint64(digit)
	for i := 0; i < digitNumbers-1; i++ {
		res = res*10 + uint64(digit)
	}
	return res
}

func pow(num uint64, b int) uint64 {
	res := uint64(1)
	for i := 0; i < b; i++ {
		res *= num
	}
	return res
}
