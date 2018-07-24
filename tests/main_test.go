package jscraper_test

import (
	"testing"

	"github.com/sironfoot/jscraper"
)

const ObjectJSON = `
	{
		"string": "Hello world",
		"string_null": null,

		"number": 123.456,
		"number_null": null,

		"boolean": true,
		"boolean_null": null,

		"object": {
			"string": "Hello world",
			"string_null": null,

			"number": 123.456,
			"number_null": null,

			"boolean": true,
			"boolean_null": null
		},
		"object_null": null,

		"arrayOfObjects": [
			{
				"string": "Hello world",
				"string_null": null,

				"number": 123.456,
				"number_null": null,

				"boolean": true,
				"boolean_null": null
			},
			{
				"string": "Hello world",
				"string_null": null,

				"number": 123.456,
				"number_null": null,

				"boolean": true,
				"boolean_null": null
			}
		],

		"arrayOfObjects_withNull": [
			{
				"string": "Hello world",
				"string_null": null,

				"number": 123.456,
				"number_null": null,

				"boolean": true,
				"boolean_null": null
			},
			{
				"string": "Hello world",
				"string_null": null,

				"number": 123.456,
				"number_null": null,

				"boolean": true,
				"boolean_null": null
			},
			null
		],

		"arrayOfObjects_invalid": [
			{
				"string": "Hello world"
			},
			"This is a string"
		],

		"arrayOfStrings": [
			"String 1",
			"String 2",
			"String 3"
		],

		"arrayOfStrings_withNull": [
			"String 1",
			"String 2",
			"String 3",
			null
		],

		"arrayOfStrings_invalid": [
			"String 1",
			123
		],

		"arrayOfNumbers": [
			123,
			1.5,
			789.123,
			0,
			-123,
			-1.5
		],

		"arrayOfNumbers_withNull": [
			123,
			1.5,
			789.123,
			0,
			-123,
			-1.5,
			null
		],

		"arrayOfNumbers_invalid": [
			123,
			"This is a string"
		],

		"arrayOfBooleans": [
			true,
			false
		],

		"arrayOfBooleans_withNull": [
			true,
			false,
			null
		],

		"arrayOfBooleans_invalid": [
			true,
			"This is a string"
		],

		"arrayOfMixed": [
			123,
			"Hello world",
			true,
			{
				"string": "Hello world"
			},
			null
		],

		"array_null": null
	}
`

func TestHasProperty(t *testing.T) {
	data := []byte(ObjectJSON)
	node, err := jscraper.NewFromBytes(data)
	if err != nil {
		t.Fatal(err)
	}

	actual := node.HasProperty("string")
	expected := true
	if actual != expected {
		t.Errorf("Expected %v but got %v when testing property 'string'", expected, actual)
	}

	actual = node.HasProperty("string_null")
	expected = true
	if actual != expected {
		t.Errorf("Expected %v but got %v when testing property 'string_null'", expected, actual)
	}

	actual = node.HasProperty("nonsense")
	expected = false
	if actual != expected {
		t.Errorf("Expected %v but got %v when testing property 'nonsense'", expected, actual)
	}
}
