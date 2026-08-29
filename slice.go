package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitrarySlice[T any] struct {
	rand   *rand.Rand
	elem   Arbitrary[T]
	minLen int
	maxLen int
}

// ArbitrarySlice - return generator for arbitrary slices.
// rnd - pseudo-random number generator.
// elem - generator of arbitrary elements of slice.
// minLen and maxLen - bounds of length of generated values.
func ArbitrarySlice[T any](
	rnd *rand.Rand,
	elem Arbitrary[T],
	minLen, maxLen int,
) Arbitrary[[]T] {
	return &arbitrarySlice[T]{
		rand:   rnd,
		elem:   elem,
		minLen: minLen,
		maxLen: maxLen,
	}
}

func (a *arbitrarySlice[T]) Generate() iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for {
			length := a.rand.IntN(a.maxLen-a.minLen+1) + int(a.minLen)

			slice := make([]T, length)
			for i := range slice {
				slice[i] = First(a.elem.Generate())
			}

			if !yield(slice) {
				return
			}
		}
	}
}

func (a *arbitrarySlice[T]) Shrink(slice []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(slice) > 0 {
			half := len(slice) / 2
			halfSlice := slice[:half]
			if !yield(halfSlice) {
				return
			}
		}

		for i := range slice {
			for smallerElem := range a.elem.Shrink(slice[i]) {
				newSlice := make([]T, len(slice))
				copy(newSlice, slice)
				newSlice[i] = smallerElem
				if !yield(newSlice) {
					return
				}
			}
		}

		if len(slice) > 0 {
			for i := range slice {
				newSlice := make([]T, 0, len(slice)-1)
				newSlice = append(newSlice, slice[:i]...)
				newSlice = append(newSlice, slice[i+1:]...)
				if !yield(newSlice) {
					return
				}
			}
		}
	}
}
