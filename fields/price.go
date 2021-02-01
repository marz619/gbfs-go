package fields

import (
	"errors"
	"reflect"
)

// Price ...
type Price struct {
	k reflect.Kind // would be used if we implement
	// internal store
	s string
	n NonNegativeFloat
}

// ErrInvalidPriceType ...
var ErrInvalidPriceType = errors.New("price must be string or float")

// UnmarshalJSON implements json.Unmarshaler interface
func (p *Price) UnmarshalJSON(data []byte) error {
	v, err := unmarshal(data)
	if err != nil {
		return err
	}

	switch v.(type) {
	default:
		return ErrInvalidPriceType
	case float64:
		p.k = reflect.Float64
		p.n, err = unmarshalToNonNegativeFloat(data)
		if err != nil {
			return err
		}
		p.s = p.n.String()
	case string:
		p.k = reflect.String
		p.s = v.(string)
		p.n, err = unmarshalToNonNegativeFloat([]byte(p.s))
		if err != nil {
			return err
		}
	}

	return nil
}

// Float64 ...
func (p Price) Float64() float64 {
	return float64(p.n)
}

func (p Price) String() string {
	return p.s
}
