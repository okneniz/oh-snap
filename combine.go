package ohsnap

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

func (c *combinedArbitrary[T, U]) Generate() Pair[T, U] {
	return Pair[T, U]{
		First:  c.arbT.Generate(),
		Second: c.arbU.Generate(),
	}
}

func (c *combinedArbitrary[T, U]) Shrink(value Pair[T, U]) []Pair[T, U] {
	var results []Pair[T, U]

	for _, t := range c.arbT.Shrink(value.First) {
		results = append(results, Pair[T, U]{t, value.Second})
	}

	for _, u := range c.arbU.Shrink(value.Second) {
		results = append(results, Pair[T, U]{value.First, u})
	}

	return results
}
