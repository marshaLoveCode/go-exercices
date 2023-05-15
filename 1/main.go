package main

import "fmt"

func main() {
	fmt.Println(multOrSum(20, 30))
	fmt.Println(multOrSum(40, 30))
}

func multOrSum(x int, y int) int {
	result := x * y
	if result > 1000 {
		sum := x + y
		return sum
	}
	return result
}
