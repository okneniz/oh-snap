package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryUint8 struct {
	rand     *rand.Rand
	from, to uint8
}

// Uint8 - return generator for arbitrary uint8 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Uint8(rnd *rand.Rand, from, to uint8) ohsnap.Arbitrary[uint8] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint8{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint8) Generate() uint8 {
	x := uint(a.to - a.from)
	if x == 0 {
		x++
	}

	return uint8(a.rand.UintN(x)) + a.from
}

func (arbitraryUint8) Shrink(value uint8) iter.Seq[uint8] {
	return shrink.Halving[uint8](0)(value)
}
