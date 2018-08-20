package jscraper_test

import (
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestIsNumber(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Is a Number", func(t *testing.T) {
		actual := node.IsNumber("number")
		expected := true

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Is Not a Number", func(t *testing.T) {
		actual := node.IsNumber("string")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Null Number", func(t *testing.T) {
		actual := node.IsNumber("number_null")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})

	t.Run("Non-existent Property", func(t *testing.T) {
		actual := node.IsNumber("nonsense")
		expected := false

		if actual != expected {
			t.Fatalf("Expected '%v', actual '%v'", expected, actual)
		}
	})
}

func TestNumber(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Number", func(t *testing.T) {
		actual, err := node.Number("number")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		expected := 123.456

		if actual != expected {
			t.Fatalf("Expected '%f', actual '%f'", expected, actual)
		}
	})

	t.Run("Null Number", func(t *testing.T) {
		actual, err := node.Number("number_null")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil, returned value was '%f'", actual)
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
		actual, err := node.Number("string")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil, returned value was '%f'", actual)
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
		actual, err := node.Number("nonsense")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if actual != 0 {
			t.Fatalf("Expected returned value to be a blank string, but it was %f", actual)
		}

		if nodeError.Type != jscraper.NodeErrorTypeNotFound {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeNotFound, nodeError.Type)
		}
	})
}

func TestNumberP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid Number", func(t *testing.T) {
		actual, err := node.NumberP("number")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		expected := 123.456

		if *actual != expected {
			t.Fatalf("Expected '%f', actual '%f'", expected, *actual)
		}
	})

	t.Run("Null Number", func(t *testing.T) {
		actual, err := node.NumberP("number_null")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %f", *actual)
		}
	})

	t.Run("Wrong Type", func(t *testing.T) {
		actual, err := node.NumberP("string")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		if actual != nil {
			t.Fatalf("Expected returned value to be nil, but it was %f", *actual)
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
		actual, err := node.NumberP("nonsense")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		if actual != nil {
			t.Fatalf("Expected return value to be nil, but got %f", *actual)
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
