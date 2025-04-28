package main

import (
	"fmt"
)

func main() {
	var name string = "Dmitri"
	fmt.Printf("Hello Words %s\n", name)
	age := 34
	fmt.Printf("My ages: %d\n", age)

	var city string

	city = "Novi Sad"
	fmt.Printf("My City: %s\n", city)

	var country, continent string = "Serbia", "Europe"
	fmt.Printf("My Country is %s in the Continent %s \n", country, continent)

	var (
		isEmployed bool = true
		salary int = 50000
		position string = "developer"
	)

	fmt.Printf("Is Employed %t, My Salary %d, My Position %s \n", isEmployed, salary, position)

	const pi = 3.14
	const (
		Monday = 1
		Tuesday = 2
		Thursday = 3
	)

	fmt.Printf("%d , %d, %d ", Monday, Tuesday, Thursday)

	const (
		Jan int = iota + 1 // 1
		Feb   // 2
		March // 3
		Apr   // 4
	)

	fmt.Printf("%d , %d, %d, %d \n", Jan, Feb, March, Apr)

	res := add(2, 6)
	fmt.Printf("%d \n", res)

	sum, product := calculate(10, 30)
	fmt.Printf("%d %d \n", sum, product)
}

func add(a int, b int) int {
	return a + b
}

func calculate(a, b int) (int, int) {
	return a + b, a * b
}