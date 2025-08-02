package ohsnap

type Mapped[T, U any] struct {
	Value U
	raw   T
}

type mappedArbitrary[T, U any] struct {
	arbitrary Arbitrary[T]
	f         func(T) U
}

// Map transforms a generator of type T into a generator of type U.
func Map[T, U any](
	arb Arbitrary[T],
	f func(T) U,
) Arbitrary[Mapped[T, U]] {
	return &mappedArbitrary[T, U]{
		arbitrary: arb,
		f:         f,
	}
}

func (a *mappedArbitrary[T, U]) Generate() Mapped[T, U] {
	raw := a.arbitrary.Generate()

	return Mapped[T, U]{
		Value: a.f(raw),
		raw:   raw,
	}
}

func (a *mappedArbitrary[T, U]) Shrink(m Mapped[T, U]) []Mapped[T, U] {
	var results []Mapped[T, U]

	for _, x := range a.arbitrary.Shrink(m.raw) {
		results = append(results, Mapped[T, U]{
			Value: a.f(x),
			raw:   x,
		})
	}

	return results
}
