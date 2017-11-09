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

// DiscoveryData struct
type DiscoveryData map[string]struct {
	Feeds []Feed `json:"feeds"`
}

// Languages of this DiscoveryData
func (d DiscoveryData) Languages() (langs []string) {
	for k := range d {
		langs = append(langs, k)
	}
	return langs
}

// Discovery struct
type Discovery struct {
	root
	Data DiscoveryData `json:"data"`
}

// Languages of this DiscoveryData
func (d Discovery) Languages() (langs []string) {
	return d.Data.Languages()
}

// Feeds for some language
func (d Discovery) Feeds(lang string) []Feed {
	if f, ok := d.Data[lang]; ok {
		return f.Feeds
	}
	return nil
}

// ErrNoDiscovery error
var ErrNoDiscovery = errors.New("no discovery url")

// Discovery /root JSON object
func (g *gbfs) Discover() (d Discovery, err error) {
	if g.discovery == "" {
		err = ErrNoDiscovery
		return
	}
	// get the Discovery doc
	err = g.get(g.discovery, &d)
	return
}
