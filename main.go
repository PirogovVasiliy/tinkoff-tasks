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
	checkEdges := make([]int, n)

	for _, el := range students {
		checkEdges[el-1] += 1
	}

	flagOne := -1 //куда нужно перенаправить
	flagTwo := -1
	for i, el := range checkEdges {
		if el == 0 {
			if flagOne != -1 {
				return -1, -1
			} else {
				flagOne = i
			}
		} else if el == 2 {
			if flagTwo != -1 {
				return -1, -1
			} else {
				flagTwo = i
			}
		} else if el > 2 {
			return -1, -1
		}
	}

	if flagOne == -1 || flagTwo == -1 {
		return -1, -1
	}

	fstEdg := -1 //первая попытка перенаправить
	sndEdg := -1 // вторая попытка перенаправить
	for i, el := range students {
		if el == flagTwo+1 {
			if fstEdg == -1 {
				fstEdg = i
			} else {
				sndEdg = i
				break
			}
		}
	}

	if checkCycle(n, students, fstEdg, flagOne) {
		return fstEdg + 1, flagOne + 1
	} else if checkCycle(n, students, sndEdg, flagOne) {
		return sndEdg + 1, flagOne + 1
	} else {
		return -1, -1
	}

}

func checkCycle(n int, studentsOrigin []int, edgeFrom int, edgeTo int) bool {

	students := make([]int, n)
	copy(students, studentsOrigin)

	students[edgeFrom] = edgeTo + 1

	flagCycle := make([]bool, n)

	temp := 1
	flagCycle[temp-1] = true
	for i := 0; i < n; i++ {
		flagCycle[temp-1] = true
		temp = students[temp-1]
	}

	for _, el := range flagCycle {
		if !el {
			return false
		}
	}
	return true

}

func changer(students []int) bool {
	students[0] = 10
	return true
}
