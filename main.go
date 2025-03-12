package main

import (
	"fmt"
)

// Correct naming convention (camelCase)
func PrintMessage() {
	fmt.Println("Hello, GolangCI-Lint!")
	fmt.Println("Hello, GolangCI-Lint!")

}

// Simplified without unnecessary else block
func checkNumber(n int, t int) string {
	if n > 0 {
		return "Positive"
	}
	return "Non-positive"
}

func main() {
	PrintMessage()
	fmt.Println(checkNumber(-5, 0))
}
