package gbfs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GBFS interface
type GBFS interface {
	Discover() (Discover, error)
	SystemInfo(url string) (SystemInfo, error)
}

// New GBFS with default http.Client
func New(discovery string) GBFS {
	return &gbfs{
		discovery: discovery,
		client:    http.DefaultClient,
	}
}

// NewClient GBFS with custom http.Client
func NewClient(discovery string, client *http.Client) GBFS {
	if client == nil {
		panic("nil client")
	}
	return &gbfs{
		discovery: discovery,
		client:    client,
	}
}

// root json structure
type root struct {
	LastUpdated int             `json:"last_updated"`
	TTL         int             `json:"ttl"`
	Data        json.RawMessage `json:"data"`
}

// GBFS type
type gbfs struct {
	discovery string
	client    *http.Client

	// cached feeds
	feeds DiscoverData
}

func (g *gbfs) get(url string, dst interface{}) error {
	res, err := g.client.Get(url)
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
