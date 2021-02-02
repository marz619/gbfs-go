package fields

import (
	"errors"
	"strconv"
)

// NonNegativeInt ...
type NonNegativeInt int

// ErrNonNegativeInt error
var ErrNonNegativeInt = errors.New("NonNegativeInt must have value >= 0")

// UnmarshalJSON implements json.Unmarshaler interface
func (n *NonNegativeInt) UnmarshalJSON(data []byte) error {
	i, err := unmarshalToInt(data)
	if err != nil {
		return err
	}

	if i < 0 {
		return ErrNonNegativeInt
	}

	*n = NonNegativeInt(i)
	return nil
}

func (n NonNegativeInt) String() string {
	return strconv.Itoa(int(n))
}

// NonNegativeFloat ...
type NonNegativeFloat float64

// ErrNonNegativeFloat error
var ErrNonNegativeFloat = errors.New("NonNegativeFloat must have value >= 0.0")

// UnmarshalJSON implements json.Unmarshaler interface
func (n *NonNegativeFloat) UnmarshalJSON(data []byte) error {
	f, err := unmarshalToFloat64(data)
	if err != nil {
		return err
	}

	if f < 0.0 {
		return ErrNonNegativeFloat
	}

	*n = NonNegativeFloat(f)
	return nil
}

func (n NonNegativeFloat) String() string {
	// special case 0
	if n == 0 {
		return "0.0"
	}
	return strconv.FormatFloat(float64(n), 'g', -1, 64)
}
