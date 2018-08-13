package jscraper

// PanicNode ...
type PanicNode struct {
	Node *Node
}

// HasProperty ...
func (n *PanicNode) HasProperty(fieldName string) bool {
	return n.Node.HasProperty(fieldName)
}

// GetRoot ...
func (n *PanicNode) GetRoot() *PanicNode {
	root := n.Node.GetRoot()

	panicNode := PanicNode{
		Node: root,
	}

	return &panicNode
}
