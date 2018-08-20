package jscraper_test

import (
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestIsBoolean(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Is a Boolean", func(t *testing.T) {
		actual := node.IsBoolean("boolean")
		expected := true

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Is Not a Boolean", func(t *testing.T) {
		actual := node.IsBoolean("string")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Null Boolean", func(t *testing.T) {
		actual := node.IsBoolean("boolean_null")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Non-existent Property", func(t *testing.T) {
		actual := node.IsBoolean("nonsense")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})
}

func TestBoolean(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Boolean", func(t *testing.T) {
		actual, err := node.Boolean("boolean")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		expected := true

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Null Boolean", func(t *testing.T) {
		actual, err := node.Boolean("boolean_null")
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
		actual, err := node.Boolean("string")
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
		actual, err := node.Boolean("nonsense")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if actual != false {
			t.Fatalf("Expected returned value to be false, but it was %v", actual)
		}

		if nodeError.Type != jscraper.NodeErrorTypeNotFound {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeNotFound, nodeError.Type)
		}
	})
}

func TestBooleanP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Boolean", func(t *testing.T) {
		actual, err := node.BooleanP("boolean")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		expected := true

		if *actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, *actual)
		}
	})

	t.Run("Null Boolean", func(t *testing.T) {
		actual, err := node.BooleanP("boolean_null")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %v", *actual)
		}
	})

	t.Run("Wrong Type", func(t *testing.T) {
		actual, err := node.BooleanP("string")
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
		actual, err := node.BooleanP("nonsense")
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
