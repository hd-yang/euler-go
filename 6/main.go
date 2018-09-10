package main

import (
	"fmt"
)

func main() {
	sum1 := 0
	sum2 := 0
	i := 0
	for i = 1; i <= 100; i++ {
		sum1 += i * i
	}
	sum2 = (1 + 100) * 100 / 2
	sum2 = sum2 * sum2
	fmt.Println(sum2 - sum1)
}
