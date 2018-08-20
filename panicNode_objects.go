package jscraper

// IsObject ...
func (n *PanicNode) IsObject(fieldName string) bool {
	return n.Node.IsObject(fieldName)
}

// ObjectP ...
func (n *PanicNode) ObjectP(fieldName string) *PanicNode {
	node, err := n.Node.ObjectP(fieldName)
	panicOnErr(err)

	if node == nil {
		return nil
	}
	return node.UsePanicAPI()
}

// Object ...
func (n *PanicNode) Object(fieldName string) PanicNode {
	node, err := n.Node.Object(fieldName)
	panicOnErr(err)

	return *node.UsePanicAPI()
}

// ArrayPOfObjectsP ...
func (n *PanicNode) ArrayPOfObjectsP(fieldName string) []*PanicNode {
	nodes, err := n.Node.ArrayPOfObjectsP(fieldName)
	panicOnErr(err)

	if nodes == nil {
		return nil
	}

	var pNodes []*PanicNode
	for i := range nodes {
		node := nodes[i]

		if node == nil {
			pNodes = append(pNodes, nil)
		} else {
			pNodes = append(pNodes, node.UsePanicAPI())
		}
	}
	return pNodes
}

// ArrayPOfObjects ...
func (n *PanicNode) ArrayPOfObjects(fieldName string) []PanicNode {
	nodes, err := n.Node.ArrayPOfObjects(fieldName)
	panicOnErr(err)

	if nodes == nil {
		return nil
	}

	var pNodes []PanicNode
	for i := range nodes {
		pNode := nodes[i].UsePanicAPI()
		pNodes = append(pNodes, *pNode)
	}
	return pNodes
}

// ArrayOfObjectsP ...
func (n *PanicNode) ArrayOfObjectsP(fieldName string) []*PanicNode {
	nodes, err := n.Node.ArrayOfObjectsP(fieldName)
	panicOnErr(err)

	var pNodes []*PanicNode
	for i := range nodes {
		node := nodes[i]

		if node == nil {
			pNodes = append(pNodes, nil)
		} else {
			pNodes = append(pNodes, node.UsePanicAPI())
		}
	}
	return pNodes
}

// ArrayOfObjects ...
func (n *PanicNode) ArrayOfObjects(fieldName string) []PanicNode {
	nodes, err := n.Node.ArrayOfObjects(fieldName)
	panicOnErr(err)

	var pNodes []PanicNode
	for i := range nodes {
		pNode := nodes[i].UsePanicAPI()
		pNodes = append(pNodes, *pNode)
	}

	return pNodes
}
