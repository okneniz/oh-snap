package ohsnap

import "iter"

// First returns the first element of the sequence.
// It is useful to pull a single random value from Generate.
func First[T any](seq iter.Seq[T]) T {
	for value := range seq {
		return value
	}

	var zero T
	return zero
}

// Empty returns an empty sequence.
// It is useful for Shrink implementations without candidates.
func Empty[T any]() iter.Seq[T] {
	return func(yield func(T) bool) {}
}
