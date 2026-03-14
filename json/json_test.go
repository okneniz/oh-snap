package json

import (
	"bytes"
	stdjson "encoding/json"
	"math"
	"strconv"
	"testing"
	"time"

	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
)

// TestEncodingJSON_RoundTrip generates JSON-like values and ensures that
// encoding/json.Marshal + Unmarshal round-trips them without losing data
// (numbers decoded as float64), using ohsnap.Check property-based style.
func TestEncodingJSON_RoundTrip(t *testing.T) {
	t.Parallel()
	const iterations = 10_000
	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)
	rnd := rand.New(rand.NewPCG(0, uint64(seed)))
	arb := ArbitraryJSON(rnd, 4, 6)

	ohsnap.Check(t, iterations, arb, func(v Value) bool {
		b, err := stdjson.Marshal(v)
		if err != nil {
			t.Logf("marshal failed for value (seed %d): %#v", seed, v)
			return false
		}

		t.Logf("json: %s", string(b))

		var dst interface{}
		if err := stdjson.Unmarshal(b, &dst); err != nil {
			t.Logf("unmarshal failed for bytes %q (seed %d)", string(b), seed)
			return false
		}
		if !valuesEqual(v, dst) {
			t.Logf("round-trip mismatch (seed %d): orig=%#v decoded=%#v", seed, v, dst)
			return false
		}
		return true
	})
}

// TestEncodingJSON_UseNumber_RoundTrip same as above but decodes using
// Decoder.UseNumber() so numeric values may come back as json.Number, using ohsnap.Check property-based style.
func TestEncodingJSON_UseNumber_RoundTrip(t *testing.T) {
	t.Parallel()
	const iterations = 1000
	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)
	rnd := rand.New(rand.NewPCG(0, uint64(seed)))
	arb := ArbitraryJSON(rnd, 4, 6)

	ohsnap.Check(t, iterations, arb, func(v Value) bool {
		b, err := stdjson.Marshal(v)
		if err != nil {
			t.Logf("marshal failed for value (seed %d): %#v", seed, v)
			return false
		}
		dec := stdjson.NewDecoder(bytes.NewReader(b))
		dec.UseNumber()
		var dst interface{}
		if err := dec.Decode(&dst); err != nil {
			t.Logf("decode with UseNumber failed for bytes %q (seed %d)", string(b), seed)
			return false
		}
		if !valuesEqualWithNumber(v, dst) {
			t.Logf("round-trip mismatch (UseNumber, seed %d): orig=%#v decoded=%#v", seed, v, dst)
			return false
		}
		return true
	})
}

// TestEncodingJSON_MarshalIndent_RoundTrip ensures that encoding/json.MarshalIndent
// (pretty-printing) round-trips JSON values correctly.
func TestEncodingJSON_MarshalIndent_RoundTrip(t *testing.T) {
	t.Parallel()
	const iterations = 1000
	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)
	rnd := rand.New(rand.NewPCG(0, uint64(seed)))
	arb := ArbitraryJSON(rnd, 4, 6)

	ohsnap.Check(t, iterations, arb, func(v Value) bool {
		b, err := stdjson.MarshalIndent(v, "", "  ")
		if err != nil {
			t.Logf("MarshalIndent failed for value (seed %d): %#v", seed, v)
			return false
		}
		var dst interface{}
		if err := stdjson.Unmarshal(b, &dst); err != nil {
			t.Logf("unmarshal failed for pretty JSON bytes %q (seed %d)", string(b), seed)
			return false
		}
		if !valuesEqual(v, dst) {
			t.Logf("pretty round-trip mismatch (seed %d): orig=%#v decoded=%#v", seed, v, dst)
			return false
		}
		return true
	})
}

// TestEncodingJSON_MarshalIndent_VariousIndents ensures that encoding/json.MarshalIndent
// (pretty-printing) round-trips JSON values correctly for various indent options.
func TestEncodingJSON_MarshalIndent_VariousIndents(t *testing.T) {
	t.Parallel()
	const iterations = 1000
	seed := time.Now().UnixNano()
	t.Logf("seed: %v", seed)
	rnd := rand.New(rand.NewPCG(0, uint64(seed)))
	arb := ArbitraryJSON(rnd, 4, 6)

	indents := []string{"", " ", "  ", "\t"}
	for _, indent := range indents {
		indent := indent // capture range variable
		t.Run("indent="+strconv.Quote(indent), func(t *testing.T) {
			ohsnap.Check(t, iterations, arb, func(v Value) bool {
				b, err := stdjson.MarshalIndent(v, "", indent)
				if err != nil {
					t.Logf("MarshalIndent failed for indent %q (seed %d): %#v", indent, seed, v)
					return false
				}
				var dst interface{}
				if err := stdjson.Unmarshal(b, &dst); err != nil {
					t.Logf("unmarshal failed for pretty JSON bytes %q (indent %q, seed %d)", string(b), indent, seed)
					return false
				}
				if !valuesEqual(v, dst) {
					t.Logf("pretty round-trip mismatch (indent %q, seed %d): orig=%#v decoded=%#v", indent, seed, v, dst)
					return false
				}
				return true
			})
		})
	}
}

// valuesEqual asserts orig Value equals decoded value where numbers are float64, returns bool for property-based testing.
//
// Why not just use reflect.DeepEqual?
// - NaN: reflect.DeepEqual(math.NaN(), math.NaN()) is false, but for round-trip JSON, NaN should equal NaN.
// - Float tolerance: reflect.DeepEqual requires exact float equality, but JSON encoding/decoding may introduce tiny rounding errors.
// - Type flexibility: JSON numbers may decode as float64 or json.Number, which DeepEqual considers different types.
// - Decoded structure: JSON decoding may produce interface{} values that are not exactly the same Go type as the original.
//
// This function handles all those cases for robust property-based JSON tests.
func valuesEqual(orig Value, decoded interface{}) bool {
	switch o := orig.(type) {
	case nil:
		return decoded == nil
	case bool:
		db, ok := decoded.(bool)
		return ok && o == db
	case int64:
		switch d := decoded.(type) {
		case float64:
			return math.Abs(float64(o)-d) < 1e-9
		case int64:
			return o == d
		default:
			return false
		}
	case float64:
		switch d := decoded.(type) {
		case float64:
			if math.IsNaN(o) {
				return math.IsNaN(d)
			}
			return math.Abs(o-d) < 1e-9
		case int64:
			return math.Abs(o-float64(d)) < 1e-9
		default:
			return false
		}
	case string:
		s, ok := decoded.(string)
		return ok && o == s
	case []Value:
		slice, ok := decoded.([]interface{})
		if !ok || len(o) != len(slice) {
			return false
		}
		for i := range o {
			if !valuesEqual(o[i], slice[i]) {
				return false
			}
		}
		return true
	case map[string]Value:
		m, ok := decoded.(map[string]interface{})
		if !ok || len(o) != len(m) {
			return false
		}
		for k, v := range o {
			dv, exists := m[k]
			if !exists || !valuesEqual(v, dv) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// valuesEqualWithNumber compares where decoded numbers may be stdjson.Number, returns bool for property-based testing.
func valuesEqualWithNumber(orig Value, decoded interface{}) bool {
	switch o := orig.(type) {
	case nil:
		return decoded == nil
	case bool:
		db, ok := decoded.(bool)
		return ok && o == db
	case int64:
		switch d := decoded.(type) {
		case float64:
			return math.Abs(float64(o)-d) < 1e-9
		case int64:
			return o == d
		case stdjson.Number:
			fv, err := d.Float64()
			if err != nil {
				return false
			}
			return math.Abs(float64(o)-fv) < 1e-9
		default:
			return false
		}
	case float64:
		switch v := decoded.(type) {
		case float64:
			if math.IsNaN(o) {
				return math.IsNaN(v)
			}
			return math.Abs(o-v) < 1e-9
		case int64:
			return math.Abs(o-float64(v)) < 1e-9
		case stdjson.Number:
			fv, err := v.Float64()
			if err != nil {
				return false
			}
			if math.IsNaN(o) {
				return math.IsNaN(fv)
			}
			return math.Abs(o-fv) < 1e-9
		default:
			return false
		}
	case string:
		s, ok := decoded.(string)
		return ok && o == s
	case []Value:
		slice, ok := decoded.([]interface{})
		if !ok || len(o) != len(slice) {
			return false
		}
		for i := range o {
			if !valuesEqualWithNumber(o[i], slice[i]) {
				return false
			}
		}
		return true
	case map[string]Value:
		m, ok := decoded.(map[string]interface{})
		if !ok || len(o) != len(m) {
			return false
		}
		for k, v := range o {
			dv, exists := m[k]
			if !exists || !valuesEqualWithNumber(v, dv) {
				return false
			}
		}
		return true
	default:
		return false
	}
}
