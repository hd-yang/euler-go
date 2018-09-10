package main

import (
	"fmt"
)

func main() {
	ans := 1
	i := 0
	for i = 2; i <= 20; i++ {
		ans = lcm(ans, i)
	}
	fmt.Println(ans)
}

// get the greatest common divisor of a and b
func gcd(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// get the Least Common Multiple of a and b
func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}
