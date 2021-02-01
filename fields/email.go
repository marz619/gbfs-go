package fields

import "net/mail"

// Email ...
type Email struct {
	*mail.Address
}

// UnmarshalJSON implements json.Unmarshaler interface
func (e *Email) UnmarshalJSON(data []byte) error {
	if e == nil {
		e = new(Email)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	// validate email
	(*e).Address, err = mail.ParseAddress(s)
	return err
}
