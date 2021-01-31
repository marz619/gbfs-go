package gbfs

import "errors"

// Output Format: https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#output-format
type Output struct {
	LastUpdated Timestamp      `json:"last_updated"`
	TTL         NonNegativeInt `json:"ttl"`
	Version     string         `json:"version"`
	// Data is implemented in underlying objects
}

// GBFS https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#gbfsjson
type GBFS struct {
	Output
	Data map[Language]struct {
		Feeds []struct {
			Name string `json:"name"`
			URL  URL    `json:"url"`
		} `json:"feeds"`
	} `json:"data"`
}

// Versions https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#gbfs_versionsjson-added-in-v11
type Versions struct {
	Output
	Data struct {
		Versions []struct {
			Version string `json:"version"`
			URL     URL    `json:"url"`
		} `json:"versions"`
	} `json:"data"`
}

// mobile tags
type mobile string

// mobile constants
const (
	android mobile = "android"
	iOS            = "ios"
)

// ErrUnknownMobile ...
var ErrUnknownMobile = errors.New("unknown mobile")

func (m *mobile) UnmarshalJSON(data []byte) error {
	if m == nil {
		m = new(mobile)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	mob := mobile(s)
	switch mob {
	default:
		return ErrUnknownMobile
	case android, iOS:
	}

	*m = mob
	return nil
}

// SystemInformation https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_informationjson
type SystemInformation struct {
	Output
	Data struct {
		SystemID         ID          `json:"system_id"`
		Language         Language    `json:"language"`
		Name             string      `json:"name"`
		ShortName        string      `json:"short_string"`
		Operator         string      `json:"operator"`
		URL              URL         `json:"url"`
		PurchaseURL      URL         `json:"purchase_url"`
		StartDate        Date        `json:"start_date"`
		PhoneNumber      PhoneNumber `json:"phone_number"`
		Email            Email       `json:"email"`
		FeedContactEmail Email       `json:"feed_contact_email"`
		Timezone         Timezone    `json:"timezone"`
		LicenseURL       URL         `json:"license_url"`
		RentalApps       map[mobile]struct {
			StoreURI     URI `json:"store_uri"`
			DiscoveryURI URI `json:"discovery_uri"`
		} `json:"rental_apps"`
	} `json:"data"`
}

// rentalMethod ...
type rentalMethod string

// rentalMethod constants
const (
	rmKey           rentalMethod = "KEY"
	rmCreditcard                 = "CREDITCARD"
	rmPaypass                    = "PAYPASS"
	rmApplepay                   = "APPLEPAY"
	rmAndroidpay                 = "ANDROIDPAY"
	rmTransitcard                = "TRANSITCARD"
	rmAccountnumber              = "ACCOUNTNUMBER"
	rmPhone                      = "PHONE"
)

// ErrUnknownRentalMethod ...
var ErrUnknownRentalMethod = errors.New("unknown rental method")

func (r *rentalMethod) UnmarshalJSON(data []byte) error {
	if r == nil {
		r = new(rentalMethod)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	rm := rentalMethod(s)
	switch rm {
	default:
		return ErrUnknownRentalMethod
	case rmKey, rmCreditcard, rmPaypass, rmApplepay, rmAndroidpay, rmTransitcard, rmAccountnumber, rmPhone:
	}

	*r = rm
	return nil
}

// StationInformation https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#station_informationjson
type StationInformation struct {
	Output
	Data struct {
		Stations []struct {
			StationID     ID             `json:"station_id"`
			Name          string         `json:"name"`
			ShortName     string         `json:"short_name"`
			Latitutde     Latitude       `json:"lat"`
			Longitude     Longitude      `json:"lon"`
			Address       string         `json:"address"`
			CrossStreet   string         `json:"cross_street"`
			RegionID      ID             `json:"region_id"`
			PostCode      string         `json:"post_code"`
			RentalMethods []rentalMethod `json:"rental_methods"`
			Capacity      NonNegativeInt `json:"capacity"`
			RentalURIs    struct {
				Android URI `json:"android"`
				IOS     URI `json:"ios"`
				Web     URL `json:"web"`
			} `json:"rental_uris"`
		} `json:"stations"`
	} `json:"data"`
}

// StationStatus https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#station_statusjson
type StationStatus struct {
	Output
	Data struct {
		Stations []struct {
			StationID         ID             `json:"station_id"`
			NumBikesAvailable NonNegativeInt `json:"num_bikes_available"`
			NumBikesDisabled  NonNegativeInt `json:"num_bikes_disabled"`
			NumDocksAvailable NonNegativeInt `json:"num_docs_available"`
			NumDocsDisabled   NonNegativeInt `json:"num_docs_disabled"`
			IsInstalled       bool           `json:"is_installed"`
			IsRenting         bool           `json:"is_renting"`
			IsReturning       bool           `json:"is_returning"`
			LastReported      Timestamp      `json:"last_reported"`
		} `json:"stations"`
	} `json:"data"`
}

// FreeBikeStatus https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#free_bike_statusjson
type FreeBikeStatus struct {
	Output
	Data struct {
		Bikes []struct {
			BikeID     ID        `json:"bike_id"`
			Latitude   Latitude  `json:"latitude"`
			Longitude  Longitude `json:"longitude"`
			IsReserved bool      `json:"is_reserved"`
			IsDisabled bool      `json:"is_disabled"`
			RentalURIs struct {
				Android URI `json:"android"`
				IOS     URI `json:"ios"`
				Web     URL `json:"web"`
			} `json:"rental_uris"`
		} `json:"bikes"`
	} `json:"data"`
}

// userType ...
type userType string

// userType constants
const (
	member    userType = "member"
	nonmember userType = "nonmember"
)

// ErrUnknownUserType ...
var ErrUnknownUserType = errors.New("unknown user type")

func (u *userType) UnmarshalJSON(data []byte) error {
	if u == nil {
		u = new(userType)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	ut := userType(s)
	switch ut {
	default:
		return ErrUnknownUserType
	case member, nonmember:
	}

	*u = ut
	return nil
}

// dayOfWeek ...
type dayOfWeek string

// dayOfWeek constants
const (
	mon dayOfWeek = "mon"
	tue           = "tue"
	web           = "web"
	thu           = "thu"
	fri           = "fri"
	sat           = "sat"
	sun           = "sun"
)

// ErrUnknownDay ...
var ErrUnknownDay = errors.New("unknown day of Week")

func (d *dayOfWeek) UnmarshalJSON(data []byte) error {
	if d == nil {
		d = new(dayOfWeek)
	}

	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	_d := dayOfWeek(s)
	switch _d {
	default:
		return ErrUnknownUserType
	case mon, tue, web, thu, fri, sat, sun:
	}

	*d = _d
	return nil
}

// SystemHours https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_hoursjson
type SystemHours struct {
	Output
	Data struct {
		RentalHours []struct {
			UserTypes []userType  `json:"user_types"`
			Days      []dayOfWeek `json:"days"`
			StartTime Time        `json:"start_time"`
			EndTime   Time        `json:"end_time"`
		} `json:"rental_hours"`
	} `json:"data"`
}

// SystemCalendar https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_calendarjson
type SystemCalendar struct {
	Output
	Data struct {
		Calendars []struct {
			StartDay   Day   `json:"start_day"`
			StartMonth Month `json:"start_month"`
			StartYear  Year  `json:"start_year"`
			EndDay     Day   `json:"eend_day"`
			EndMonth   Month `json:"end_month"`
			EndYear    Year  `json:"end_year"`
		} `json:"calendar"`
	} `json:"data"`
}

// SystemRegions https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_regionsjson
type SystemRegions struct {
	Output
	Data struct {
		Regions []struct {
			RegionID ID     `json:"region_id"`
			Name     string `json:"name"`
		} `json:"regions"`
	} `json:"data"`
}

// SystemPricingPlans https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_pricing_plansjson
type SystemPricingPlans struct {
	Output
	Data struct {
		Plans []struct {
			PlanID      ID       `json:"plan_id"`
			URL         URL      `json:"url"`
			Name        string   `json:"name"`
			Currency    Currency `json:"currency"`
			Price       Price    `json:"price"`
			IsTaxable   bool     `json:"is_taxable"`
			Description string   `json:"description"`
		} `json:"plans"`
	} `json:"data"`
}

type alertType string

const (
	atSystemClosure  alertType = "SYSTEM_CLOSURE"
	atStationClosure           = "STATION_CLOSURE"
	atStationMove              = "STATION_MOVE"
	atOther                    = "OTHER"
)

// ErrUnknownAlertType ...
var ErrUnknownAlertType = errors.New("unknown alert type")

func (a *alertType) UnmarshalJSON(data []byte) error {
	s, err := unmarshalToString(data)
	if err != nil {
		return err
	}

	at := alertType(s)
	switch at {
	default:
		return ErrUnknownAlertType
	case atSystemClosure, atStationClosure, atStationMove, atOther:
	}

	*a = at
	return nil
}

func (a alertType) String() string {
	return string(a)
}

// SystemAlerts https://github.com/NABSA/gbfs/blob/v2.0/gbfs.md#system_alertsjson
type SystemAlerts struct {
	Output
	Data []struct {
		Alerts []struct {
			AlertID ID        `json:"alert_id"`
			Type    alertType `json:"type"`
			Times   []struct {
				Start Timestamp `json:"start"`
				End   Timestamp `json:"end"`
			} `json:"times"`
			StationIds  []ID      `json:"station_ids"`
			RegionIds   []ID      `json:"region_ids"`
			URL         URL       `json:"url"`
			Summary     string    `json:"summary"`
			Description string    `json:"description"`
			LastUpdated Timestamp `json:"last_updated"`
		} `json:"alerts"`
	} `json:"data"`
}
