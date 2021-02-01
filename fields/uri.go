package fields

import "net/url"

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
