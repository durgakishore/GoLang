package organization

//Identifiable interafce
type Identifiable interface {
	ID() string
}

//Person struct
type Person struct {
}

//ID is an Identifiable interface function
func (p Person) ID() string {

	return "1234"
}
