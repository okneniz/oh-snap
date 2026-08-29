package ohsnap

import (
	"iter"
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

func (a *arbitraryOneOf[T]) Generate() iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			idx := a.rand.IntN(len(a.arbs))
			arb := a.arbs[idx]
			value := First(arb.Generate())
			if !yield(value) {
				return
			}
		}
	}
}

func (a *arbitraryOneOf[T]) Shrink(T) iter.Seq[T] {
	return Empty[T]()
}
