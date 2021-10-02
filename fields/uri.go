package fields

import "net/url"

// URI ...
type URI struct {
	*url.URL
}

// UnmarshalJSON satisifies json.Unmarshaler interface
func (u *URI) UnmarshalJSON(data []byte) (err error) {
	(*u).URL, err = unmarshalToURL(data)
	return err
}
