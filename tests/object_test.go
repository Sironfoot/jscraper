package jscraper_test

import (
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestIsObject(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Is an Object", func(t *testing.T) {
		actual := node.IsObject("object")
		expected := true

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Is Not an Object", func(t *testing.T) {
		actual := node.IsObject("string")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Null Object", func(t *testing.T) {
		actual := node.IsObject("object_null")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Non-existent Property", func(t *testing.T) {
		actual := node.IsObject("nonsense")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})
}

func TestObject(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Object", func(t *testing.T) {
		actual, err := node.Object("object")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}

		if len(actual.Value) == 0 {
			t.Fatalf("Expect object properties, but there were none")
		}

		actualString, err := actual.String("string")
		if err != nil {
			t.Fatalf("Expecting a property called 'string' on the object, but got error %s", err)
		}

		expectedString := "Hello world"

		if actualString != expectedString {
			t.Fatalf("Expecting object's 'string' property to be %s but was %s", expectedString, actualString)
		}
	})

	t.Run("Null Object", func(t *testing.T) {
		actual, err := node.Object("object_null")
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

	t.Run("Wrong Type", func(t *testing.T) {
		actual, err := node.Object("string")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil, returned value was '%v'", actual)
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeWrongType {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeWrongType, nodeError.Type)
		}
	})

	t.Run("Non-existent Property", func(t *testing.T) {
		actual, err := node.Object("nonsense")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if actual.ParentNode != nil {
			t.Fatalf("Expected returned value to be false, but it was %v", actual)
		}

		if nodeError.Type != jscraper.NodeErrorTypeNotFound {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeNotFound, nodeError.Type)
		}
	})
}

func TestObjectP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Object", func(t *testing.T) {
		actual, err := node.ObjectP("object")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		if len(actual.Value) == 0 {
			t.Fatalf("Expect object properties, but there were none")
		}

		actualString, err := actual.String("string")
		if err != nil {
			t.Fatalf("Expecting a property called 'string' on the object, but got error %s", err)
		}

		expectedString := "Hello world"

		if actualString != expectedString {
			t.Fatalf("Expecting object's 'string' property to be %s but was %s", expectedString, actualString)
		}
	})

	t.Run("Null Object", func(t *testing.T) {
		actual, err := node.ObjectP("object_null")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", *actual)
		}
	})

	t.Run("Wrong Type", func(t *testing.T) {
		actual, err := node.ObjectP("string")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		if actual != nil {
			t.Fatalf("Expected returned value to be nil, but it was %v", *actual)
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeWrongType {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeWrongType, nodeError.Type)
		}
	})

	t.Run("Non-existent Property", func(t *testing.T) {
		actual, err := node.ObjectP("nonsense")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		if actual != nil {
			t.Fatalf("Expected return value to be nil, but got %v", *actual)
		}

		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if nodeError.Type != jscraper.NodeErrorTypeNotFound {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeNotFound, nodeError.Type)
		}
	})
}
