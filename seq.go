package ohsnap

import "iter"

// empty returns an empty sequence.
// It is useful for Shrink implementations without candidates.
func empty[T any]() iter.Seq[T] {
	return func(func(T) bool) {}
}
