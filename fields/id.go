package fields

import "errors"

// ID type
type ID string

// ErrIDSpaces ...
var ErrIDSpaces = errors.New("ID cannot contain spaces")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (id *ID) UnmarshalJSON(data []byte) error {
	if id == nil {
		id = new(ID)
	}

	raw, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	if containsSpaces(raw) {
		return ErrIDSpaces
	}

	(*id) = ID(raw)
	return nil
}
