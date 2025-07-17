package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	fmt.Println(solve(k, nums))
}

func solve(k int, nums []int) int {
	difference := 0
	for i := 0; i < k; i++ {
		minNum := 0
		maxDig := 0
		elemToFix := 0
		for j, el := range nums {
			fixNum, fixDig := getFixDigit(el, getDigitNumber(el))
			if fixDig > maxDig {
				maxDig = fixDig
				minNum = fixNum
				elemToFix = j
			} else if fixDig == maxDig {
				if fixNum < minNum {
					minNum = fixNum
					elemToFix = j
				}
			}
		}
		if maxDig != 0 {
			var dif int
			dif, nums[elemToFix] = fixElem(nums[elemToFix], minNum, maxDig)
			difference += dif
		}
	}
	return difference
}

func getDigitNumber(num int) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}

func getFixDigit(number int, digit int) (fixNumber int, fixDigit int) {
	for digit-1 >= 0 {
		a := number / pow(10, digit-1) % 10
		if a != 9 {
			return a, digit
		} else {
			digit--
		}
	}
	return 0, 0
}

func pow(num int, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= num
	}
	return res
}

func fixElem(num int, fixNumber int, fixDigit int) (int, int) {
	res := num/pow(10, fixDigit)*pow(10, fixDigit) + 9*pow(10, fixDigit-1) + num%pow(10, fixDigit-1)
	difference := (9 - fixNumber) * pow(10, fixDigit-1)
	return difference, res
}
