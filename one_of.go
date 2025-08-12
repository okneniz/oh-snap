package ohsnap

import (
	"math/rand/v2"
)

type arbitraryOneOf[T any] struct {
	rand *rand.Rand
	arbs []Arbitrary[T]
}

func OneOf[T any](
	rnd *rand.Rand,
	arbs []Arbitrary[T],
) Arbitrary[T] {
	return &arbitraryOneOf[T]{
		rand: rnd,
		arbs: arbs,
	}
}

func (a *arbitraryOneOf[T]) Generate() T {
	idx := a.rand.IntN(len(a.arbs))
	arb := a.arbs[idx]
	return arb.Generate()
}

func (a *arbitraryOneOf[T]) Shrink(variant T) []T {
	return nil
}
