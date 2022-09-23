package comparison

type Person1 struct {
	firstName string //private field
	lastName  string // private field
	ids       []int
}

// constructor
func NewPerson(fn, ln string, ids []int) Person1 {
	return Person1{
		firstName: fn,
		lastName:  ln,
		ids:       ids,
	}
}

/* if we coompare 2 structs of person1 the compiler will
throw an error. this is because the "==" operator in golang
cannot compare dynamic data structures such as maps and slices.

the way to compare both is a bit long and tedious as we have to
create a function to compare slices and then create a function
to compare the structs. the functions are shown below.
*/

func EqualSlice(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}

func (p1 *Person1) Equals(p2 *Person1) bool {
	fn := p1.firstName == p2.firstName
	ln := p1.lastName == p2.lastName
	slice := EqualSlice(p1.ids, p2.ids)

	if fn && ln && slice {
		return true
	}

	return false
}
