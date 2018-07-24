package jscraper_test

import (
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestArrayPOfBooleansP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfBooleansP("arrayOfBooleans")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 array items, but there were %d", len(actual))
		}

		for i, item := range actual {
			if item == nil {
				t.Fatalf("Not expecting nil item, but got one at index %d", i)
			}
		}

		if *actual[0] != true {
			t.Fatalf("First array item should be true but was %v", *actual[0])
		}

		if *actual[1] != false {
			t.Fatalf("Second array item should be false but was %v", *actual[1])
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayPOfBooleansP("arrayOfBooleans_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 array items, but there were %d", len(actual))
		}

		if *actual[0] != true {
			t.Fatalf("First array item should be true but was %v", *actual[0])
		}

		if *actual[1] != false {
			t.Fatalf("Second array item should be false but was %v", *actual[1])
		}

		if actual[2] != nil {
			t.Fatalf("Third array item should be nil but was %v", *actual[2])
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayPOfBooleansP("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfBooleansP("arrayOfBooleans_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfBooleansP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfBooleansP("nonsense")
	})
}

func TestArrayPOfBooleans(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayPOfBooleans("arrayOfBooleans")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 array items, but there were %d", len(actual))
		}

		if actual[0] != true {
			t.Fatalf("First array item should be true but was %v", actual[0])
		}

		if actual[1] != false {
			t.Fatalf("Second array item should be false but was %v", actual[1])
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		_, err := node.ArrayPOfBooleans("arrayOfBooleans_withNull")
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
		actual, err := node.ArrayPOfBooleans("array_null")
		if err != nil {
			t.Fatalf("Not expecting an error but got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", actual)
		}
	})

	checkInvalidArray(t, func() (interface{}, error) {
		return node.ArrayPOfBooleans("arrayOfBooleans_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayPOfBooleans("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayPOfBooleans("nonsense")
	})
}

func TestArrayOfBooleansP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfBooleansP("arrayOfBooleans")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 array items, but there were %d", len(actual))
		}

		if *actual[0] != true {
			t.Fatalf("First array item should be true but was %v", *actual[0])
		}

		if *actual[1] != false {
			t.Fatalf("Second array item should be false but was %v", *actual[1])
		}
	})

	t.Run("Array with Null item", func(t *testing.T) {
		actual, err := node.ArrayOfBooleansP("arrayOfBooleans_withNull")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 3 {
			t.Fatalf("Expected 3 array items, but there were %d", len(actual))
		}

		if *actual[0] != true {
			t.Fatalf("First array item should be true but was %v", *actual[0])
		}

		if *actual[1] != false {
			t.Fatalf("Second array item should be false but was %v", *actual[1])
		}

		if actual[2] != nil {
			t.Fatalf("Third array item should be nil but was %v", *actual[2])
		}
	})

	t.Run("Null Array", func(t *testing.T) {
		actual, err := node.ArrayOfBooleansP("array_null")
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
		return node.ArrayOfBooleansP("arrayOfBooleans_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfBooleansP("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfBooleansP("nonsense")
	})
}

func TestArrayOfBooleans(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Array", func(t *testing.T) {
		actual, err := node.ArrayOfBooleans("arrayOfBooleans")
		if err != nil {
			t.Fatalf("Not expecting an error but got: %s", err)
		}

		if len(actual) != 2 {
			t.Fatalf("Expected 2 array items, but there were %d", len(actual))
		}

		if actual[0] != true {
			t.Fatalf("First array item should be true but was %v", actual[0])
		}

		if actual[1] != false {
			t.Fatalf("Second array item should be false but was %v", actual[1])
		}
	})

	t.Run("Invalid Array", func(t *testing.T) {
		_, err := node.ArrayOfBooleans("arrayOfBooleans_invalid")
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
		_, err := node.ArrayOfBooleans("arrayOfBooleans_withNull")
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
		actual, err := node.ArrayOfBooleans("array_null")
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
		return node.ArrayOfBooleans("arrayOfBooleans_invalid")
	})

	checkArrayWrongType(t, func() (interface{}, error) {
		return node.ArrayOfBooleans("string")
	})

	checkArrayNonExistentProperty(t, func() (interface{}, error) {
		return node.ArrayOfBooleans("nonsense")
	})
}
