package jscraper

// BooleanP ...
func (n *PanicNode) BooleanP(fieldName string) *bool {
	v, err := n.Node.BooleanP(fieldName)
	panicOnErr(err)

	return v
}

// Boolean ...
func (n *PanicNode) Boolean(fieldName string) bool {
	v, err := n.Node.Boolean(fieldName)
	panicOnErr(err)

	return v
}

// ArrayPOfBooleansP ...
func (n *PanicNode) ArrayPOfBooleansP(fieldName string) []*bool {
	v, err := n.Node.ArrayPOfBooleansP(fieldName)
	panicOnErr(err)

	return v
}

// ArrayPOfBooleans ...
func (n *PanicNode) ArrayPOfBooleans(fieldName string) []bool {
	v, err := n.Node.ArrayPOfBooleans(fieldName)
	panicOnErr(err)

	return v
}

// ArrayOfBooleansP ...
func (n *PanicNode) ArrayOfBooleansP(fieldName string) []*bool {
	v, err := n.Node.ArrayOfBooleansP(fieldName)
	panicOnErr(err)

	return v
}

// ArrayOfBooleans ...
func (n *PanicNode) ArrayOfBooleans(fieldName string) []bool {
	v, err := n.Node.ArrayOfBooleans(fieldName)
	panicOnErr(err)

	return v
}
