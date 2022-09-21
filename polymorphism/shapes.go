package polymorphism

import "fmt"

type Shape interface {
	Render()
}

type Circle struct{}

func (c Circle) Render() {
	// this is how to use interface functions
	fmt.Println("Circle method is implementing the render function")
}

type Square struct{}

func (c Square) Render() {
	// this is how to use interface functions
	fmt.Println("Square method is implementing the render function")
}
