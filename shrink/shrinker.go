// Package shrink provides composable shrinking strategies for oh-snap.
//
// A Shrinker is a function from a value to a lazy sequence of simpler
// candidate values. Strategies compose with Concat, Limit and Interleave,
// and plug into any generator via ohsnap.WithShrinker:
//
//	arb := ohsnap.WithShrinker(
//	    ohsnap.ArbitraryInt(rnd, 100, 200),
//	    shrink.Halving(100), // stay inside the range, shrink toward 100
//	)
package shrink

import "iter"

// Shrinker is a shrinking strategy: it returns a lazy sequence of
// candidates simpler than the input value.
// Candidates must differ from the input value, otherwise the search
// may not terminate.
type Shrinker[T any] func(T) iter.Seq[T]

// Number is a constraint for numeric types supported by numeric strategies.
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// NoShrink returns a strategy without candidates.
func NoShrink[T any]() Shrinker[T] {
	return func(T) iter.Seq[T] {
		return func(yield func(T) bool) {}
	}
}
