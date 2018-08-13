package jscraper

// ObjectP ...
func (n *PanicNode) ObjectP(fieldName string) *PanicNode {
	v, err := n.Node.ObjectP(fieldName)
	panicOnErr(err)

	return v.UsePanicAPI()
}

// Object ...
func (n *PanicNode) Object(fieldName string) PanicNode {
	v, err := n.Node.Object(fieldName)
	panicOnErr(err)

	return *v.UsePanicAPI()
}

// ArrayPOfObjectsP ...
func (n *PanicNode) ArrayPOfObjectsP(fieldName string) []*PanicNode {
	vs, err := n.Node.ArrayPOfObjectsP(fieldName)
	panicOnErr(err)

	var pv []*PanicNode
	for _, v := range vs {
		pv = append(pv, v.UsePanicAPI())
	}
	return pv
}

// ArrayPOfObjects ...
func (n *PanicNode) ArrayPOfObjects(fieldName string) []PanicNode {
	vs, err := n.Node.ArrayPOfObjects(fieldName)
	panicOnErr(err)

	var pv []PanicNode
	for _, v := range vs {
		pv = append(pv, *v.UsePanicAPI())
	}
	return pv
}

// ArrayOfObjectsP ...
func (n *PanicNode) ArrayOfObjectsP(fieldName string) []*PanicNode {
	vs, err := n.Node.ArrayOfObjectsP(fieldName)
	panicOnErr(err)

	var pv []*PanicNode
	for _, v := range vs {
		pv = append(pv, v.UsePanicAPI())
	}
	return pv
}

// ArrayOfObjects ...
func (n *PanicNode) ArrayOfObjects(fieldName string) []PanicNode {
	vs, err := n.Node.ArrayOfObjects(fieldName)
	panicOnErr(err)

	var pv []PanicNode
	for _, v := range vs {
		pv = append(pv, *v.UsePanicAPI())
	}
	return pv
}
