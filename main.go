package main

import (
	"fmt"

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

}
