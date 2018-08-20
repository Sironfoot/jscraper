package jscraper

import (
	"fmt"
	"reflect"
)

// IsBoolean ...
func (n *Node) IsBoolean(fieldName string) bool {
	_, err := n.Boolean(fieldName)
	return err == nil
}

// BooleanP ...
func (n *Node) BooleanP(fieldName string) (*bool, error) {
	value, ok := n.Value[fieldName]
	if !ok {
		return nil, notFoundError(n.Path + "." + fieldName)
	}

	if value == nil {
		return nil, nil
	}

	valueAsNBoolean, ok := value.(bool)
	if !ok {
		correctType := reflect.TypeOf(value).Name()
		return nil, wrongTypeError("boolean", correctType, n.Path+"."+fieldName)
	}

	return &valueAsNBoolean, nil
}

// Boolean ...
func (n *Node) Boolean(fieldName string) (bool, error) {
	value, err := n.BooleanP(fieldName)
	if err != nil {
		return false, err
	}

	if value == nil {
		return false, nullTypeError("boolean", n.Path+"."+fieldName)
	}

	return *value, nil
}

// ArrayPOfBooleansP ...
func (n *Node) ArrayPOfBooleansP(fieldName string) ([]*bool, error) {
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

	var booleans []*bool

	for i, item := range items {
		if item == nil {
			booleans = append(booleans, nil)
			continue
		}

		boolItem, ok := item.(bool)

		if !ok {
			correctType := reflect.TypeOf(item).Name()
			return nil, arrayItemWrongTypeError(i, "bool", correctType, fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		booleans = append(booleans, &boolItem)
	}

	return booleans, nil
}

// ArrayPOfBooleans ...
func (n *Node) ArrayPOfBooleans(fieldName string) ([]bool, error) {
	booleans, err := n.ArrayPOfBooleansP(fieldName)
	if err != nil {
		return nil, err
	}

	if booleans == nil {
		return nil, nil
	}

	var derefBooleans []bool

	for i, boolItem := range booleans {
		if boolItem == nil {
			return nil, arrayItemNullTypeError(i, "bool", fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		derefBooleans = append(derefBooleans, *boolItem)
	}

	return derefBooleans, nil
}

// ArrayOfBooleansP ...
func (n *Node) ArrayOfBooleansP(fieldName string) ([]*bool, error) {
	booleans, err := n.ArrayPOfBooleansP(fieldName)
	if err != nil {
		return nil, err
	}

	if booleans == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return booleans, nil
}

// ArrayOfBooleans ...
func (n *Node) ArrayOfBooleans(fieldName string) ([]bool, error) {
	booleans, err := n.ArrayPOfBooleans(fieldName)
	if err != nil {
		return nil, err
	}

	if booleans == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return booleans, nil
}
