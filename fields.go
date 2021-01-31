package gbfs

import (
	"encoding/json"
	"errors"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

// Field Types: https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#field-types

// TODO(ammaar): define custom error type to wrap internal unmarshaling errors
//
// e.g. url.Error
//      json.UnmarshalTypeError

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

// Date ...
type Date struct {
	*time.Time
}

// date format YYYY-MM-DD
const dateFmt = "2006-01-02"

// UnmarshalJSON implements json.Unmarshaler interface
func (d *Date) UnmarshalJSON(data []byte) error {
	if d == nil {
		d = new(Date)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	*((*d).Time), err = time.Parse(dateFmt, s)
	return err
}

// String implements Stringer interface
func (d *Date) String() string {
	return d.Format(dateFmt)
}

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

// TODO: EnumField

// ID type
type ID string

// ErrIDSpaces ...
var ErrIDSpaces = errors.New("ID cannot contain spaces")

// UnmarshalJSON implements json.Unmarshaler interface
func (id *ID) UnmarshalJSON(data []byte) error {
	if id == nil {
		id = new(ID)
	}

	raw, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	if ContainsSpaces(raw) {
		return ErrIDSpaces
	}

	(*id) = ID(raw)
	return nil
}

// Latitude ...
type Latitude float64

// ErrLatitude ...
var ErrLatitude = errors.New("Latitude must in range [-90.0, 90.0]")

// UnmarshalJSON implements json.Unmarshaler interface
func (l *Latitude) UnmarshalJSON(data []byte) error {
	if l == nil {
		l = new(Latitude)
	}

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
	if l == nil {
		l = new(Longitude)
	}

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

// ErrNonNegativeInt error
var ErrNonNegativeInt = errors.New("NonNegativeInt must have value >= 0")

// NonNegativeInt ...
type NonNegativeInt int

// UnmarshalJSON implements json.Unmarshaler interface
func (n *NonNegativeInt) UnmarshalJSON(data []byte) error {
	if n == nil {
		n = new(NonNegativeInt)
	}

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

// ErrNonNegativeFloat error
var ErrNonNegativeFloat = errors.New("NonNegativeFloat must have value >= 0.0")

// NonNegativeFloat ...
type NonNegativeFloat float64

// UnmarshalJSON implements json.Unmarshaler interface
func (n *NonNegativeFloat) UnmarshalJSON(data []byte) error {
	if n == nil {
		n = new(NonNegativeFloat)
	}

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

// String implements the Stringer interface
func (n *NonNegativeFloat) String() string {
	return strconv.FormatFloat(float64(*n), 'g', -1, 64)
}

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

// Day ...
type Day NonNegativeInt

// ErrInvalidDay ...
var ErrInvalidDay = errors.New("invalid day")

// UnmarshalJSON implements json.Unmarshaler interface
func (d *Day) UnmarshalJSON(data []byte) error {
	n, err := unmarshalToNonNegativeInt(data)
	if err != nil {
		return err
	}

	if *n < 1 || *n > 31 {
		return ErrInvalidDay
	}

	*d = Day(*n)
	return nil
}

// Month ...
type Month NonNegativeInt

// ErrInvalidMonth ...
var ErrInvalidMonth = errors.New("invalid month")

// UnmarshalJSON implements json.Unmarshaler interface
func (m *Month) UnmarshalJSON(data []byte) error {
	n, err := unmarshalToNonNegativeInt(data)
	if err != nil {
		return err
	}

	if *n < 1 || *n > 12 {
		return ErrInvalidMonth
	}

	*m = Month(*n)
	return nil
}

// Year ...
type Year NonNegativeInt

// ErrInvalidYear ...
var ErrInvalidYear = errors.New("invalid year")

// UnmarshalJSON implements json.Unmarshaler interface
func (y *Year) UnmarshalJSON(data []byte) error {
	n, err := unmarshalToNonNegativeInt(data)
	if err != nil {
		return err
	}

	if *n < 0 || *n > 9999 {
		return ErrInvalidYear
	}

	*y = Year(*n)
	return nil
}

// PhoneNumber ...
type PhoneNumber string

// Time ...
type Time struct {
	*time.Time
}

// time format HH:mm:ss
const timeFmt = "15:04:05"

// UnmarshalJSON implements json.Unmarshaler interface
func (t *Time) UnmarshalJSON(data []byte) error {
	if t == nil {
		t = new(Time)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	*((*t).Time), err = time.Parse(timeFmt, s)
	if err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface
func (t *Time) String() string {
	return t.Format(timeFmt)
}

// Timestamp ...
type Timestamp struct {
	*time.Time
}

// UnmarshalJSON implements json.Unmarshaler interface
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	if t == nil {
		t = new(Timestamp)
	}

	i, err := unmarshalToInt(data)
	if err != nil {
		return err
	}

	*((*t).Time) = time.Unix(int64(i), 0).UTC()
	return nil
}

// String implements the Stringer interface
func (t *Timestamp) String() string {
	return strconv.FormatInt(t.Unix(), 10)
}

// Timezone ...
type Timezone struct {
	*time.Location
}

// UnmarshalJSON implements json.Unmarshaler interface
func (t *Timezone) UnmarshalJSON(data []byte) error {
	if t == nil {
		t = new(Timezone)
	}

	zone, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	(*t).Location, err = time.LoadLocation(zone)
	return err
}

// URI ...
type URI struct {
	*url.URL
}

// UnmarshalJSON implements json.Unmarshaler interface
func (u *URI) UnmarshalJSON(data []byte) (err error) {
	if u == nil {
		u = new(URI)
	}

	(*u).URL, err = unmarshalToURL(data)
	return err
}

// ErrURLScheme ...
var ErrURLScheme = errors.New("URL Scheme must be 'http' OR 'https'")

// URL ...
type URL struct {
	*url.URL
}

// UnmarshalJSON implements json.Unmarshaler interface
func (u *URL) UnmarshalJSON(data []byte) error {
	if u == nil {
		u = new(URL)
	}

	url, err := unmarshalToURL(data)
	if err != nil {
		return err
	}

	// check the URL scheme
	switch url.Scheme {
	default:
		return ErrURLScheme
	case "http", "https":
		// valid
	}

	(*u).URL = url
	return nil
}

// unmarshalToFloat64 is a convenience method to unmarshal some bytes into a
// float64
func unmarshalToFloat64(data []byte) (f float64, err error) {
	err = json.Unmarshal(data, &f)
	return
}

// unmarshalToInt is a convenience method to unmarshal some bytes into an int
func unmarshalToInt(data []byte) (i int, err error) {
	err = json.Unmarshal(data, &i)
	return
}

// unmarshalToString is a convenience method to unmarshal some bytes into a
// string
//
// this piece of code:
//
//	 var s string
//	 err := json.Unmarshal(data, &s)
//	 if err != nil {
//	 	return err
//	 }
//
// becomes:
//
//	 s, err := unmarshalToString(data)
//	 if err != nil {
//	 	return err
//	 }
func unmarshalToString(data []byte) (s string, err error) {
	err = json.Unmarshal(data, &s)
	return
}

// unmarshalToURL is a convenience method to unmarshal some bytes into *url.URL
func unmarshalToURL(data []byte) (u *url.URL, err error) {
	raw, err := unmarshalToString(data)
	if err != nil {
		return
	}

	u, err = url.Parse(raw)
	return
}

// unmarshalToNonNegativeInt is a convenience method to unmarshal some bytes
// into *NonNegativeInt
func unmarshalToNonNegativeInt(data []byte) (n *NonNegativeInt, err error) {
	err = json.Unmarshal(data, n)
	return
}

// unmarshalToNonNegativeFloat is a convenience method to unmarshal some bytes
// into *NonNegativeInt
func unmarshalToNonNegativeFloat(data []byte) (n *NonNegativeFloat, err error) {
	err = json.Unmarshal(data, n)
	return
}
