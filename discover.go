package gbfs

import (
	"errors"
	"fmt"
)

// Feed struct
type Feed struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (f Feed) String() string {
	return fmt.Sprintf("Feed<%s>: %s", f.Name, f.URL)
}

// DiscoverData struct
type DiscoverData map[string]struct {
	Feeds []Feed `json:"feeds"`
}

// Languages of this DiscoverData
func (d DiscoverData) Languages() (langs []string) {
	for k := range d {
		langs = append(langs, k)
	}
	return langs
}

// Discover struct
type Discover struct {
	root
	Data DiscoverData `json:"data"`
}

// Languages of this DiscoverData
func (d Discover) Languages() (langs []string) {
	return d.Data.Languages()
}

// Feeds for some language
func (d Discover) Feeds(lang string) []Feed {
	if f, ok := d.Data[lang]; ok {
		return f.Feeds
	}
	return nil
}

// ErrNoDiscover error
var ErrNoDiscover = errors.New("no discovery url")

// Discover /root JSON object
func (g *gbfs) Discover() (d Discover, err error) {
	if g.discovery == "" {
		err = ErrNoDiscover
		return
	}
	// get the Discover doc
	err = g.get(g.discovery, &d)
	return
}
