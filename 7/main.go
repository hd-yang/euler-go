package main

import "fmt"

var prime [10001]int

func main() {
	prime[0] = 2
	i := 1
	num := 2
	for i < 10001 {
		num++
		var j int
		for j = 0; j < i; j++ {
			if num%prime[j] == 0 {
				break
			}
		}
		if j == i {
			prime[i] = num
			i++
		}
	}
	fmt.Println(prime[10000])
}
