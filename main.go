package main

import (
	"fmt"

	"example.com/go-oop/composition"
	"example.com/go-oop/polymorphism"
	"example.com/go-oop/structs"
)

func main() {
	// struct Person fields can be accessed in another package because they are public
	p := structs.Person{
		FirstName: "goku",
		LastName:  "son",
	}

	s := p.Walk()

	fmt.Println(p, s)

	/* struct Person1 fields are private and hence can't be directly accessed
	in a different package, so we use setter function  as seen below */
	p1 := structs.Person1{}

	p1.SetFirstName("kakarot")

	fmt.Println(p1)
	fmt.Println(p1.FirstName())

	//using constructor
	p11 := structs.NewPerson("goku", "son")

	fmt.Println(p11)

	// implementing polymorphism
	var c polymorphism.Shape = polymorphism.Circle{}
	var d polymorphism.Shape = polymorphism.Square{}

	c.Render()
	d.Render()

	// composition
	mercedes := composition.NewCar("bmw", "V4 engine", "4 wheels", 23, 4)

	fmt.Println(mercedes)

	//getting the first field (unique fields)
	fmt.Println("Unique field: ", mercedes.HP())
	fmt.Println("Unique field: ", mercedes.WheelDimension())

	//getting the second field (non unique fields)
	fmt.Println("Unique field: ", mercedes.EngineName())
	fmt.Println("Unique field: ", mercedes.WheelName())

	/* comparable

	the lines of code below looks at comparing structs
	*/

	p2 := structs.NewPerson("goku", "son")
	p3 := structs.NewPerson("goku", "son")

	if p2 == p3 {
		// this checks if both structs are equal
		fmt.Println("both structs are equal")
	} else {
		fmt.Println("both structs are not equal")
	}

	/*
		the direct comparison made between structs p2 and p3 is possible
		because they only contain primitive data structures in this case string.
		If they contained non primitive/ dynamic data structure such as slice and
		maps, direct comparism can't be made i.e the compiler would show error.

		there is a way around this. which is shown in the comparison.go file */

}
