package ohsnap

import (
	"math/rand/v2"
)

type Variant[T any] struct {
	Value     T
	arbitrary Arbitrary[T]
}

type arbitraryOneOf[T any] struct {
	rand *rand.Rand
	arbs []Arbitrary[T]
}

func OneOf[T any](
	rnd *rand.Rand,
	arbs []Arbitrary[T],
) Arbitrary[Variant[T]] {
	return &arbitraryOneOf[T]{
		rand: rnd,
		arbs: arbs,
	}
}

func (a *arbitraryOneOf[T]) Generate() Variant[T] {
	idx := a.rand.IntN(len(a.arbs))
	arb := a.arbs[idx]

	return Variant[T]{
		Value:     arb.Generate(),
		arbitrary: arb,
	}
}

func (a *arbitraryOneOf[T]) Shrink(variant Variant[T]) []Variant[T] {
	var results []Variant[T]

	for _, x := range variant.arbitrary.Shrink(variant.Value) {
		results = append(results, Variant[T]{
			Value:     x,
			arbitrary: variant.arbitrary,
		})
	}

	return results
}
