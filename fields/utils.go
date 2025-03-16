package fields

import (
	"encoding/json"
	"net/url"
	"unicode"
)

// containsRuneFunc uses a provided predicate func to determine if some rune(s)
// are contained in the string
func containsRuneFunc(s string, f func(rune) bool) bool {
	for _, r := range s {
		if f(r) {
			return true
		}
	}
	return false
}

// containsSpaces determines if there are spaces in the string
func containsSpaces(s string) bool {
	return containsRuneFunc(s, unicode.IsSpace)
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
//	var s string
//	err := json.Unmarshal(data, &s)
//	if err != nil {
//		return err
//	}
//
// becomes:
//
//	s, err := unmarshalToString(data)
//	if err != nil {
//		return err
//	}
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
func unmarshalToNonNegativeInt(data []byte) (n NonNegativeInt, err error) {
	err = json.Unmarshal(data, &n)
	return
}
