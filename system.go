package gbfs

import (
	"errors"
)

// SystemInfoData struct
type SystemInfoData struct {
	// Required
	SystemID string `json:"system_id"`
	Language string `json:"language"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
	// Optional
	ShortName   string `json:"short_name"`
	Operator    string `json:"operator"`
	URL         string `json:"url"`
	PurchaseURL string `json:"purchase_urL"`
	StartDate   string `json:"start_date"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	LicenseURL  string `json:"license_url"`
}

// SystemInfo struct
type SystemInfo struct {
	root
	Data SystemInfoData `json:"data"`
}

// ErrNoSystemInfoURL error
var ErrNoSystemInfoURL = errors.New("must provide system information url")

// SystemInfo data
func (g *gbfs) SystemInfo(url string) (s SystemInfo, err error) {
	if url == "" {
		err = ErrNoSystemInfoURL
		return
	}
	// retrieve the system information
	err = g.get(url, &s)
	return
}
