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
	arbs []Arbitrary[T],
) Arbitrary[T] {
	return &arbitraryOneOf[T]{
		rand: rnd,
		arbs: arbs,
	}
}

func (a *arbitraryOneOf[T]) Generate() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			idx := a.rand.IntN(len(a.arbs))
			arb := a.arbs[idx]
			a.lastArb = arb
			value := First(arb.Generate())
			if !yield(value) {
				return
			}
		}
	}
}

// Shrink delegates shrinking to the generator picked by the last Generate
// call. Before the first Generate it yields no candidates.
func (a *arbitraryOneOf[T]) Shrink(variant T) iter.Seq[T] {
	if a.lastArb == nil {
		return Empty[T]()
	}

	return a.lastArb.Shrink(variant)
}
