package gbfs

import (
	"encoding/json"
	"errors"
	"reflect"
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
	// non-standard data
	misc map[string]interface{}
}

func (s *SystemInfoData) loadMisc(m map[string]interface{}) {
	s.misc = make(map[string]interface{})

	t := reflect.TypeOf(*s)
	tags := make(map[string]bool, t.NumField())
	// extract tag names
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if name, ok := f.Tag.Lookup("json"); ok {
			tags[name] = true
		}
	}
	// store in our private field
	for k, v := range m {
		if !tags[k] {
			s.misc[k] = v
		}
	}
}

// SystemInfo struct
type SystemInfo struct {
	root
	Data SystemInfoData `json:"data"`
}

// UnmarshalJSON to extract non-standard fields
func (s *SystemInfo) UnmarshalJSON(data []byte) error {
	// extract root
	err := json.Unmarshal(data, &s.root)
	if err != nil {
		return err
	}
	// extract SysInfoData
	err = json.Unmarshal(data, &s.Data)
	if err != nil {
		return err
	}
	// extract misc fields
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	s.Data.loadMisc(m["data"].(map[string]interface{}))
	// success
	return nil
}

// MiscFields not part of the standard
func (s SystemInfo) MiscFields() (m []string) {
	for k := range s.Data.misc {
		m = append(m, k)
	}
	return
}

// Misc field value
func (s SystemInfo) Misc(field string) interface{} {
	if v, ok := s.Data.misc[field]; ok {
		return v
	}
	return nil
}

// ErrNoSysInfoURL error
var ErrNoSysInfoURL = errors.New("must provide system information url")

// SystemInfo data
func (g *gbfs) SystemInfo(url string) (s SystemInfo, err error) {
	if url == "" {
		err = ErrNoSysInfoURL
		return
	}
	// retrieve the system information
	err = g.get(url, &s)
	return
}
