package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryUint16 struct {
	rand     *rand.Rand
	from, to uint16
}

// Uint16 - return generator for arbitrary uint16 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Uint16(rnd *rand.Rand, from, to uint16) ohsnap.Arbitrary[uint16] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryUint16{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryUint16) Generate() uint16 {
	x := uint(a.to - a.from)
	if x == 0 {
		x++
	}

	return uint16(a.rand.UintN(x)) + a.from
}

func (arbitraryUint16) Shrink(value uint16) iter.Seq[uint16] {
	return shrink.Halving[uint16](0)(value)
}
