package jscraper_test

import (
	"fmt"
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestArrayPOfObjectsP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfObjectsP("arrayOfObjects")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) == 0 {
			t.Fatalf("Expected array items, but there were none")
		}

		for i, item := range actual {
			actualString, err := item.String("string")
			if err != nil {
				t.Fatalf("Expecting a property called 'string' on the object at index %d, but got error %s", i, err)
			}

			expectedString := "Hello world"

			if actualString != expectedString {
				t.Fatalf("Expecting object's 'string' property to be %s at index %d but was %s", expectedString, i, actualString)
			}

			expectedPath := fmt.Sprintf(".arrayOfObjects[%d]", i)

			if item.Path != expectedPath {
				t.Fatalf("Expecting object's path to be %s at index %d but was %s", expectedPath, i, item.Path)
			}
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayPOfObjectsP("arrayOfObjects_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) == 0 {
			t.Fatalf("Expected array items, but there were none")
		}

		foundNullItem := false

		for _, item := range actual {
			if item == nil {
				foundNullItem = true
				break
			}
		}

		if !foundNullItem {
			t.Fatalf("One of the items in the array should be null")
		}

		foundNullItem = false

		for i, item := range actual {
			if item == nil {
				if foundNullItem {
					t.Fatalf("Only one of the items in the array should be null")
				}
				foundNullItem = true
				continue
			}

			actualString, err := item.String("string")
			if err != nil {
				t.Fatalf("Expecting a property called 'string' on the object at index %d, but got error %s", i, err)
			}

			expectedString := "Hello world"

			if actualString != expectedString {
				t.Fatalf("Expecting object's 'string' property to be %s at index %d but was %s", expectedString, i, actualString)
			}

			expectedPath := fmt.Sprintf(".arrayOfObjects_withNull[%d]", i)

			if item.Path != expectedPath {
				t.Fatalf("Expecting object's path to be %s at index %d but was %s", expectedPath, i, item.Path)
			}
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayPOfObjectsP("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfObjectsP("arrayOfObjects_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfObjectsP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfObjectsP("nonsense")
	})
}

func TestArrayPOfObjects(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfObjects("arrayOfObjects")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		if len(actual) == 0 {
			t.Fatalf("Expected array items, but there were none")
		}

		for i, item := range actual {
			actualString, err := item.String("string")
			if err != nil {
				t.Fatalf("Expecting a property called 'string' on the object at index %d, but got error %s", i, err)
			}

			expectedString := "Hello world"

			if actualString != expectedString {
				t.Fatalf("Expecting object's 'string' property to be %s at index %d but was %s", expectedString, i, actualString)
			}

			expectedPath := fmt.Sprintf(".arrayOfObjects[%d]", i)

			if item.Path != expectedPath {
				t.Fatalf("Expecting object's path to be %s at index %d but was %s", expectedPath, i, item.Path)
			}
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		_, err := node.ArrayPOfObjects("arrayOfObjects_withNull")
		if err == nil {
			t.Fatal("Expecting an error to be returned, but was nil")
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeArrayItemIsNull {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeArrayItemIsNull, nodeError.Type)
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayPOfObjects("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfObjects("arrayOfObjects_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfObjects("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfObjects("nonsense")
	})
}

func TestArrayOfObjectsP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfObjectsP("arrayOfObjects")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) == 0 {
			t.Fatalf("Expected array items, but there were none")
		}

		for i, item := range actual {
			actualString, err := item.String("string")
			if err != nil {
				t.Fatalf("Expecting a property called 'string' on the object at index %d, but got error %s", i, err)
			}

			expectedString := "Hello world"

			if actualString != expectedString {
				t.Fatalf("Expecting object's 'string' property to be %s at index %d but was %s", expectedString, i, actualString)
			}

			expectedPath := fmt.Sprintf(".arrayOfObjects[%d]", i)

			if item.Path != expectedPath {
				t.Fatalf("Expecting object's path to be %s at index %d but was %s", expectedPath, i, item.Path)
			}
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayOfObjectsP("arrayOfObjects_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) == 0 {
			t.Fatalf("Expected array items, but there were none")
		}

		foundNullItem := false

		for _, item := range actual {
			if item == nil {
				foundNullItem = true
				break
			}
		}

		if !foundNullItem {
			t.Fatalf("One of the items in the array should be null")
		}

		foundNullItem = false

		for i, item := range actual {
			if item == nil {
				if foundNullItem {
					t.Fatalf("Only one of the items in the array should be null")
				}
				foundNullItem = true
				continue
			}

			actualString, err := item.String("string")
			if err != nil {
				t.Fatalf("Expecting a property called 'string' on the object at index %d, but got error %s", i, err)
			}

			expectedString := "Hello world"

			if actualString != expectedString {
				t.Fatalf("Expecting object's 'string' property to be %s at index %d but was %s", expectedString, i, actualString)
			}

			expectedPath := fmt.Sprintf(".arrayOfObjects_withNull[%d]", i)

			if item.Path != expectedPath {
				t.Fatalf("Expecting object's path to be %s at index %d but was %s", expectedPath, i, item.Path)
			}
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayOfObjectsP("array_null")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil, returned value was '%v'", actual)
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeIsNull {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeIsNull, nodeError.Type)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayOfObjectsP("arrayOfObjects_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfObjectsP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfObjectsP("nonsense")
	})
}

func TestArrayOfObjects(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfObjects("arrayOfObjects")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) == 0 {
			t.Fatalf("Expected array items, but there were none")
		}

		for i, item := range actual {
			actualString, err := item.String("string")
			if err != nil {
				t.Fatalf("Expecting a property called 'string' on the object at index %d, but got error %s", i, err)
			}

			expectedString := "Hello world"

			if actualString != expectedString {
				t.Fatalf("Expecting object's 'string' property to be %s at index %d but was %s", expectedString, i, actualString)
			}

			expectedPath := fmt.Sprintf(".arrayOfObjects[%d]", i)

			if item.Path != expectedPath {
				t.Fatalf("Expecting object's path to be %s at index %d but was %s", expectedPath, i, item.Path)
			}
		}
	})

	t.Run("Invalid Array", func(t *testing.T) {
		_, err := node.ArrayOfObjects("arrayOfObjects_invalid")
		if err == nil {
			t.Fatal("Expecting an error to be returned, but was nil")
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeArrayItemWrongType {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeArrayItemWrongType, nodeError.Type)
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		_, err := node.ArrayOfObjects("arrayOfObjects_withNull")
		if err == nil {
			t.Fatal("Expecting an error to be returned, but was nil")
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeArrayItemIsNull {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeArrayItemIsNull, nodeError.Type)
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayOfObjects("array_null")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil, returned value was '%v'", actual)
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeIsNull {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeIsNull, nodeError.Type)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayOfObjects("arrayOfObjects_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfObjects("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfObjects("nonsense")
	})
}

func checkInvalidArray(t *testing.T, getNodes func() (interface{}, error)) {
	t.Run("Invalid Array", func(t *testing.T) {
		_, err := getNodes()
		if err == nil {
			t.Fatal("Expecting an error to be returned, but was nil")
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeArrayItemWrongType {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeArrayItemWrongType, nodeError.Type)
		}
	})
}

func checkArrayWrongType(t *testing.T, getNode func() (interface{}, error)) {
	t.Run("Wrong Type", func(t *testing.T) {
		data, err := getNode()

		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}

		if actual, ok := data.([]jscraper.Node); ok {
			if actual != nil {
				t.Fatalf("Expected returned value to be nil, but it was %v", actual)
			}
		}

		if actual, ok := data.([]*jscraper.Node); ok {
			if actual != nil {
				t.Fatalf("Expected returned value to be nil, but it was %v", actual)
			}
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeWrongType {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeWrongType, nodeError.Type)
		}
	})
}

func checkArrayNonExistentProperty(t *testing.T, getNode func() (interface{}, error)) {
	t.Run("Non-existent Property", func(t *testing.T) {
		data, err := getNode()
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if actual, ok := data.([]jscraper.Node); ok {
			if actual != nil {
				t.Fatalf("Expected returned value to be nil, but it was %v", actual)
			}
		}

		if actual, ok := data.([]*jscraper.Node); ok {
			if actual != nil {
				t.Fatalf("Expected returned value to be nil, but it was %v", actual)
			}
		}

		if nodeError.Type != jscraper.NodeErrorTypeNotFound {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeNotFound, nodeError.Type)
		}
	})
}
