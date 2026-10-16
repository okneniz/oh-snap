package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryInt32 struct {
	rand     *rand.Rand
	from, to int32
}

// Int32 - return generator for arbitray int32 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Int32(rnd *rand.Rand, from, to int32) ohsnap.Arbitrary[int32] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt32{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt32) Generate() int32 {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Int32N(x) + a.from
}

func (arbitraryInt32) Shrink(value int32) iter.Seq[int32] {
	return shrink.Halving[int32](0)(value)
}
