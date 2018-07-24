package jscraper_test

import (
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestArrayPOfNumbersP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfNumbersP("arrayOfNumbers")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 6 {
			t.Fatalf("Expected 6 array items, but there were %d", len(actual))
		}

		for i, item := range actual {
			if item == nil {
				t.Fatalf("Not expecting nil item, but got one at index %d", i)
			}
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayPOfNumbersP("arrayOfNumbers_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 7 {
			t.Fatalf("Expected 7 array items, but there were %d", len(actual))
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

		for _, item := range actual {
			if item == nil {
				if foundNullItem {
					t.Fatalf("Only one of the items in the array should be null")
				}
				foundNullItem = true
				continue
			}
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayPOfNumbersP("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfNumbersP("arrayOfNumbers_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfNumbersP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfNumbersP("nonsense")
	})
}

func TestArrayPOfNumbers(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfNumbers("arrayOfNumbers")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		if len(actual) != 6 {
			t.Fatalf("Expected 6 array items, but there were %d", len(actual))
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		_, err := node.ArrayPOfNumbers("arrayOfNumbers_withNull")
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
		actual, err := node.ArrayPOfNumbers("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfNumbers("arrayOfNumbers_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfNumbers("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfNumbers("nonsense")
	})
}

func TestArrayOfNumbersP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfNumbersP("arrayOfNumbers")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 6 {
			t.Fatalf("Expected 6 array items, but there were %d", len(actual))
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayOfNumbersP("arrayOfNumbers_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 7 {
			t.Fatalf("Expected 7 array items, but there were %d", len(actual))
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

		for _, item := range actual {
			if item == nil {
				if foundNullItem {
					t.Fatalf("Only one of the items in the array should be null")
				}
				foundNullItem = true
				continue
			}
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayOfNumbersP("array_null")
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
		return node.ArrayOfNumbersP("arrayOfNumbers_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfNumbersP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfNumbersP("nonsense")
	})
}

func TestArrayOfNumbers(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfNumbers("arrayOfNumbers")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 6 {
			t.Fatalf("Expected 6 array items, but there were %d", len(actual))
		}
	})

	t.Run("Invalid Array", func(t *testing.T) {
		_, err := node.ArrayOfNumbers("arrayOfNumbers_invalid")
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
		_, err := node.ArrayOfNumbers("arrayOfNumbers_withNull")
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
		actual, err := node.ArrayOfNumbers("array_null")
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
		return node.ArrayOfNumbers("arrayOfNumbers_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfNumbers("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfNumbers("nonsense")
	})
}
