package main

import "fmt"

func main() {
	sumWithPrevious(10)
}

func sumWithPrevious(x int) {
	currentNum, previousNum := 0, 0

	for i := 0; i < x; i++ {
		currentNum = previousNum + 1
		fmt.Printf(
			"Current Number %d Previous Number  %d Sum:%d  \n",
			currentNum,
			previousNum,
			currentNum+previousNum)
		previousNum = i
	}
}
