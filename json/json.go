package json

import (
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
)

// Value represents a JSON value. It can be one of:
// - nil
// - bool
// - float64 (for numbers)
// - int64 (for integer numbers)
// - string
// - []Value (array)
// - map[string]Value (object)
type Value any

type arbitraryJSON struct {
	rnd      *rand.Rand
	maxDepth int
	maxSize  int
	letters  string

	// cached underlying arbitraries
	arbFloat  ohsnap.Arbitrary[float64]
	arbInt    ohsnap.Arbitrary[int64]
	arbString ohsnap.Arbitrary[string]
	arbBool   ohsnap.Arbitrary[bool]
	arbKey    ohsnap.Arbitrary[string]

	// collection arbitraries (arrays and objects) that reuse a depth-wrapped value arb
	arbArray  ohsnap.Arbitrary[[]Value]
	arbObject ohsnap.Arbitrary[map[string]Value]

	// small-range int arbitraries for choice selection (created once)
	arbChoice ohsnap.Arbitrary[int] // Used with % to select among type variants
}

// depthArbitrary wraps parent arbitrary to produce values limited by a fixed remaining depth.
// It implements ohsnap.Arbitrary[Value].
type depthArbitrary struct {
	parent *arbitraryJSON
	depth  int
}

func (d *depthArbitrary) Generate() Value {
	return d.parent.generateAt(d.depth)
}

func (d *depthArbitrary) Shrink(v Value) []Value {
	return d.parent.Shrink(v)
}

// ArbitraryJSON creates and caches all required arbitraries (including collection arbitraries).
// rnd - pseudo-random number generator (seedable).
// maxDepth - maximum nesting depth for arrays/objects (>=0).
// maxSize - maximum array length / number of object keys.
func ArbitraryJSON(rnd *rand.Rand, maxDepth, maxSize int) ohsnap.Arbitrary[Value] {
	if maxDepth < 0 {
		maxDepth = 0
	}
	if maxSize < 0 {
		maxSize = 0
	}

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
	aj := &arbitraryJSON{
		rnd:      rnd,
		maxDepth: maxDepth,
		maxSize:  maxSize,
		letters:  letters,
	}

	// Create and cache primitive arbitraries once.
	// create primitive arbitraries once and reuse them on every call
	aj.arbFloat = ohsnap.ArbitraryFloat64(rnd, -1000, 1000)
	aj.arbInt = ohsnap.ArbitraryInt64(rnd, -1000000, 1000000)
	aj.arbString = ohsnap.ArbitraryString(rnd, letters, 0, 20)
	aj.arbBool = ohsnap.ArbitraryBool(rnd)
	aj.arbKey = ohsnap.ArbitraryString(rnd, letters, 1, 8)

	// Use a single arbitrary for all type choices, using % to select the variant needed
	aj.arbChoice = ohsnap.ArbitraryInt(rnd, 0, 1000)

	// element depth for nested collections
	elementDepth := maxDepth - 1
	if elementDepth < 0 {
		elementDepth = 0
	}
	depthElem := &depthArbitrary{parent: aj, depth: elementDepth}

	// create collection arbitraries once using depth-wrapped element arbitrary
	aj.arbArray = ohsnap.ArbitrarySlice(rnd, depthElem, 0, maxSize)
	aj.arbObject = ohsnap.ArbitraryMap(rnd, aj.arbKey, depthElem, 0, maxSize)

	return aj
}

// Generate implements ohsnap.Arbitrary[Value]
// It uses cached primitive and collection arbitraries.
func (a *arbitraryJSON) Generate() Value {
	return a.generateAt(a.maxDepth)
}

// generateAt generates a JSON Value using remainingDepth to decide whether nesting is allowed.
// Note: arrays/objects are generated via the cached collection arbitraries which were created
// with element depth = maxDepth-1 at construction time.
func (a *arbitraryJSON) generateAt(remainingDepth int) Value {
	// if no nesting allowed or maxSize==0, generate primitives only
	if remainingDepth <= 0 || a.maxSize == 0 {
		choice := a.arbChoice.Generate() % 5 // null, bool, float, int, string
		switch choice {
		case 0:
			return nil
		case 1:
			return a.arbBool.Generate()
		case 2:
			return a.arbFloat.Generate()
		case 3:
			return a.arbInt.Generate()
		default:
			return a.arbString.Generate()
		}
	}

	// when nesting allowed, include arrays and objects
	choice := a.arbChoice.Generate() % 7 // null, bool, float, int, string, array, object
	switch choice {
	case 0:
		return nil
	case 1:
		return a.arbBool.Generate()
	case 2:
		return a.arbFloat.Generate()
	case 3:
		return a.arbInt.Generate()
	case 4:
		return a.arbString.Generate()
	case 5:
		return a.arbArray.Generate()
	default:
		return a.arbObject.Generate()
	}
}

// Shrink delegates to cached arbitraries' Shrink methods and applies structural shrinking
// for arrays and objects (by converting returned values to []Value).
func (a *arbitraryJSON) Shrink(v Value) []Value {
	switch x := v.(type) {
	case nil:
		return nil
	case bool:
		bs := a.arbBool.Shrink(x)
		out := make([]Value, 0, len(bs))
		for _, b := range bs {
			out = append(out, b)
		}
		return out
	case float64:
		fs := a.arbFloat.Shrink(x)
		out := make([]Value, 0, len(fs))
		for _, f := range fs {
			out = append(out, f)
		}
		return out
	case int64:
		is := a.arbInt.Shrink(x)
		out := make([]Value, 0, len(is))
		for _, i := range is {
			out = append(out, i)
		}
		return out
	case string:
		ss := a.arbString.Shrink(x)
		out := make([]Value, 0, len(ss))
		for _, s := range ss {
			out = append(out, s)
		}
		return out
	case []Value:
		// ArbitrarySlice.Shrink returns [][]T
		arrs := a.arbArray.Shrink(x)
		out := make([]Value, 0, len(arrs))
		for _, arr := range arrs {
			out = append(out, arr)
		}
		// also include structural shrinks
		out = append(out, a.shrinkArray(x)...)
		return out
	case map[string]Value:
		maps := a.arbObject.Shrink(x)
		out := make([]Value, 0, len(maps))
		for _, m := range maps {
			out = append(out, m)
		}
		out = append(out, a.shrinkObject(x)...)
		return out
	default:
		return nil
	}
}

// structural shrinking helpers

func (a *arbitraryJSON) shrinkArray(arr []Value) []Value {
	var shrunk []Value
	n := len(arr)
	if n == 0 {
		return nil
	}

	// prefix half
	half := n / 2
	if half > 0 {
		prefix := make([]Value, half)
		copy(prefix, arr[:half])
		shrunk = append(shrunk, prefix)
	}

	// shrink elements
	for i := range arr {
		elemShrinks := a.Shrink(arr[i])
		for _, s := range elemShrinks {
			newArr := make([]Value, n)
			copy(newArr, arr)
			newArr[i] = s
			shrunk = append(shrunk, newArr)
		}
	}

	// remove single elements
	for i := range arr {
		newArr := make([]Value, 0, n-1)
		newArr = append(newArr, arr[:i]...)
		newArr = append(newArr, arr[i+1:]...)
		shrunk = append(shrunk, newArr)
	}

	// empty
	shrunk = append(shrunk, []Value{})

	return shrunk
}

func (a *arbitraryJSON) shrinkObject(m map[string]Value) []Value {
	var shrunk []Value
	if len(m) == 0 {
		return nil
	}

	// produce smaller object with ~half keys
	halfSize := len(m) / 2
	if halfSize > 0 {
		smaller := make(map[string]Value, halfSize)
		i := 0
		for k, v := range m {
			if i >= halfSize {
				break
			}
			smaller[k] = v
			i++
		}
		shrunk = append(shrunk, smaller)
	}

	// shrink individual values
	for k, v := range m {
		for _, sv := range a.Shrink(v) {
			newMap := make(map[string]Value, len(m))
			for k2, v2 := range m {
				newMap[k2] = v2
			}
			newMap[k] = sv
			shrunk = append(shrunk, newMap)
		}
	}

	// remove single keys
	for k := range m {
		newMap := make(map[string]Value, len(m)-1)
		for k2, v := range m {
			if k2 != k {
				newMap[k2] = v
			}
		}
		shrunk = append(shrunk, newMap)
	}

	// empty object
	shrunk = append(shrunk, map[string]Value{})

	return shrunk
}
