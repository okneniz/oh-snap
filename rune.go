package ohsnap

import (
	"iter"
	"math/rand/v2"
)

type arbitraryRune struct {
	rand     *rand.Rand
	from, to rune
}

// ArbitraryRune - return generator for arbitrary rune values.
// rnd - pseudo-random number generator.
// from and to - bounds of generated values.
func ArbitraryRune(rnd *rand.Rand, from, to rune) Arbitrary[rune] {
	if from > to {
		from, to = to, from
	}

	return &arbitraryRune{
		rand: rnd,
		from: from,
		to:   to,
	}
}

func (a arbitraryRune) Generate() iter.Seq[rune] {
	return func(yield func(rune) bool) {
		x := a.to - a.from
		if x == 0 {
			x++
		}

		for {
			value := a.rand.Int32N(x) + a.from
			if !yield(value) {
				return
			}
		}
	}
}

func (arbitraryRune) Shrink(value rune) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for value != 0 {
			value /= 2
			if !yield(value) {
				return
			}
		}
	}
}
