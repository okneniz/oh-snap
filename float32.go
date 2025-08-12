package ohsnap

import (
	"math/rand/v2"
)

type arbitraryFloat32 struct {
	rand     *rand.Rand
	from, to float32
}

// ArbitraryFloat32 - return generator for arbitrary float32 values.
// - rnd - pseudo-random number generator.
// - from and to - bounds of generated values.
func ArbitraryFloat32(rnd *rand.Rand, from, to float32) Arbitrary[float32] {
	return &arbitraryFloat32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryFloat32) Generate() float32 {
	return a.from + a.rand.Float32()*(a.to-a.from)
}

func (arbitraryFloat32) Shrink(value float32) []float32 {
	var results []float32

	for value != 0 {
		value /= 2.0
		results = append(results, value)
	}

	return results
}
