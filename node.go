package jscraper

// Node ...
type Node struct {
	ParentNode *Node
	Path       string
	Value      map[string]interface{}
}

// GetRoot ...
func (n *Node) GetRoot() *Node {
	if n.ParentNode == nil {
		return n
	}

	parent := n.ParentNode

	for {
		if parent.ParentNode == nil {
			break
		}

		parent = parent.ParentNode
	}

	return parent
}

// HasProperty ...
func (n *Node) HasProperty(fieldName string) bool {
	_, ok := n.Value[fieldName]
	return ok
}
