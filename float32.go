package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryFloat32 struct {
	rand     *rand.Rand
	from, to float32
}

// ArbitraryFloat32 - return generator for arbitrary float32 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryFloat32(rnd *rand.Rand, from, to float32) Arbitrary[float32] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryFloat32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryFloat32) Generate() float32 {
	return a.from + a.rand.Float32()*(a.to-a.from)
}

func (arbitraryFloat32) Shrink(value float32) iter.Seq[float32] {
	return shrink.Halving[float32](0)(value)
}
