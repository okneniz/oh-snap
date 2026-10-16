package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryUint32 struct {
	rand     *rand.Rand
	from, to uint32
}

// Uint32 - return generator for arbitrary uint32 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Uint32(rnd *rand.Rand, from, to uint32) ohsnap.Arbitrary[uint32] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint32) Generate() uint32 {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Uint32N(x) + a.from
}

func (arbitraryUint32) Shrink(value uint32) iter.Seq[uint32] {
	return shrink.Halving[uint32](0)(value)
}
