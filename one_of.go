package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryOneOf[T any] struct {
	rand    *rand.Rand
	arbs    []Arbitrary[T]
	lastArb Arbitrary[T] // generator picked by the last Generate call
}

func OneOf[T any](
	rnd *rand.Rand,
	arbs ...Arbitrary[T],
) Arbitrary[T] {
	if len(arbs) == 1 {
		return arbs[0]
	}

	return &arbitraryOneOf[T]{
		rand: rnd,
		arbs: arbs,
	}
}

func (a *arbitraryOneOf[T]) Generate() T {
	idx := a.rand.IntN(len(a.arbs))
	arb := a.arbs[idx]
	a.lastArb = arb
	return arb.Generate()
}

// Shrink delegates shrinking to the generator picked by the last Generate
// call. Before the first Generate it yields no candidates.
func (a *arbitraryOneOf[T]) Shrink(variant T) iter.Seq[T] {
	if a.lastArb == nil {
		return empty[T]()
	}

	return a.lastArb.Shrink(variant)
}
