package fields

import "errors"

// Latitude ...
type Latitude float64

// ErrLatitude ...
var ErrLatitude = errors.New("Latitude must in range [-90.0, 90.0]")

// UnmarshalJSON implements json.Unmarshaler interface
func (l *Latitude) UnmarshalJSON(data []byte) error {
	f, err := unmarshalToFloat64(data)
	if err != nil {
		return err
	}

	if f < -90.0 || f > 90.0 {
		return ErrLatitude
	}

	*l = Latitude(f)
	return nil
}

// Longitude ...
type Longitude float64

// ErrLongitude ...
var ErrLongitude = errors.New("Longitude must in range [-180.0, 180.0]")

// UnmarshalJSON implements json.Unmarshaler interface
func (l *Longitude) UnmarshalJSON(data []byte) error {
	f, err := unmarshalToFloat64(data)
	if err != nil {
		return err
	}

	if f < -180.0 || f > 180.0 {
		return ErrLongitude
	}

	*l = Longitude(f)
	return nil
}
