package ohsnap

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

func (a *mappedArbitrary[T, U]) Generate() U {
	return a.f(a.arbitrary.Generate())
}

func (a *mappedArbitrary[T, U]) Shrink(m U) []U {
	return nil
}
