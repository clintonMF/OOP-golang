package structs

type Person struct {
	FirstName string //public field
	LastName  string //public field
}

/* in the struct above all the fields are public because they are in capitals
if we wish to keep them private we should start them in small letters */

// This is a struct method also called receiver functions
func (p Person) Walk() string {
	return p.FirstName + " " + p.LastName + " likes walking"
}

/* the struct below shows how we can deal with encapsulation in golang
since the fields are private we shouldn't be able to access them from another
package. the way around this is to use getter and setter functions
- this is implemented below */

type Person1 struct {
	firstName string //private field
	lastName  string // private field
}

// constructor
func NewPerson(fn, ln string) Person1 {
	return Person1{
		firstName: fn,
		lastName:  ln,
	}
}

func (p *Person1) SetFirstName(f string) {
	// this function sets the first name of the struct Person1
	// the argument is passed by reference (*)
	p.firstName = f
}

// getter function
func (p *Person1) FirstName() string {
	return p.firstName
}
