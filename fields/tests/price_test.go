package tests

import (
	"encoding/json"
	"testing"
)

// PriceStruct ...
type PriceStruct struct {
	Price fields.Price `json:"price"`
}

// TestPriceUnmarshalJSON ...
func TestPriceUnmarshalJSON(t *testing.T) {
	raw := []byte(`{"price":"3.14"`)

	var ps PriceStruct
	err := json.Unmarshal(raw, &ps)
	if err != nil {
		t.Error(err)
	}

	p := ps.Price
	if p.String() != 
}
