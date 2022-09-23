package composition

type engine struct {
	hp   int
	name string
}

/* creating getter functions for the engine struct
the first getter function for the hp field doesn't have
to be made special as the field name is unique to engine.

the second getter function which is for the name field
needs some more upgrades to avoid confusion as it shares
wheel struct also has a field with the same name.
*/

// first field
func (e engine) HP() int {
	return e.hp
}

// second field - needs extra function to be called by the car struct
func (e engine) Name() string {
	return e.name
}

type wheel struct {
	wheelDimension int
	name           string
}

// first field
func (w wheel) WheelDimension() int {
	return w.wheelDimension
}

// second field - needs extra function to be called by the car struct
func (w wheel) Name() string {
	return w.name
}

type Car struct {
	Name string
	engine
	wheel
}

func NewCar(N, en, wn string, hp, wd int) Car {
	return Car{
		Name: N,
		engine: engine{
			hp:   hp,
			name: en,
		},
		wheel: wheel{
			wheelDimension: wd,
			name:           wn,
		},
	}
}

// the extra function needed to get the car field name
func (c Car) EngineName() string {
	return c.engine.Name()
}

// the extra function needed to get the wheel field name
func (c Car) WheelName() string {
	return c.wheel.Name()
}
