package jscraper_test

import (
	"testing"

	"github.com/sironfoot/jscraper"
)

func TestString(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid String", func(t *testing.T) {
		actual, err := node.String("string")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		expected := "Hello world"

		if actual != expected {
			t.Fatalf("Expected '%s', actual '%s'", expected, actual)
		}
	})

	t.Run("Null String", func(t *testing.T) {
		actual, err := node.String("string_null")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil, returned value was '%s'", actual)
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
		actual, err := node.String("number")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil, returned value was '%s'", actual)
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
		actual, err := node.String("nonsense")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		nodeError, ok := err.(*jscraper.NodeError)
		if !ok {
			t.Fatalf("Expected a error to be of type *NodeError")
		}

		if len(actual) > 0 {
			t.Fatalf("Expected returned value to be a blank string, but it was %s", actual)
		}

		if nodeError.Type != jscraper.NodeErrorTypeNotFound {
			t.Fatalf("Expected error type to be %s but got %s", jscraper.NodeErrorTypeNotFound, nodeError.Type)
		}
	})
}

func TestStringP(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Valid String", func(t *testing.T) {
		actual, err := node.StringP("string")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual == nil {
			t.Fatalf("Returned value shouldn't be nil if no error was returned")
		}

		expected := "Hello world"

		if *actual != expected {
			t.Fatalf("Expected '%s', actual '%s'", expected, *actual)
		}
	})

	t.Run("Null String", func(t *testing.T) {
		actual, err := node.StringP("string_null")
		if err != nil {
			t.Fatalf("Not expecting an error buy got %s", err)
		}
		if actual != nil {
			t.Fatalf("Returned value should be nil but got %s", *actual)
		}
	})

	t.Run("Wrong Type", func(t *testing.T) {
		actual, err := node.StringP("number")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		if actual != nil {
			t.Fatalf("Expected returned value to be nil, but it was %s", *actual)
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
		actual, err := node.StringP("nonsense")
		if err == nil {
			t.Fatalf("Expected an error, but error was nil")
		}
		if actual != nil {
			t.Fatalf("Expected return value to be nil, but got %s", *actual)
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
