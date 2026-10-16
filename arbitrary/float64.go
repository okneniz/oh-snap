package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryFloat64 struct {
	rand     *rand.Rand
	from, to float64
}

// Float64 - return generator for arbitrary float64 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Float64(rnd *rand.Rand, from, to float64) ohsnap.Arbitrary[float64] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryFloat64{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryFloat64) Generate() float64 {
	return a.from + a.rand.Float64()*(a.to-a.from)
}

func (arbitraryFloat64) Shrink(value float64) iter.Seq[float64] {
	return shrink.Halving[float64](0)(value)
}
