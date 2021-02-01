package gbfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	f "github.com/marz619/gbfs-go/fields"
)

// Client interface
type Client interface {
	Discover() (GBFS, error)
	GBFS() (GBFS, error)
	SystemInformation(string) (SystemInformation, error)
}

// New Client with default http.Client
func New(discovery string) Client {
	return NewWithClient(discovery, nil)
}

// NewWithClient Client with custom http.Client
func NewWithClient(discovery string, c *http.Client) Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &gbfs{
		c:         c,
		discovery: discovery,
	}
}

// Client
type gbfs struct {
	c         *http.Client
	discovery string

	// cached feeds
	feeds map[f.Language][]Feed
}

func (g *gbfs) get(url string, dst interface{}) error {
	res, err := g.c.Get(url)
	if err != nil {
		return err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("HTTP<%d>: %s", res.StatusCode, content)
	}

	// try to unmarshal as json
	return json.NewDecoder(res.Body).Decode(dst)
}

// ErrNoDiscoveryURL error
var ErrNoDiscoveryURL = errors.New("no discovery url")

// Discover /root JSON object
func (g *gbfs) Discover() (d GBFS, err error) {
	if g.discovery == "" {
		err = ErrNoDiscoveryURL
		return
	}
	// get the Discover doc
	err = g.get(g.discovery, &d)
	return
}

func (g *gbfs) GBFS() (GBFS, error) {
	return g.Discover()
}

// ErrNoSystemInfoURL error
var ErrNoSystemInfoURL = errors.New("must provide system information url")

// SystemInfo data
func (g *gbfs) SystemInformation(url string) (s SystemInformation, err error) {
	if url == "" {
		err = ErrNoSystemInfoURL
		return
	}
	// retrieve the system information
	err = g.get(url, &s)
	return
}
