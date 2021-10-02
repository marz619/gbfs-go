package fields

import (
	"errors"
	"strconv"
	"time"
)

// Date ...
type Date struct {
	time.Time
}

// date format YYYY-MM-DD
const dateFmt = "2006-01-02"

// UnmarshalJSON satisifies json.Unmarshaler interface
func (d *Date) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	(*d).Time, err = time.Parse(dateFmt, s)
	return err
}

func (d *Date) String() string {
	return d.Format(dateFmt)
}

// Day ...
type Day NonNegativeInt

// ErrInvalidDay ...
var ErrInvalidDay = errors.New("invalid day")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (d *Day) UnmarshalJSON(data []byte) error {
	n, err := unmarshalToNonNegativeInt(data)
	if err != nil {
		return err
	}

	if n < 1 || n > 31 {
		return ErrInvalidDay
	}

	*d = Day(n)
	return nil
}

// Month ...
type Month NonNegativeInt

// ErrInvalidMonth ...
var ErrInvalidMonth = errors.New("invalid month")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (m *Month) UnmarshalJSON(data []byte) error {
	n, err := unmarshalToNonNegativeInt(data)
	if err != nil {
		return err
	}

	if n < 1 || n > 12 {
		return ErrInvalidMonth
	}

	*m = Month(n)
	return nil
}

// Year ...
type Year NonNegativeInt

// ErrInvalidYear ...
var ErrInvalidYear = errors.New("invalid year")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (y *Year) UnmarshalJSON(data []byte) error {
	n, err := unmarshalToNonNegativeInt(data)
	if err != nil {
		return err
	}

	if n < 0 || n > 9999 {
		return ErrInvalidYear
	}

	*y = Year(n)
	return nil
}

// Time ...
type Time struct {
	time.Time
}

// time format HH:mm:ss
const timeFmt = "15:04:05"

// UnmarshalJSON satisifies json.Unmarshaler interface
func (t *Time) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	(*t).Time, err = time.Parse(timeFmt, s)
	if err != nil {
		return err
	}

	return nil
}

func (t *Time) String() string {
	return t.Format(timeFmt)
}

// Timestamp ...
type Timestamp struct {
	time.Time
}

// UnmarshalJSON satisifies json.Unmarshaler interface
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	i, err := unmarshalToInt(data)
	if err != nil {
		return err
	}

	(*t).Time = time.Unix(int64(i), 0).UTC()
	return nil
}

func (t *Timestamp) String() string {
	return strconv.FormatInt(t.Unix(), 10)
}

// Timezone ...
type Timezone struct {
	*time.Location
}

// UnmarshalJSON satisifies json.Unmarshaler interface
func (t *Timezone) UnmarshalJSON(data []byte) error {
	zone, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	(*t).Location, err = time.LoadLocation(zone)
	return err
}
