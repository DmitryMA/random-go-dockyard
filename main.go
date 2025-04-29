package main

import "fmt"

func main() {
	numbers := [5] int{ 10, 20, 30, 40, 55}

	allNumbers := numbers[:]
	threeNumbers := numbers[0:3]

	fruits := []string{"7","8","9","10"}
	fruits = append(fruits, "11")
	fruits = append(fruits, "12", "13")

	otherFruits := []string{"20", "21", "22"}

	fruits = append(fruits, otherFruits...)

	fmt.Printf("%v \n %v \n %v", allNumbers, threeNumbers, fruits)

	for idx, fruit := range fruits {
		fmt.Printf("%d => %s \n", idx, fruit)
	}
}

func add(a int, b int) int {
	return a + b
}

func calculate(a, b int) (int, int) {
	return a + b, a * b
}