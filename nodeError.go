package jscraper

import "fmt"

const (
	// NodeErrorTypeNotFound ...
	NodeErrorTypeNotFound = "NotFound"

	// NodeErrorTypeWrongType ...
	NodeErrorTypeWrongType = "WrongType"

	// NodeErrorTypeIsNull ...
	NodeErrorTypeIsNull = "IsNull"

	// NodeErrorTypeArrayItemIsNull ...
	NodeErrorTypeArrayItemIsNull = "ArrayItemIsNull"

	// NodeErrorTypeArrayItemWrongType ...
	NodeErrorTypeArrayItemWrongType = "ArrayItemWrongType"
)

// NodeError ...
type NodeError struct {
	Type    string
	Path    string
	message string
}

func (err *NodeError) Error() string {
	return err.message
}

func notFoundError(path string) *NodeError {
	nodeError := NodeError{
		Type:    NodeErrorTypeNotFound,
		Path:    path,
		message: fmt.Sprintf("property '%s' is missing from JSON response", path),
	}

	return &nodeError
}

func wrongTypeError(expectedType, actualType, path string) *NodeError {
	nodeError := NodeError{
		Type:    NodeErrorTypeWrongType,
		Path:    path,
		message: fmt.Sprintf("property '%s' is not of type %s from JSON response, actual type was %s", path, expectedType, actualType),
	}

	return &nodeError
}

func nullTypeError(expectedType, path string) *NodeError {
	nodeError := NodeError{
		Type:    NodeErrorTypeIsNull,
		Path:    path,
		message: fmt.Sprintf("property '%s' is null from JSON response when expecting a value of type %s", path, expectedType),
	}

	return &nodeError
}

func arrayItemNullTypeError(index int, expectedType, path string) *NodeError {
	nodeError := NodeError{
		Type: NodeErrorTypeArrayItemIsNull,
		Path: path,
		message: fmt.Sprintf("array item index %d at property '%s' is null when expecting a value of type %s",
			index, path, expectedType),
	}

	return &nodeError
}

func arrayItemWrongTypeError(index int, expectedType, actualType, path string) *NodeError {
	nodeError := NodeError{
		Type: NodeErrorTypeArrayItemWrongType,
		Path: path,
		message: fmt.Sprintf("array item index %d at property '%s' is not of type %s, actual type was %s",
			index, path, expectedType, actualType),
	}

	return &nodeError
}
