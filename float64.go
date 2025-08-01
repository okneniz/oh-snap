package ohsnap

import (
	"math/rand/v2"
)

type arbitraryFloat64 struct {
	rand     *rand.Rand
	from, to float64
}

func ArbitraryFloat64(rnd *rand.Rand, from, to float64) Arbitrary[float64] {
	return &arbitraryFloat64{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryFloat64) Generate() float64 {
	return a.from + a.rand.Float64()*(a.to-a.from)
}

func (arbitraryFloat64) Shrink(value float64) []float64 {
	var results []float64

	for value != 0 {
		value /= 2.0
		results = append(results, value)
	}

	return results
}
