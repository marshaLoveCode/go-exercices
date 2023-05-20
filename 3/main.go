package main

import "fmt"

func main() {
	oddChar()
}

func oddChar() {
	var userInput string
	fmt.Scanln(&userInput)
	for i := 0; i < len(userInput); i++ {
		fmt.Println(string(userInput[i]))
	}
}
