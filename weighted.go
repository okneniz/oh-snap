package ohsnap

import (
	"iter"
	"math/rand/v2"
	"sort"
)

// Weighted returns an Arbitrary that selects from the provided arbitraries
// according to their weights. The map key is the weight (must be > 0), and
// the value is the Arbitrary to use for that weight.
//
// Example:
//
//	arb := Weighted(rnd, map[int]Arbitrary[int]{
//	    1: ArbitraryInt(rnd, 0, 10),
//	    3: ArbitraryInt(rnd, 100, 200),
//	})
func Weighted[T any](rnd *rand.Rand, weightedArbs map[int]Arbitrary[T]) Arbitrary[T] {
	// Flatten map into slices for deterministic iteration order
	type entry struct {
		weight int
		arb    Arbitrary[T]
	}
	var entries []entry
	totalWeight := 0
	for w, arb := range weightedArbs {
		if w <= 0 {
			continue // skip zero or negative weights
		}
		entries = append(entries, entry{weight: w, arb: arb})
		totalWeight += w
	}
	// Sort for deterministic behavior (optional, but good for tests)
	sort.Slice(entries, func(i, j int) bool { return entries[i].weight < entries[j].weight })

	// Convert []entry to []arbitraryWeightedEntry for storage in struct
	typedEntries := make([]*arbitraryWeightedEntry[T], len(entries))
	for i, e := range entries {
		typedEntries[i] = &arbitraryWeightedEntry[T]{
			weight: e.weight,
			arb:    e.arb,
		}
	}

	return &arbitraryWeighted[T]{
		rand:        rnd,
		entries:     typedEntries,
		totalWeight: totalWeight,
	}
}

type arbitraryWeightedEntry[T any] struct {
	weight int
	arb    Arbitrary[T]
}

type arbitraryWeighted[T any] struct {
	rand        *rand.Rand
	entries     []*arbitraryWeightedEntry[T]
	totalWeight int
	lastArb     Arbitrary[T] // generator picked by the last Generate call
}

func (a *arbitraryWeighted[T]) Generate() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			value := a.pick()
			if !yield(value) {
				return
			}
		}
	}
}

func (a *arbitraryWeighted[T]) pick() T {
	if a.totalWeight == 0 || len(a.entries) == 0 {
		var zero T
		return zero
	}
	r := a.rand.IntN(a.totalWeight)
	acc := 0
	for _, e := range a.entries {
		acc += e.weight
		if r < acc {
			a.lastArb = e.arb
			return First(e.arb.Generate())
		}
	}
	// Fallback (should not happen)
	last := a.entries[len(a.entries)-1].arb
	a.lastArb = last
	return First(last.Generate())
}

// Shrink delegates shrinking to the generator picked by the last Generate
// call. Before the first Generate it yields no candidates.
func (a *arbitraryWeighted[T]) Shrink(value T) iter.Seq[T] {
	if a.lastArb == nil {
		return Empty[T]()
	}

	return a.lastArb.Shrink(value)
}
