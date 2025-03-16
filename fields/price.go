package fields

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Price is represented as float64 or string
type Price struct {
	json.Number
}

// ErrInvalidPriceType returned when the .(type) of unmarshaled data is not a string
// or float
var ErrInvalidPriceType = errors.New("price must be string or float")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (p *Price) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &(p.Number))
	if err != nil {
		return fmt.Errorf("%w %w", err, ErrInvalidPriceType)
	}
	return nil
}

// Float64 returns the float64 value of this Price
func (p Price) Float64() float64 {
	f64, _ := p.Number.Float64()
	return f64
}

// String returns the string value of this Price
func (p Price) String() string {
	return p.Number.String()
}
