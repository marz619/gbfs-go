package fields

import (
	"reflect"
	"strconv"
)

// Price ...
type Price struct {
	k reflect.Kind // would be used if we implement
	// internal store
	s string
	n NonNegativeFloat
}

// UnmarshalJSON implements json.Unmarshaler interface
func (p *Price) UnmarshalJSON(data []byte) error {
	if p == nil {
		p = new(Price)
	}

	switch data[0] {
	default:
		p.k = reflect.Float64

		n, err := unmarshalToNonNegativeFloat(data)
		if err != nil {
			return err
		}

		p.n = *n
		p.s = n.String()
	case '"':
		p.k = reflect.String

		s, err := unmarshalToString(data)
		if err != nil {
			return err
		}

		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}

		p.s = s
		p.n = NonNegativeFloat(f)
	}

	return nil
}

// Float64 ...
func (p *Price) Float64() float64 {
	return float64(p.n)
}

func (p *Price) String() string {
	return p.s
}
