package main

import (
	"fmt"
)

// Incorrect naming convention (should be camelCase)
func Print_Message() {
	fmt.Println("Hello, GolangCI-Lint!")
}

// Unnecessary else block
func checkNumber(n int) string {
	if n > 0 {
		return "Positive"
	} else { // `revive` will flag this as an unnecessary else block
		return "Non-positive"
	}
}

// Unused function
func unusedFunction() { // `revive` will flag this as an unused function
	fmt.Println("This function is not used")
}

func main() {
	Print_Message()
	fmt.Println(checkNumber(-5))
}
