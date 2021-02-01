package gbfs

import (
	"errors"

	f "github.com/marz619/gbfs-go/fields"
)

// Output https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#output-format
type Output struct {
	LastUpdated f.Timestamp      `json:"last_updated"`
	TTL         f.NonNegativeInt `json:"ttl"`
	Version     string           `json:"version"`
	// Data is implemented in underlying objects
}

// Feed ...
type Feed struct {
	Name string `json:"name"`
	URL  f.URL  `json:"url"`
}

// GBFS https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#gbfsjson
type GBFS struct {
	Output
	Data map[f.Language]struct {
		Feeds []Feed `json:"feeds"`
	} `json:"data"`
}

// Languages ...
func (g GBFS) Languages() []f.Language {
	ls := make([]f.Language, 0, len(g.Data))
	for l := range g.Data {
		ls = append(ls, l)
	}
	return ls
}

// ErrNoFeed ...
var ErrNoFeed = errors.New("no feed for language")

// Feeds ...
func (g GBFS) Feeds(l f.Language) []Feed {
	return g.Data[l].Feeds
}

// Versions https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#gbfs_versionsjson-added-in-v11
type Versions struct {
	Output
	Data struct {
		Versions []struct {
			Version string `json:"version"`
			URL     f.URL  `json:"url"`
		} `json:"versions"`
	} `json:"data"`
}

// SystemInformation https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_informationjson
type SystemInformation struct {
	Output
	Data struct {
		SystemID         f.ID          `json:"system_id"`
		Language         f.Language    `json:"language"`
		Name             string        `json:"name"`
		ShortName        string        `json:"short_string"`
		Operator         string        `json:"operator"`
		URL              f.URL         `json:"url"`
		PurchaseURL      f.URL         `json:"purchase_url"`
		StartDate        f.Date        `json:"start_date"`
		PhoneNumber      f.PhoneNumber `json:"phone_number"`
		Email            f.Email       `json:"email"`
		FeedContactEmail f.Email       `json:"feed_contact_email"`
		Timezone         f.Timezone    `json:"timezone"`
		LicenseURL       f.URL         `json:"license_url"`
		RentalApps       map[f.Mobile]struct {
			StoreURI     f.URI `json:"store_uri"`
			DiscoveryURI f.URI `json:"discovery_uri"`
		} `json:"rental_apps"`
	} `json:"data"`
}

// StationInformation https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#station_informationjson
type StationInformation struct {
	Output
	Data struct {
		Stations []struct {
			StationID     f.ID             `json:"station_id"`
			Name          string           `json:"name"`
			ShortName     string           `json:"short_name"`
			Latitutde     f.Latitude       `json:"lat"`
			Longitude     f.Longitude      `json:"lon"`
			Address       string           `json:"address"`
			CrossStreet   string           `json:"cross_street"`
			RegionID      f.ID             `json:"region_id"`
			PostCode      string           `json:"post_code"`
			RentalMethods []f.RentalMethod `json:"rental_methods"`
			Capacity      f.NonNegativeInt `json:"capacity"`
			RentalURIs    struct {
				Android f.URI `json:"android"`
				IOS     f.URI `json:"ios"`
				Web     f.URL `json:"web"`
			} `json:"rental_uris"`
		} `json:"stations"`
	} `json:"data"`
}

// StationStatus https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#station_statusjson
type StationStatus struct {
	Output
	Data struct {
		Stations []struct {
			StationID         f.ID             `json:"station_id"`
			NumBikesAvailable f.NonNegativeInt `json:"num_bikes_available"`
			NumBikesDisabled  f.NonNegativeInt `json:"num_bikes_disabled"`
			NumDocksAvailable f.NonNegativeInt `json:"num_docs_available"`
			NumDocsDisabled   f.NonNegativeInt `json:"num_docs_disabled"`
			IsInstalled       bool             `json:"is_installed"`
			IsRenting         bool             `json:"is_renting"`
			IsReturning       bool             `json:"is_returning"`
			LastReported      f.Timestamp      `json:"last_reported"`
		} `json:"stations"`
	} `json:"data"`
}

// FreeBikeStatus https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#free_bike_statusjson
type FreeBikeStatus struct {
	Output
	Data struct {
		Bikes []struct {
			BikeID     f.ID        `json:"bike_id"`
			Latitude   f.Latitude  `json:"latitude"`
			Longitude  f.Longitude `json:"longitude"`
			IsReserved bool        `json:"is_reserved"`
			IsDisabled bool        `json:"is_disabled"`
			RentalURIs struct {
				Android f.URI `json:"android"`
				IOS     f.URI `json:"ios"`
				Web     f.URL `json:"web"`
			} `json:"rental_uris"`
		} `json:"bikes"`
	} `json:"data"`
}

// SystemHours https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_hoursjson
type SystemHours struct {
	Output
	Data struct {
		RentalHours []struct {
			UserTypes []f.UserType  `json:"user_types"`
			Days      []f.DayOfWeek `json:"days"`
			StartTime f.Time        `json:"start_time"`
			EndTime   f.Time        `json:"end_time"`
		} `json:"rental_hours"`
	} `json:"data"`
}

// SystemCalendar https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_calendarjson
type SystemCalendar struct {
	Output
	Data struct {
		Calendars []struct {
			StartDay   f.Day   `json:"start_day"`
			StartMonth f.Month `json:"start_month"`
			StartYear  f.Year  `json:"start_year"`
			EndDay     f.Day   `json:"eend_day"`
			EndMonth   f.Month `json:"end_month"`
			EndYear    f.Year  `json:"end_year"`
		} `json:"calendar"`
	} `json:"data"`
}

// SystemRegions https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_regionsjson
type SystemRegions struct {
	Output
	Data struct {
		Regions []struct {
			RegionID f.ID   `json:"region_id"`
			Name     string `json:"name"`
		} `json:"regions"`
	} `json:"data"`
}

// SystemPricingPlans https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_pricing_plansjson
type SystemPricingPlans struct {
	Output
	Data struct {
		Plans []struct {
			PlanID      f.ID       `json:"plan_id"`
			URL         f.URL      `json:"url"`
			Name        string     `json:"name"`
			Currency    f.Currency `json:"currency"`
			Price       f.Price    `json:"price"`
			IsTaxable   bool       `json:"is_taxable"`
			Description string     `json:"description"`
		} `json:"plans"`
	} `json:"data"`
}

// SystemAlerts https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_alertsjson
type SystemAlerts struct {
	Output
	Data []struct {
		Alerts []struct {
			AlertID f.ID        `json:"alert_id"`
			Type    f.AlertType `json:"type"`
			Times   []struct {
				Start f.Timestamp `json:"start"`
				End   f.Timestamp `json:"end"`
			} `json:"times"`
			StationIds  []f.ID      `json:"station_ids"`
			RegionIds   []f.ID      `json:"region_ids"`
			URL         f.URL       `json:"url"`
			Summary     string      `json:"summary"`
			Description string      `json:"description"`
			LastUpdated f.Timestamp `json:"last_updated"`
		} `json:"alerts"`
	} `json:"data"`
}
