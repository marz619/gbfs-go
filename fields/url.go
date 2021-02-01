package fields

import (
	"errors"
	"net/url"
)

// URL ...
type URL struct {
	*url.URL
}

// ErrURLScheme ...
var ErrURLScheme = errors.New("URL Scheme must be 'http' OR 'https'")

// UnmarshalJSON implements json.Unmarshaler interface
func (u *URL) UnmarshalJSON(data []byte) error {
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
