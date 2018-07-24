package jscraper

import (
	"fmt"
	"reflect"
)

// ObjectP ...
func (n *Node) ObjectP(fieldName string) (*Node, error) {
	value, ok := n.Value[fieldName]
	if !ok {
		return nil, notFoundError(n.Path + "." + fieldName)
	}

	if value == nil {
		return nil, nil
	}

	valueAsObject, ok := value.(map[string]interface{})
	if !ok {
		correctType := reflect.TypeOf(value).Name()
		return nil, wrongTypeError("object", correctType, n.Path+"."+fieldName)
	}

	node := Node{
		ParentNode: n,
		Path:       n.Path + "." + fieldName,
		Value:      valueAsObject,
	}

	return &node, nil
}

// Object ...
func (n *Node) Object(fieldName string) (Node, error) {
	node, err := n.ObjectP(fieldName)
	if err != nil {
		return Node{}, err
	}

	if node == nil {
		return Node{}, nullTypeError("object", n.Path+"."+fieldName)
	}

	return *node, nil
}

// ArrayPOfObjectsP ...
func (n *Node) ArrayPOfObjectsP(fieldName string) ([]*Node, error) {
	value, ok := n.Value[fieldName]
	if !ok {
		return nil, notFoundError(n.Path + "." + fieldName)
	}

	if value == nil {
		return nil, nil
	}

	items, ok := value.([]interface{})
	if !ok {
		correctType := reflect.TypeOf(value).Name()
		return nil, wrongTypeError("array", correctType, n.Path+"."+fieldName)
	}

	var nodes []*Node

	for i, item := range items {
		if item == nil {
			nodes = append(nodes, nil)
			continue
		}

		valueItem, ok := item.(map[string]interface{})

		if !ok {
			correctType := reflect.TypeOf(item).Name()
			return nil, arrayItemWrongTypeError(i, "object", correctType, fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		nodes = append(nodes, &Node{
			ParentNode: n,
			Path:       fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i),
			Value:      valueItem,
		})
	}

	return nodes, nil
}

// ArrayPOfObjects ...
func (n *Node) ArrayPOfObjects(fieldName string) ([]Node, error) {
	nodes, err := n.ArrayPOfObjectsP(fieldName)
	if err != nil {
		return nil, err
	}

	if nodes == nil {
		return nil, nil
	}

	var derefNodes []Node

	for i, node := range nodes {
		if node == nil {
			return nil, arrayItemNullTypeError(i, "object", fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		derefNodes = append(derefNodes, *node)
	}

	return derefNodes, nil
}

// ArrayOfObjectsP ...
func (n *Node) ArrayOfObjectsP(fieldName string) ([]*Node, error) {
	nodes, err := n.ArrayPOfObjectsP(fieldName)
	if err != nil {
		return nil, err
	}

	if nodes == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return nodes, nil
}

// ArrayOfObjects ...
func (n *Node) ArrayOfObjects(fieldName string) ([]Node, error) {
	nodes, err := n.ArrayPOfObjects(fieldName)
	if err != nil {
		return nil, err
	}

	if nodes == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return nodes, nil
}
