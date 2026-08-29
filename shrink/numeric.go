package shrink

import (
	"iter"
)

// Halving returns a strategy that halves the distance between the value
// and the target. It generalizes the default behavior of numeric
// generators (target 0) and keeps candidates inside the generator range
// when target is the lower bound of the range.
func Halving[T Number](target T) Shrinker[T] {
	return func(value T) iter.Seq[T] {
		return func(yield func(T) bool) {
			for value != target {
				var candidate T
				if value > target {
					candidate = target + (value-target)/2
				} else {
					candidate = target - (target-value)/2
				}

				if !yield(candidate) {
					return
				}

				value = candidate
			}
		}
	}
}

// Boundaries returns a strategy that yields the given interesting values
// (0, 1, -1, range bounds and so on) in order.
// The input value itself is skipped to keep the search terminating.
func Boundaries[T Number](values ...T) Shrinker[T] {
	return func(value T) iter.Seq[T] {
		return func(yield func(T) bool) {
			for _, v := range values {
				if v == value {
					continue
				}

				if !yield(v) {
					return
				}
			}
		}
	}
}
