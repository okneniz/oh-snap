package ohsnap

import "iter"

type mappedArbitrary[T, U any] struct {
	arbitrary Arbitrary[T]
	f         func(T) U
}

// Map transforms a generator of type T into a generator of type U.
func Map[T, U any](
	arb Arbitrary[T],
	f func(T) U,
) Arbitrary[U] {
	return &mappedArbitrary[T, U]{
		arbitrary: arb,
		f:         f,
	}
}

func (a *mappedArbitrary[T, U]) Generate() iter.Seq[U] {
	return func(yield func(U) bool) {
		for {
			value := a.f(First(a.arbitrary.Generate()))
			if !yield(value) {
				return
			}
		}
	}
}

func (a *mappedArbitrary[T, U]) Shrink(U) iter.Seq[U] {
	return Empty[U]()
}
