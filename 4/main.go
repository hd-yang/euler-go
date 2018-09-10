package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, j, a, b, max := 0, 0, 0, 0, 0
	for i = 999; i > 0; i-- {
		for j = 999; j > 0; j-- {
			if isPal(i * j) {
				if i*j > max {
					max = i * j
					a, b = i, j
				}
				break
			}
		}
	}
	// for i = 1; i < 1000; i++ {
	// 	for j = 1; j < 1000; j++ {
	// 		if isPal(i*j) && i*j > max {
	// 			max = i * j
	// 			a, b = i, j
	// 		}
	// 	}
	// }
	fmt.Println("max = ", max, " a = ", a, " b = ", b)
}

func isPal(num int) bool {
	s := strconv.Itoa(num)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}
