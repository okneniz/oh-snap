package ohsnap

import (
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

func (a *arbitrarySlice[T]) Generate() []T {
	length := a.rand.IntN(a.maxLen-a.minLen+1) + int(a.minLen)

	slice := make([]T, length)
	for i := range slice {
		slice[i] = a.elem.Generate()
	}

	return slice
}

func (a *arbitrarySlice[T]) Shrink(slice []T) [][]T {
	var shrunk [][]T

	if len(slice) > 0 {
		half := len(slice) / 2
		shrunk = append(shrunk, slice[:half])
	}

	for i := range slice {
		for _, smallerElem := range a.elem.Shrink(slice[i]) {
			newSlice := make([]T, len(slice))
			copy(newSlice, slice)
			newSlice[i] = smallerElem
			shrunk = append(shrunk, newSlice)
		}
	}

	if len(slice) > 0 {
		for i := range slice {
			newSlice := make([]T, 0, len(slice)-1)
			newSlice = append(newSlice, slice[:i]...)
			newSlice = append(newSlice, slice[i+1:]...)
			shrunk = append(shrunk, newSlice)
		}
	}

	return shrunk
}
