package jscraper

// NumberP ...
func (n *PanicNode) NumberP(fieldName string) *float64 {
	v, err := n.Node.NumberP(fieldName)
	panicOnErr(err)

	return v
}

// Number ...
func (n *PanicNode) Number(fieldName string) float64 {
	v, err := n.Node.Number(fieldName)
	panicOnErr(err)

	return v
}

// ArrayPOfNumbersP ...
func (n *PanicNode) ArrayPOfNumbersP(fieldName string) []*float64 {
	v, err := n.Node.ArrayPOfNumbersP(fieldName)
	panicOnErr(err)

	return v
}

// ArrayPOfNumbers ...
func (n *PanicNode) ArrayPOfNumbers(fieldName string) []float64 {
	v, err := n.Node.ArrayPOfNumbers(fieldName)
	panicOnErr(err)

	return v
}

// ArrayOfNumbersP ...
func (n *PanicNode) ArrayOfNumbersP(fieldName string) []*float64 {
	v, err := n.Node.ArrayOfNumbersP(fieldName)
	panicOnErr(err)

	return v
}

// ArrayOfNumbers ...
func (n *PanicNode) ArrayOfNumbers(fieldName string) []float64 {
	v, err := n.Node.ArrayOfNumbers(fieldName)
	panicOnErr(err)

	return v
}
