package jscraper_test

import (
	"fmt"
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestArrayPOfStringsP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfStringsP("arrayOfStrings")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 array items, but there were %d", len(actual))
		}

		for i, item := range actual {
			if item == nil {
				t.Fatalf("Not expecting nil item, but got one at index %d", i)
			}

			expected := fmt.Sprintf("String %d", i+1)
			if *item != expected {
				t.Fatalf("Expected %s at index %d but got %s", expected, i, *item)
			}
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayPOfStringsP("arrayOfStrings_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 4 {
			t.Fatalf("Expected 4 array items, but there were %d", len(actual))
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

			expected := fmt.Sprintf("String %d", i+1)
			if *item != expected {
				t.Fatalf("Expected %s at index %d but got %s", expected, i, *item)
			}
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayPOfStringsP("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfStringsP("arrayOfStrings_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfStringsP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfStringsP("nonsense")
	})
}

func TestArrayPOfStrings(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfStrings("arrayOfStrings")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 array items, but there were %d", len(actual))
		}

		for i, item := range actual {
			expected := fmt.Sprintf("String %d", i+1)
			if item != expected {
				t.Fatalf("Expected %s at index %d but got %s", expected, i, item)
			}
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		_, err := node.ArrayPOfStrings("arrayOfStrings_withNull")
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
		actual, err := node.ArrayPOfStrings("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfStrings("arrayOfStrings_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfStrings("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfStrings("nonsense")
	})
}

func TestArrayOfStringsP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfStringsP("arrayOfStrings")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 array items, but there were %d", len(actual))
		}

		for i, item := range actual {
			expected := fmt.Sprintf("String %d", i+1)
			if *item != expected {
				t.Fatalf("Expected %s at index %d but got %s", expected, i, *item)
			}
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayOfStringsP("arrayOfStrings_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 4 {
			t.Fatalf("Expected 4 array items, but there were %d", len(actual))
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

			expected := fmt.Sprintf("String %d", i+1)
			if *item != expected {
				t.Fatalf("Expected %s at index %d but got %s", expected, i, *item)
			}
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayOfStringsP("array_null")
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
		return node.ArrayOfStringsP("arrayOfStrings_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfStringsP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfStringsP("nonsense")
	})
}

func TestArrayOfStrings(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfStrings("arrayOfStrings")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 array items, but there were %d", len(actual))
		}

		for i, item := range actual {
			expected := fmt.Sprintf("String %d", i+1)
			if item != expected {
				t.Fatalf("Expected %s at index %d but got %s", expected, i, item)
			}
		}
	})

	t.Run("Invalid Array", func(t *testing.T) {
		_, err := node.ArrayOfStrings("arrayOfStrings_invalid")
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
		_, err := node.ArrayOfStrings("arrayOfStrings_withNull")
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
		actual, err := node.ArrayOfStrings("array_null")
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
		return node.ArrayOfStrings("arrayOfStrings_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfStrings("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfStrings("nonsense")
	})
}
