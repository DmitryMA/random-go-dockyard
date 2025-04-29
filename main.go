package main

import (
	"fmt"
)

func main() {
	age := 23

	if age >= 24 {
		fmt.Printf("You are adult")
	} else if age < 23 {
		fmt.Printf("You are child")
	} else {
		fmt.Printf("Someone else")
	}

	day := "Monday"

	switch day {
		case "Monday":
			fmt.Printf("Start of the week")
			fallthrough
		case "Thursday":
			fmt.Printf("mid week")
		default: fmt.Printf("other day")
	}
}

func add(a int, b int) int {
	return a + b
}

func calculate(a, b int) (int, int) {
	return a + b, a * b
}