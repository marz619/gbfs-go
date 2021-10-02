package fields

import "net/mail"

// Email ...
type Email struct {
	*mail.Address
}

// UnmarshalJSON satisifies json.Unmarshaler interface
func (e *Email) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	// validate email
	(*e).Address, err = mail.ParseAddress(s)
	return err
}
