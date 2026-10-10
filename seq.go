package ohsnap

import "iter"

// Empty returns an empty sequence.
// It is useful for Shrink implementations without candidates.
func Empty[T any]() iter.Seq[T] {
	return func(yield func(T) bool) {}
}
