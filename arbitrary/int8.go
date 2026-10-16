package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryInt8 struct {
	rand     *rand.Rand
	from, to int8
}

// Int8 - return generator for arbitrary int8 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Int8(rnd *rand.Rand, from, to int8) ohsnap.Arbitrary[int8] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt8{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt8) Generate() int8 {
	x := int(a.to - a.from)
	if x == 0 {
		x++
	}

	return int8(a.rand.IntN(x)) + a.from
}

func (arbitraryInt8) Shrink(value int8) iter.Seq[int8] {
	return shrink.Halving[int8](0)(value)
}
