package jscraper

// StringP ...
func (n *PanicNode) StringP(fieldName string) *string {
	v, err := n.Node.StringP(fieldName)
	panicOnErr(err)

	return v
}

// String ...
func (n *PanicNode) String(fieldName string) string {
	v, err := n.Node.String(fieldName)
	panicOnErr(err)

	return v
}

// ArrayPOfStringsP ...
func (n *PanicNode) ArrayPOfStringsP(fieldName string) []*string {
	v, err := n.Node.ArrayPOfStringsP(fieldName)
	panicOnErr(err)

	return v
}

// ArrayPOfStrings ...
func (n *PanicNode) ArrayPOfStrings(fieldName string) []string {
	v, err := n.Node.ArrayPOfStrings(fieldName)
	panicOnErr(err)

	return v
}

// ArrayOfStringsP ...
func (n *PanicNode) ArrayOfStringsP(fieldName string) []*string {
	v, err := n.Node.ArrayOfStringsP(fieldName)
	panicOnErr(err)

	return v
}

// ArrayOfStrings ...
func (n *PanicNode) ArrayOfStrings(fieldName string) []string {
	v, err := n.Node.ArrayOfStrings(fieldName)
	panicOnErr(err)

	return v
}
