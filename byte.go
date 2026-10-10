package ohsnap

import (
	"iter"
	"math/rand/v2"

	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryByte struct {
	rand     *rand.Rand
	from, to byte
}

// ArbitraryByte - return generator for arbitrary byte values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryByte(rnd *rand.Rand, from, to byte) Arbitrary[byte] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryByte{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryByte) Generate() byte {
	x := uint(a.to - a.from)
	if x == 0 {
		x++
	}

	return byte(a.rand.UintN(x)) + a.from
}

func (arbitraryByte) Shrink(value byte) iter.Seq[byte] {
	return shrink.Halving[byte](0)(value)
}
