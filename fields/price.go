package fields

import (
	"errors"
	"reflect"
)

// Price is represented as NonNegativeFloat OR string
type Price struct {
	k reflect.Kind // would be used if we implement MarshalJSON
	// internal store
	s string
	n NonNegativeFloat
}

// ErrInvalidPriceType returned when the .(type) of unmarshaled data is not a
// string or float
var ErrInvalidPriceType = errors.New("price must be string or float")

// UnmarshalJSON satisifies json.Unmarshaler interface
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
		p.n, err = unmarshalToNonNegativeFloat([]byte(v.(string)))
		if err != nil {
			return err
		}
		p.s = v.(string)
	}

	return nil
}

// Float64 returns the float64 value contained in this NonNegativeFloat
func (p Price) Float64() float64 {
	return float64(p.n)
}

func (p Price) String() string {
	return p.s
}
