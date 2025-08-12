package ohsnap

import (
	"math/rand/v2"
)

type arbitraryOneOfValue[T any] struct {
	rand   *rand.Rand
	values []T
}

// OneOfValue - return generator for arbitrary values from list.
// - rnd - pseudo-random number generator.
// - values - allowed values.
func OneOfValue[T any](
	rnd *rand.Rand,
	values ...T,
) Arbitrary[T] {
	return &arbitraryOneOfValue[T]{
		rand:   rnd,
		values: values,
	}
}

func (a *arbitraryOneOfValue[T]) Generate() T {
	idx := a.rand.IntN(len(a.values))
	return a.values[idx]
}

func (a *arbitraryOneOfValue[T]) Shrink(value T) []T {
	return nil
}
