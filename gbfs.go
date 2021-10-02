package gbfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	f "github.com/marz619/gbfs-go/fields"
)

// Client interface is the public Client interface used to retrieve data from
// a GBFS API
type Client interface {
	GBFS() (GBFS, error)
}

// AutoRefreshClient extends the Client interface providing the ability to auto
// refresh documents based on the returned TTL
type AutoRefreshClient interface {
	Client
	Pause() RefreshState
	Resume() RefreshState
}

// RefreshState enum
type RefreshState uint8

const (
	_ RefreshState = iota
	Paused
	Refreshing
	Errored
)

// New Client with default http.Client
func New(rootURL string) Client {
	return NewClient(rootURL, nil)
}

// NewClient returns a Client
func NewClient(rootURL string, c *http.Client) Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &clientImpl{
		Client:      c,
		rootURL:     rootURL,
		autoRefresh: false,
	}
}

// NewAutoRefreshClient returns a Client that will self update based on the
// returned TTL
func NewAutoRefreshClient(rootURL string, c *http.Client) AutoRefreshClient {
	if c == nil {
		c = http.DefaultClient
	}
	return &clientImpl{
		Client:      c,
		rootURL:     rootURL,
		autoRefresh: true,
		state:       Paused,
	}
}

// client interface
type client interface {
	get(string, interface{}) error
	set(client)
}

func setC(c client, dst interface{}) {
	if t, ok := dst.(client); ok {
		t.set(c)
	}
}

// internal client implementation
type clientImpl struct {
	*http.Client
	rootURL     string
	autoRefresh bool
	state       RefreshState
}

func (c *clientImpl) set(_ client) {} // noop to satisfy interface

// get retrieves
func (c *clientImpl) get(url string, dst interface{}) error {
	defer setC(c, dst)

	res, err := c.Get(url)
	if err != nil {
		return err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("HTTP<%d>: %s", res.StatusCode, string(content))
	}

	// try to unmarshal as json
	return json.NewDecoder(res.Body).Decode(dst)
}

// ErrNoRootURL error
var ErrNoRootURL = errors.New("no rootURL url")

// GBFS satisfies Client interface
func (c *clientImpl) GBFS() (g GBFS, err error) {
	if c.rootURL == "" {
		err = ErrNoRootURL
		return
	}
	// get the Discover doc
	err = c.get(c.rootURL, &g)
	return
}

func (c *clientImpl) Pause() RefreshState {
	if c.state == Paused {
		return c.state
	}
	err := c.pause()
	if err != nil {
		return Errored.with(err)
	}
}

// SystemInformation ...
func (g GBFS) SystemInformation(l f.Language) (s SystemInformation, err error) {
	err = g.c.get(g.Feeds(l).URL("system_information").String(), &s)
	return
}
