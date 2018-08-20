package jscraper

import (
	"fmt"
	"reflect"
)

// IsString ...
func (n *Node) IsString(fieldName string) bool {
	_, err := n.String(fieldName)
	return err == nil
}

// StringP ...
func (n *Node) StringP(fieldName string) (*string, error) {
	value, ok := n.Value[fieldName]
	if !ok {
		return nil, notFoundError(n.Path + "." + fieldName)
	}

	if value == nil {
		return nil, nil
	}

	valueAsString, ok := value.(string)
	if !ok {
		correctType := reflect.TypeOf(value).Name()
		return nil, wrongTypeError("string", correctType, n.Path+"."+fieldName)
	}

	return &valueAsString, nil
}

// String ...
func (n *Node) String(fieldName string) (string, error) {
	value, err := n.StringP(fieldName)
	if err != nil {
		return "", err
	}

	if value == nil {
		return "", nullTypeError("string", n.Path+"."+fieldName)
	}

	return *value, nil
}

// ArrayPOfStringsP ...
func (n *Node) ArrayPOfStringsP(fieldName string) ([]*string, error) {
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

	var strings []*string

	for i, item := range items {
		if item == nil {
			strings = append(strings, nil)
			continue
		}

		stringItem, ok := item.(string)

		if !ok {
			correctType := reflect.TypeOf(item).Name()
			return nil, arrayItemWrongTypeError(i, "string", correctType, fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		strings = append(strings, &stringItem)
	}

	return strings, nil
}

// ArrayPOfStrings ...
func (n *Node) ArrayPOfStrings(fieldName string) ([]string, error) {
	strings, err := n.ArrayPOfStringsP(fieldName)
	if err != nil {
		return nil, err
	}

	if strings == nil {
		return nil, nil
	}

	var derefStrings []string

	for i, stringItem := range strings {
		if stringItem == nil {
			return nil, arrayItemNullTypeError(i, "string", fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		derefStrings = append(derefStrings, *stringItem)
	}

	return derefStrings, nil
}

// ArrayOfStringsP ...
func (n *Node) ArrayOfStringsP(fieldName string) ([]*string, error) {
	strings, err := n.ArrayPOfStringsP(fieldName)
	if err != nil {
		return nil, err
	}

	if strings == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return strings, nil
}

// ArrayOfStrings ...
func (n *Node) ArrayOfStrings(fieldName string) ([]string, error) {
	strings, err := n.ArrayPOfStrings(fieldName)
	if err != nil {
		return nil, err
	}

	if strings == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return strings, nil
}
