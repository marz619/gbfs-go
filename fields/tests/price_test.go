package tests

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"

	"github.com/marz619/gbfs-go/fields"
	"github.com/stretchr/testify/assert"
)

// TestPriceUnmarshalJSON ...
func TestPriceUnmarshalJSON(t *testing.T) {
	// for testing math.MaxFloat64, math.SmallestNonzeroFloat64
	maxF, maxFS := math.MaxFloat64, strconv.FormatFloat(math.MaxFloat64, 'g', -1, 64)
	minNzF, minNzFS := math.SmallestNonzeroFloat64, strconv.FormatFloat(math.SmallestNonzeroFloat64, 'g', -1, 64)

	// standalone
	for _, tc := range []struct {
		name string
		raw  []byte
		expS string
		expF float64
		err  error
	}{
		{"string", []byte(`"3.14159"`), "3.14159", 3.14159, nil},
		{"float", []byte(`3.14159`), "3.14159", 3.14159, nil},
		{"zero_string_no_point", []byte(`"0"`), "0", 0.0, nil},
		{"zero_string", []byte(`"0.0"`), "0.0", 0.0, nil},
		{"zero_float", []byte(`0.0`), "0.0", 0.0, nil},
		{"float_max", []byte(maxFS), maxFS, maxF, nil},
		{"float_smallest_non_zero", []byte(minNzFS), minNzFS, minNzF, nil},
		// {"empty", []byte(`""`), "", 0., nil},
		// errors invalid type
		{"error_type_object", []byte(`{}`), "", 0., fields.ErrInvalidPriceType},
		{"error_type_array", []byte(`[]`), "", 0., fields.ErrInvalidPriceType},
		{"error_type_bool_true", []byte(`true`), "", 0., fields.ErrInvalidPriceType},
		{"error_type_bool_false", []byte(`false`), "", 0., fields.ErrInvalidPriceType},
		// error negative value
		{"non_negative_string", []byte(`"-3.14159"`), "", 0., fields.ErrNonNegativeFloat},
		{"non_negative_float", []byte(`-3.14159`), "", 0., fields.ErrNonNegativeFloat},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// unmarhsal container
			var p fields.Price

			if tc.err != nil {
				assert.ErrorIs(t, json.Unmarshal(tc.raw, &p), tc.err)
			} else {
				assert.NoError(t, json.Unmarshal(tc.raw, &p))
				assert.Equal(t, tc.expS, p.String())
				assert.Equal(t, tc.expF, p.Float64())
			}
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
