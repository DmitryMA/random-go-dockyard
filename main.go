package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	// person := Person{
	// 	Name: "Dmitrii",
	// 	Age: 23,
	// }	

	// employee := struct {
	// 	name string
	// 	id int
	// 	} {
	// 		name: "Dmitrii",
	// 		id: 1,
	// 	}



	type Address struct {
		Street string
		City string
	}

	type Contact struct {
		Name string
		Address Address
		Phone string
	}

	contact := Contact {
		Name: "Mark",
		Address: Address {
			Street: "Lenin Street",
			City: "Moscow",
		},
		Phone: "+8 0000000",
	}

	fmt.Printf("%+v \n", contact)
}