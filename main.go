package jscraper

import "encoding/json"

// NewFromBytes ...
func NewFromBytes(data []byte) (*Node, error) {
	var value map[string]interface{}
	err := json.Unmarshal(data, &value)
	if err != nil {
		return nil, err
	}

	node := Node{
		ParentNode: nil,
		Path:       "",
		Value:      value,
	}

	return &node, nil
}

// NewFromString ...
func NewFromString(data string) (*Node, error) {
	return NewFromBytes([]byte(data))
}
