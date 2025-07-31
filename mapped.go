package ohsnap

// Map transforms a generator of type T into a generator of type U.
func Map[T, U any](arb Arbitrary[T], f func(T) U) Arbitrary[U] {
	return &mappedArbitrary[T, U]{arb, f}
}

type mappedArbitrary[T, U any] struct {
	arb Arbitrary[T]
	f   func(T) U
}

func (m *mappedArbitrary[T, U]) Generate() U {
	return m.f(m.arb.Generate())
}

func (m *mappedArbitrary[T, U]) Shrink(value U) []U {
	// Shrinking is not supported for mapped generators
	return nil
}
