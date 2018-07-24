package jscraper

import (
	"fmt"
	"reflect"
)

// NumberP ...
func (n *Node) NumberP(fieldName string) (*float64, error) {
	value, ok := n.Value[fieldName]
	if !ok {
		return nil, notFoundError(n.Path + "." + fieldName)
	}

	if value == nil {
		return nil, nil
	}

	valueAsNumber, ok := value.(float64)
	if !ok {
		correctType := reflect.TypeOf(value).Name()
		return nil, wrongTypeError("number", correctType, n.Path+"."+fieldName)
	}

	return &valueAsNumber, nil
}

// Number ...
func (n *Node) Number(fieldName string) (float64, error) {
	value, err := n.NumberP(fieldName)
	if err != nil {
		return 0, err
	}

	if value == nil {
		return 0, nullTypeError("number", n.Path+"."+fieldName)
	}

	return *value, nil
}

// ArrayPOfNumbersP ...
func (n *Node) ArrayPOfNumbersP(fieldName string) ([]*float64, error) {
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

	var numbers []*float64

	for i, item := range items {
		if item == nil {
			numbers = append(numbers, nil)
			continue
		}

		numberItem, ok := item.(float64)

		if !ok {
			correctType := reflect.TypeOf(item).Name()
			return nil, arrayItemWrongTypeError(i, "float64", correctType, fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		numbers = append(numbers, &numberItem)
	}

	return numbers, nil
}

// ArrayPOfNumbers ...
func (n *Node) ArrayPOfNumbers(fieldName string) ([]float64, error) {
	numbers, err := n.ArrayPOfNumbersP(fieldName)
	if err != nil {
		return nil, err
	}

	if numbers == nil {
		return nil, nil
	}

	var derefNumbers []float64

	for i, numberItem := range numbers {
		if numberItem == nil {
			return nil, arrayItemNullTypeError(i, "float64", fmt.Sprintf("%s.%s[%d]", n.Path, fieldName, i))
		}

		derefNumbers = append(derefNumbers, *numberItem)
	}

	return derefNumbers, nil
}

// ArrayOfNumbersP ...
func (n *Node) ArrayOfNumbersP(fieldName string) ([]*float64, error) {
	numbers, err := n.ArrayPOfNumbersP(fieldName)
	if err != nil {
		return nil, err
	}

	if numbers == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return numbers, nil
}

// ArrayOfNumbers ...
func (n *Node) ArrayOfNumbers(fieldName string) ([]float64, error) {
	numbers, err := n.ArrayPOfNumbers(fieldName)
	if err != nil {
		return nil, err
	}

	if numbers == nil {
		return nil, nullTypeError("array", n.Path+"."+fieldName)
	}

	return numbers, nil
}
