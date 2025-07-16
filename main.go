package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(checkExpenses(a, b, c, d))
}

func checkExpenses(a int, b int, c int, d int) int {
	if d > b {
		return a + (d-b)*c
	} else {
		return a
	}
}
