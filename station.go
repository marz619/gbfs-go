package gbfs

// Station struct
type Station struct {
	// Required
	StationID string  `json:"station_id"`
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	// Optional
	ShortName     string   `json:"short_name"`
	Address       string   `json:"adress"`
	CrossStreet   string   `json:"cross_street"`
	RegionID      string   `json:"region_id"`
	PostCode      string   `json:"post_code"`
	RentalMethods []string `json:"rental_methods"`
	Capacity      int      `json:"capacity"`
}

// StationInfoData struct
type StationInfoData struct {
	Stations []Station `json:"stations"`
}

// StationInfo struct
type StationInfo struct {
	root
	Data StationInfoData `json:"data"`
}

func (g *gbfs) StationInfO(url string) (StationInfo, error) {
	return StationInfo{}, nil
}
