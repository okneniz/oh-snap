package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryInt64 struct {
	rand     *rand.Rand
	from, to int64
}

// Int64 - return generator for arbitrary int64 values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Int64(rnd *rand.Rand, from, to int64) ohsnap.Arbitrary[int64] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryInt64{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryInt64) Generate() int64 {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Int64N(x) + a.from
}

func (arbitraryInt64) Shrink(value int64) iter.Seq[int64] {
	return shrink.Halving[int64](0)(value)
}
