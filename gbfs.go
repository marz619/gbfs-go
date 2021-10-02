package gbfs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	f "github.com/marz619/gbfs-go/fields"
)

const DefaultRefreshDuration = 10 * time.Second

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
	//
	next(interface{}, int)
}

// RefreshState enum
type RefreshState uint8

const (
	_ RefreshState = iota
	Paused
	Refreshing
	Errored
	Noop
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
		ts:          make(map[interface{}]time.Ticker),
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
	// protected by mutex
	m     sync.Mutex
	state RefreshState                // global state
	ts    map[interface{}]time.Ticker // per object ticker
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
	err = json.NewDecoder(res.Body).Decode(dst)
	if err != nil {
		return err
	}
	if o, ok := dst.(Output); ok {
		o.self = url
	}
	return nil
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
		return Noop
	}
	c.m.Lock()
	defer c.m.Unlock()
	return c.pause()
}

func (c *clientImpl) pause() RefreshState {
	if c.state == Paused {
		return Noop
	}
	// stop all the tickers
	for _, t := range c.ts {
		t.Stop()
	}
	c.state = Paused
	return c.state
}

func (c *clientImpl) Resume() RefreshState {
	if c.state == Refreshing {
		return Noop
	}
	c.m.Lock()
	defer c.m.Unlock()
	return c.resume()
}

func (c *clientImpl) resume() RefreshState {
	if c.state == Refreshing {
		return Noop
	}
	for _, t := range c.ts {
		t.Reset(DefaultRefreshDuration)
	}
	c.state = Refreshing
	return c.state
}

func (c *clientImpl) next(i interface{}, ttl int) {
	// based on TTL set the next timer tick for this object
}

// SystemInformation ...
func (g GBFS) SystemInformation(l f.Language) (s SystemInformation, err error) {
	err = g.c.get(g.Feeds(l).URL("system_information").String(), &s)
	return
}
