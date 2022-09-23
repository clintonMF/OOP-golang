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

}
