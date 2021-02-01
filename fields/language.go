package fields

import "golang.org/x/text/language"

// Language ...
type Language struct {
	*language.Tag
}

// UnmarshalJSON implements json.Unmarshaler interface
func (l *Language) UnmarshalJSON(data []byte) error {
	if l == nil {
		l = new(Language)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	*((*l).Tag), err = language.Parse(s)
	return err
}
