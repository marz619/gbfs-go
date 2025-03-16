package fields

import (
	"errors"
	"time"
)

// AlertType ...
type AlertType string

const (
	atSystemClosure  AlertType = "SYSTEM_CLOSURE"
	atStationClosure AlertType = "STATION_CLOSURE"
	atStationMove    AlertType = "STATION_MOVE"
	atOther          AlertType = "OTHER"
)

// ErrUnknownAlertType ...
var ErrUnknownAlertType = errors.New("unknown alert type")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (a *AlertType) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	at := AlertType(s)
	switch at {
	default:
		return ErrUnknownAlertType
	case atSystemClosure, atStationClosure, atStationMove, atOther:
	}

	*a = at
	return nil
}

func (a AlertType) String() string {
	return string(a)
}

// DayOfWeek ...
type DayOfWeek string

// DayOfWeek constants
const (
	mon DayOfWeek = "mon"
	tue DayOfWeek = "tue"
	wed DayOfWeek = "wed"
	thu DayOfWeek = "thu"
	fri DayOfWeek = "fri"
	sat DayOfWeek = "sat"
	sun DayOfWeek = "sun"
)

var dowWeekday = map[DayOfWeek]time.Weekday{
	mon: time.Monday,
	tue: time.Tuesday,
	wed: time.Wednesday,
	thu: time.Thursday,
	fri: time.Friday,
	sat: time.Saturday,
	sun: time.Sunday,
}

// Weekday returns the time.Weekday value for this DayOfWeek
func (d DayOfWeek) Weekday() time.Weekday {
	return dowWeekday[d]
}

// ErrUnknownDayOfWeek ...
var ErrUnknownDayOfWeek = errors.New("unknown day")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (d *DayOfWeek) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	dow := DayOfWeek(s)
	switch dow {
	default:
		return ErrUnknownDayOfWeek
	case mon, tue, wed, thu, fri, sat, sun:
	}

	*d = dow
	return nil
}

// Mobile tags
type Mobile string

// Mobile constants
const (
	Android Mobile = "android"
	IOS     Mobile = "ios"
)

// ErrUnknownMobile ...
var ErrUnknownMobile = errors.New("unknown mobile")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (m *Mobile) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	mob := Mobile(s)
	switch mob {
	default:
		return ErrUnknownMobile
	case Android, IOS:
	}

	*m = mob
	return nil
}

// RentalMethod ...
type RentalMethod string

// RentalMethod constants
const (
	RMKey           RentalMethod = "KEY"
	RMCreditcard    RentalMethod = "CREDITCARD"
	RMPaypass       RentalMethod = "PAYPASS"
	RMApplepay      RentalMethod = "APPLEPAY"
	RMAndroidpay    RentalMethod = "ANDROIDPAY"
	RMTransitcard   RentalMethod = "TRANSITCARD"
	RMAccountnumber RentalMethod = "ACCOUNTNUMBER"
	RMPhone         RentalMethod = "PHONE"
)

// ErrUnknownRentalMethod ...
var ErrUnknownRentalMethod = errors.New("unknown rental method")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (r *RentalMethod) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	rm := RentalMethod(s)
	switch rm {
	default:
		return ErrUnknownRentalMethod
	case RMKey, RMCreditcard, RMPaypass, RMApplepay, RMAndroidpay, RMTransitcard, RMAccountnumber, RMPhone:
	}

	*r = rm
	return nil
}

// UserType ...
type UserType string

// UserType constants
const (
	member    UserType = "member"
	nonmember UserType = "nonmember"
)

// ErrUnknownUserType ...
var ErrUnknownUserType = errors.New("unknown user type")

// UnmarshalJSON satisifies json.Unmarshaler interface
func (u *UserType) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	ut := UserType(s)
	switch ut {
	default:
		return ErrUnknownUserType
	case member, nonmember:
	}

	*u = ut
	return nil
}
