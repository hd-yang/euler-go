package main

import "fmt"

func main() {
	num := 600851475143
	f := 2
	for f < num {
		for num%f == 0 {
			num /= f
		}
		f++
	}
	fmt.Println(num)
}
