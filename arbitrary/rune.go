package arbitrary

import (
	"iter"
	"math/rand/v2"

	ohsnap "github.com/okneniz/oh-snap"
	"github.com/okneniz/oh-snap/shrink"
)

type arbitraryRune struct {
	rand     *rand.Rand
	from, to rune
}

// Rune - return generator for arbitrary rune values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func Rune(rnd *rand.Rand, from, to rune) ohsnap.Arbitrary[rune] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryRune{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryRune) Generate() rune {
	x := a.to - a.from
	if x == 0 {
		x++
	}

	return a.rand.Int32N(x) + a.from
}

func (arbitraryRune) Shrink(value rune) iter.Seq[rune] {
	return shrink.Halving[rune](0)(value)
}
