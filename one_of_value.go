package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryOneOfValue[T any] struct {
	rand   *rand.Rand
	values []T
}

// OneOfValue - return generator for arbitrary values from list.
// rnd - pseudo-random number generator.
// values - allowed values.
func OneOfValue[T any](
	rnd *rand.Rand,
	values ...T,
) Arbitrary[T] {
	return &arbitraryOneOfValue[T]{
		rand:   rnd,
		values: values,
	}
}

func (a *arbitraryOneOfValue[T]) Generate() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			idx := a.rand.IntN(len(a.values))
			value := a.values[idx]
			if !yield(value) {
				return
			}
		}
	}
}

func (a *arbitraryOneOfValue[T]) Shrink(T) iter.Seq[T] {
	return Empty[T]()
}
