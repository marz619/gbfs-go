package fields

import "golang.org/x/text/currency"

// Currency ...
type Currency struct {
	*currency.Unit
}

// UnmarshalJSON implements json.Unmarshaler interface
func (c *Currency) UnmarshalJSON(data []byte) error {
	if c == nil {
		c = new(Currency)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	(*(*c).Unit), err = currency.ParseISO(s)
	return err
}
