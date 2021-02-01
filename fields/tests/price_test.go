package tests

import (
	"encoding/json"
	"testing"

	"github.com/marz619/gbfs-go/fields"
	"github.com/stretchr/testify/assert"
)

// TestPriceUnmarshalJSON ...
func TestPriceUnmarshalJSON(t *testing.T) {
	// standalone
	for _, tc := range []struct {
		raw  []byte
		expS string
		expF float64
	}{
		{[]byte(`"3.14159"`), "3.14159", 3.14159},
		{[]byte(`3.14159`), "3.14159", 3.14159},
	} {
		t.Run("field: "+string(tc.raw), func(t *testing.T) {
			var p fields.Price
			assert.NoError(t, json.Unmarshal(tc.raw, &p))
			assert.Equal(t, tc.expS, p.String())
			assert.Equal(t, tc.expF, p.Float64())
		})
	}

	// simple json struct
	type S struct {
		P fields.Price `json:"price"`
	}

	for _, tc := range []struct {
		raw  []byte
		expS string
		expF float64
	}{
		{[]byte(`{"price":"2.71828"}`), "2.71828", 2.71828},
		{[]byte(`{"price":2.71828}`), "2.71828", 2.71828},
	} {
		t.Run("struct with value: "+string(tc.raw), func(t *testing.T) {
			var s S
			assert.NoError(t, json.Unmarshal(tc.raw, &s))
			assert.Equal(t, tc.expS, s.P.String())
			assert.Equal(t, tc.expF, s.P.Float64())
		})
	}

	// simple json struct with pointer
	type U struct {
		P *fields.Price `json:"price"`
	}

	for _, tc := range []struct {
		raw   []byte
		expS  string
		expF  float64
		isNil bool
	}{
		{[]byte(`{"price":"1.61803"}`), "1.61803", 1.61803, false},
		{[]byte(`{"price":1.61803}`), "1.61803", 1.61803, false},
		{[]byte(`{}`), "", 0.0, true},
	} {
		t.Run("struct with pointer: "+string(tc.raw), func(t *testing.T) {
			var u U
			assert.NoError(t, json.Unmarshal(tc.raw, &u))
			if !tc.isNil {
				assert.Equal(t, tc.expS, u.P.String())
				assert.Equal(t, tc.expF, u.P.Float64())
			} else {
				assert.Nil(t, u.P)
			}
		})
	}
}
