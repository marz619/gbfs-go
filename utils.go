package gbfs

import (
	"bytes"
	"encoding/json"
	"errors"
)

// ID type
type ID struct {
	id  string
	raw []byte
}

// ID errors
var (
	ErrInvalidID    = errors.New("invalid id value")
	ErrNilOrEmptyID = errors.New("nil or empty raw id")
)

// double quote byte
const dquote = byte('"')

// const bad id starts
const invalidIDStart = "{[:"

// const null bytes
var null = []byte{'n', 'u', 'l', 'l'}

// UnmarshalJSON implements json.Unmarshaler interface
func (id *ID) UnmarshalJSON(data []byte) error {
	// check if data is somewhat valid
	if len(data) == 0 ||
		bytes.Equal(data, null) ||
		bytes.ContainsAny([]byte{data[0]}, invalidIDStart) {
		return ErrInvalidID
	}

	// create a copy of the data
	id.raw = make([]byte, len(data))
	copy(id.raw, data)

	// wrap in quotes to coerce to string
	if data[0] != dquote {
		data = append(append([]byte{dquote}, data...), dquote)
	}

	// unmarshal as string
	return json.Unmarshal(data, &id.id)
}

// MarshalJSON implements the json.Marshaler interface
func (id ID) MarshalJSON() ([]byte, error) {
	if len(id.raw) == 0 {
		return nil, ErrNilOrEmptyID
	}
	return id.raw, nil
}
