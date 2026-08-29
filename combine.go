package ohsnap

import "iter"

// Pair represents a pair of values of types T and U.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Combine combines two generators into a generator of pairs.
func Combine[T, U any](arbT Arbitrary[T], arbU Arbitrary[U]) Arbitrary[Pair[T, U]] {
	return &combinedArbitrary[T, U]{arbT, arbU}
}

type combinedArbitrary[T, U any] struct {
	arbT Arbitrary[T]
	arbU Arbitrary[U]
}

func (c *combinedArbitrary[T, U]) Generate() iter.Seq[Pair[T, U]] {
	return func(yield func(Pair[T, U]) bool) {
		for {
			pair := Pair[T, U]{
				First:  First(c.arbT.Generate()),
				Second: First(c.arbU.Generate()),
			}
			if !yield(pair) {
				return
			}
		}
	}
}

func (c *combinedArbitrary[T, U]) Shrink(value Pair[T, U]) iter.Seq[Pair[T, U]] {
	return func(yield func(Pair[T, U]) bool) {
		for t := range c.arbT.Shrink(value.First) {
			pair := Pair[T, U]{t, value.Second}
			if !yield(pair) {
				return
			}
		}

		for u := range c.arbU.Shrink(value.Second) {
			pair := Pair[T, U]{value.First, u}
			if !yield(pair) {
				return
			}
		}
	}
}
